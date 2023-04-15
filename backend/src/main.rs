#[macro_use]
extern crate rocket;

extern crate dotenv;

use std::str::FromStr;
use std::sync::Arc;

use rocket_cors::AllowedOrigins;

pub mod prisma;
pub mod routes;
pub mod util;
pub mod auth;

#[derive(Clone)]
pub struct Context {
    pub prisma: Arc<prisma::PrismaClient>,
}
pub type Ctx = rocket::State<Context>;

#[launch]
async fn rocket() -> _ {
    dotenv::dotenv().ok();

    let prisma = Arc::new(
        prisma::new_client()
            .await
            .expect("Failed to create Prisma client"),
    );
    #[cfg(debug_assert)]
    prisma._db_push(false).await.unwrap();

    let cors = rocket_cors::CorsOptions {
        allowed_origins: AllowedOrigins::All,
        allowed_methods: [
            "Get", "Put", "Post", "Delete", "Options", "Head", "Trace", "Connect", "Patch",
        ]
        .iter()
        .map(|s| FromStr::from_str(s).unwrap())
        .collect(),
        ..Default::default()
    }
    .to_cors()
    .unwrap();

    rocket::build()
        .manage(Context { prisma })
        .mount("/rtmp", routes::rtmp::routes())
        .mount("/users", routes::users::routes())
        .attach(cors)
}
