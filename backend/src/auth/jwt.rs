use std::env;

use chrono::Utc;
use jsonwebtoken::{Header, Algorithm, encode, EncodingKey, decode, Validation};
use jsonwebtoken::errors::{Error, ErrorKind};
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
pub struct Claims {
    pub sub: i32,
    pub exp: usize,
}

pub fn encode_user_token_jwt(user_id: i32) -> Result<String, Error>{
    let expiration = Utc::now()
        .checked_add_signed(chrono::Duration::hours(1))
        .expect("Invalid JWT expiration timestamp")
        .timestamp();

    let claims = Claims {
        sub: user_id,
        exp: expiration as usize,
    };

    let header = Header::new(Algorithm::HS512);

    let jwt_secret = env::var("JWT_SECRET").expect("Env variable JWT_SECRET not set");
    encode(&header, &claims, &EncodingKey::from_secret(jwt_secret.as_bytes()))
}

pub fn decode_user_token_jwt(jwt_token: String) -> Result<Claims, ErrorKind> {
    let jwt_secret = env::var("JWT_SECRET").expect("Env variable JWT_SECRET not set");

    match decode(&jwt_token, &jwt_secret, &Validation::new(Algorithm::HS512)) {
        Ok(jwt_token) => Ok(jwt_token.claims),
        Err(error) => Err(error.kind().to_owned())
    }
}