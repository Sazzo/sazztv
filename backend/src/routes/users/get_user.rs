use rocket::serde::json::Json;

use crate::prisma::user;
use crate::util::result::{Error, Result};
use crate::Ctx;

user::select!(user_without_streamkey {
    id
    username
    last_stream_at
    stream: select {
        id
    }
});

#[get("/<username>")]
pub async fn req(ctx: &Ctx, username: String) -> Result<Json<user_without_streamkey::Data>> {
    let user = ctx
        .prisma
        .user()
        .find_unique(user::username::equals(username))
        .select(user_without_streamkey::select())
        .exec()
        .await
        .unwrap();

    if user.is_none() {
        return Err(Error::NotFound);
    }

    Ok(Json(user.unwrap()))
}
