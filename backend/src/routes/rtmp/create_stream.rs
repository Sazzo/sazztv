use chrono::{FixedOffset, Utc};
use rocket::form::Form;
use rocket::response::Redirect;

use crate::prisma::user;
use crate::util::result::{Error, Result};
use crate::Ctx;

#[derive(FromForm)]
pub struct RTMPPlayArguments<'r> {
    pub name: &'r str,
}

#[post("/create-stream", data = "<args>")]
pub async fn req(ctx: &Ctx, args: Form<RTMPPlayArguments<'_>>) -> Result<Redirect> {
    let user_with_streamkey_query = ctx
        .prisma
        .user()
        .update(
            user::stream_key::equals(args.name.to_string()),
            vec![user::last_stream_at::set(
                Utc::now().with_timezone(&FixedOffset::east_opt(0).unwrap()),
            )],
        )
        .with(user::stream::fetch())
        .exec()
        .await;

    if user_with_streamkey_query.is_err() {
        return Err(Error::InvalidStreamKey);
    }

    let user_with_streamkey = user_with_streamkey_query.unwrap();

    // user_with_streamkey.stream returns Some(None) if the user has no stream.
    if user_with_streamkey.stream.unwrap().is_none() {
        ctx.prisma
            .stream()
            .create(user::id::equals(user_with_streamkey.id), vec![])
            .exec()
            .await
            .unwrap();
    }

    Ok(Redirect::to(format!("{}", user_with_streamkey.username)))
}
