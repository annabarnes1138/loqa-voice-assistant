use axum::{
    extract::Json,
    routing::post,
    Router,
};
use serde::Deserialize;
use tokio::net::TcpListener;
use tracing_subscriber;

#[derive(Deserialize, Debug)]
struct WakePayload {
    event: String,
    device_id: String,
}

async fn handle_wake(Json(payload): Json<WakePayload>) {
    tracing::info!("Wake event received from {}: {:?}", payload.device_id, payload);
}

#[tokio::main]
async fn main() {
    tracing_subscriber::fmt()
        .with_max_level(tracing::Level::INFO)
        .with_target(false)
        .init();

    let app = Router::new().route("/wake", post(handle_wake));

    let listener = TcpListener::bind("0.0.0.0:3000").await.unwrap();
    tracing::info!("🚀 loqa-hub listening on http://0.0.0.0:3000");

    axum::serve(listener, app).await.unwrap();
}