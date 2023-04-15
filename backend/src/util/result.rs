use std::io::Cursor;

use rocket::{
    http::{ContentType, Status},
    response::{self, Responder},
    serde::json::serde_json::json,
    Request, Response,
};
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(tag = "type")]
pub enum Error {
    InvalidStreamKey,
    InvalidSessionToken,

    InternalError,
    NotFound,
}

pub type Result<T, E = Error> = std::result::Result<T, E>;

impl<'r> Responder<'r, 'static> for Error {
    fn respond_to(self, _: &'r Request<'_>) -> response::Result<'static> {
        let status_code = match self {
            Error::InvalidStreamKey => Status::BadRequest,
            Error::InvalidSessionToken => Status::Unauthorized,

            Error::InternalError => Status::InternalServerError,
            Error::NotFound => Status::NotFound,
        };

        let error_body = json!(self).to_string();

        Response::build()
            .status(status_code)
            .header(ContentType::new("application", "json"))
            .sized_body(error_body.len(), Cursor::new(error_body))
            .ok()
    }
}
