use rocket::form::Form;

use super::create_stream::RTMPPlayArguments;

use crate::prisma::{stream, user};
use crate::util::result::{Error, Result};
use crate::Ctx;

#[post("/delete-stream", data = "<args>")]
pub async fn req(ctx: &Ctx, args: Form<RTMPPlayArguments<'_>>) -> Result<()> {
    let user_with_streamkey_query = ctx
        .prisma
        .user()
        .find_unique(user::stream_key::equals(args.name.to_string()))
        .exec()
        .await
        .unwrap();

    if user_with_streamkey_query.is_none() {
        return Err(Error::InvalidStreamKey);
    }

    let user_with_streamkey = user_with_streamkey_query.unwrap();
    ctx.prisma
        .stream()
        .delete(stream::streamer_id::equals(user_with_streamkey.id))
        .exec()
        .await
        .unwrap();

    Ok(())
}
