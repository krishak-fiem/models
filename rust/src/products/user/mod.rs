use db::es::es_operation;
use serde::{Deserialize, Serialize};
use serde_json::Value;

#[derive(Debug, Serialize, Deserialize)]
pub struct User {
    pub name: String,
    pub email: String,
}

impl User {
    pub async fn add_user(&self) {
        es_operation(
            &db::es::get_client(),
            serde_json::to_value(self).unwrap(),
            "users",
        )
        .await;
    }

    pub async fn find_user(&self, query: Value) {
        es_operation(&db::es::get_client(), query, "users").await;
    }
}
