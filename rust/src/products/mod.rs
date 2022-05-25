use serde::{Deserialize, Serialize};

pub mod user;

#[derive(Debug, Serialize, Deserialize)]
pub struct Product {
    pub name: String,
    pub description: String,
    pub id: String,
    pub quantity: i64,
    pub seller: String,
    pub stock: i64,
    pub created_at: String,
}
