use rusqlite::params;

use crate::cv_kawa_database::CVKawaDatabase;

#[tauri::command]
pub fn init_system() -> Result<(), String> {
    let c = CVKawaDatabase::new_connection();
    c.init_database().map_err(|e| e.to_string())?;
    Ok(())
}

#[tauri::command]
pub fn get_config(key:String) -> Result<String, String> {
    let c = CVKawaDatabase::new_connection();
    let mut stmt = c
        .conn
        .prepare("SELECT * FROM system_config WHERE key = ?1;")
        .map_err(|e| e.to_string())?;
    let result = stmt
        .query_row(params![key], |row| Ok(row.get(2)?))
        .map_err(|e| e.to_string())?;
    Ok(result)
}

#[tauri::command]
pub async fn set_config(key:String, value:String) -> Result<(), String> {
    let c = CVKawaDatabase::new_connection();
    c.conn
        .execute(
            "INSERT OR REPLACE INTO system_config (key, value) VALUES (?1, ?2);",
            params![key, value],
        )
        .map_err(|e| e.to_string())?;
    Ok(())
}