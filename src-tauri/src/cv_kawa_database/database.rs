use rusqlite::{Connection, params};
use tauri::api::path;
use time::OffsetDateTime;

pub struct CVKawaDatabase {
    pub conn: Connection,
}

impl CVKawaDatabase {
    fn file_exists(path: &str) -> bool {
        std::path::Path::new(path).exists()
    }

    pub fn new_connection() -> CVKawaDatabase {
        let database_url = path::app_data_dir(&Default::default())
            .ok_or("Failed to get app data directory".to_string())
            .unwrap()
            .join("CVKawa")
            .join("cv_kawa.db");
        if !Self::file_exists(database_url.to_str().unwrap()) {
            std::fs::create_dir_all(database_url.parent().unwrap()).unwrap();
        }
        let conn = Connection::open(&database_url).unwrap();
        CVKawaDatabase { conn }
    }

    pub fn init_database(&self) -> Result<bool, String> {
        self.conn
            .execute(
                "CREATE TABLE IF NOT EXISTS cv_version (
                id INTEGER PRIMARY KEY,
                name TEXT NOT NULL,
                version TEXT NOT NULL,
                parent_version TEXT,
                created_at INTEGER NOT NULL,
                updated_at INTEGER NOT NULL,
                deleted_at INTEGER
            )",
                [],
            )
            .map_err(|e| e.to_string())?;
        self.conn
            .execute(
                "CREATE TABLE IF NOT EXISTS system_config (
                id INTEGER PRIMARY KEY,
                key VARCHAR(255) UNIQUE NOT NULL,
                value VARCHAR(255)
            )",
                [],
            )
            .map_err(|e| e.to_string())?;
        Ok(true)
    }

    pub fn check_version(&self, version: &str) -> Result<bool, String> {
        let mut stmt = self.conn.prepare("SELECT * FROM cv_version WHERE version = ?1 AND deleted_at is null;").map_err(|e| e.to_string())?;
        let result = stmt.exists(params![version]).map_err(|e| e.to_string())?;
        Ok(result)
    }

    pub fn del_version(&self, version: &str) -> Result<(), String> {
        let current_timestamp_millis = OffsetDateTime::now_utc().unix_timestamp();
        self.conn.execute(
            "UPDATE cv_version SET deleted_at = ?1 WHERE version = ?2",
            params![
                current_timestamp_millis.clone(),
                version,
            ],
        ).map_err(|e| e.to_string())?;
        Ok(())
    }

    pub fn save_version(&self, version: &str, name: &str, parent_version: Option<&str>) -> Result<(), String> {
        let current_timestamp_millis = OffsetDateTime::now_utc().unix_timestamp();
        self.conn.execute(
            "insert into cv_version (name, version, parent_version, created_at, updated_at, deleted_at) values (?1, ?2, ?3, ?4, ?5, null);",
            params![
                name,
                version,
                parent_version.unwrap_or("null").to_string(),
                current_timestamp_millis.clone() * 1000,
                current_timestamp_millis.clone() * 1000,
            ],
        ).map_err(|e| e.to_string())?;
        Ok(())
    }
}
