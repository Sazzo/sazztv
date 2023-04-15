mod get_user;

pub fn routes() -> Vec<rocket::Route> {
    routes![get_user::req]
}
