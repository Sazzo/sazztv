mod create_stream;
mod delete_stream;

pub fn routes() -> Vec<rocket::Route> {
    routes![create_stream::req, delete_stream::req]
}
