use std::fs;
use std::path::Path;
use serde::{Deserialize, Serialize};
use tauri::api::path;
use uuid::Uuid;
use crate::cv_kawa_database::CVKawaDatabase;
use crate::cv_kawa_main::schema::CV;

#[derive(Debug, Serialize, Deserialize)]
pub struct Version {
    pub id: Option<i64>,
    pub name: String,
    pub version: Option<String>,
    pub parent_version: Option<String>,
    pub created_at: Option<i64>,
    pub updated_at: Option<i64>,
    pub deleted_at: Option<i64>,
}


#[tauri::command]
pub fn version_list() -> Result<Vec<Version>, String> {
    let c = CVKawaDatabase::new_connection();
    let mut stmt = c.conn.prepare("SELECT * FROM cv_version WHERE deleted_at is null ORDER BY updated_at DESC;")
        .map_err(|e| e.to_string()).unwrap();
    let versions_iter = stmt.query_map([], |row| {
        Ok(Version {
            id: row.get(0)?,
            name: row.get(1)?,
            version: row.get(2)?,
            parent_version: row.get(3)?,
            created_at: row.get(4)?,
            updated_at: row.get(5)?,
            deleted_at: row.get(6)?,
        })
    }).unwrap();
    let mut versions = Vec::new();
    for version in versions_iter {
        versions.push(version.unwrap());
    }
    Ok(versions)
}

#[tauri::command]
pub fn load_cv(version: &str) -> Result<CV, String> {
    let c = CVKawaDatabase::new_connection();
    let is_exist = c.check_version(version).map_err(|e| e.to_string())?;
    if !is_exist {
        return Err("version not found".to_string());
    };
    let file_path = path::app_data_dir(&Default::default())
        .ok_or("failed to get app dir".to_string())?
        .join("CVKawa")
        .join(version)
        .join("cv_data.json");
    let cv_string = fs::read_to_string(file_path).map_err(|e| e.to_string())?;
    let mut cv: CV = serde_json::from_str(&cv_string).map_err(|e| e.to_string())?;
    let photo_asset_path = path::app_data_dir(&Default::default())
        .ok_or("failed to get app dir".to_string())?
        .join("CVKawa")
        .join(version)
        .join(cv.cv_base_info.photo_path.clone().unwrap_or("".to_string()));
    cv.cv_base_info.photo_path = Option::from(photo_asset_path.to_str().ok_or("failed to get photo path".to_string())?.to_string());
    Ok(cv)
}

#[tauri::command]
pub fn save_cv(cv: CV, name: &str, parent_version: Option<&str>) -> bool {
    let version = match save2file(&cv) {
        Ok(v) => v,
        Err(e) => {
            eprintln!("failed to save cv to file: {}", e);
            return false;
        }
    };
    let c = CVKawaDatabase::new_connection();
    match c.save_version(&version, name, parent_version) {
        Ok(_) => true,
        Err(_) => {
            eprintln!("failed to save version to file");
            return false;
        }
    }
}

#[tauri::command]
pub fn soft_del(version: &str) -> bool {
    let c = CVKawaDatabase::new_connection();
    match c.del_version(version) {
        Ok(_) => true,
        Err(_) => {
            eprintln!("failed to soft delete version");
            return false;
        }
    }
}

fn save2file(cv: &CV) -> Result<String, String> {
    let version = format!("{}", Uuid::new_v4());
    let dir_path = path::app_data_dir(&Default::default())
        .ok_or("failed to get app dir".to_string())?
        .join("CVKawa")
        .join(version.clone());
    fs::create_dir_all(&dir_path).map_err(|e| e.to_string())?;
    let dest_path = match save_photo(&cv.cv_base_info.photo_path.clone().unwrap_or("".to_string()), &version).map_err(|e| e.to_string()) {
        Ok(p) => p,
        Err(e) => {
            eprintln!("failed to save photo: {}", e);
            return Err(e);
        }
    };
    let mut cv_clone = cv.clone();
    cv_clone.cv_base_info.photo_path = Option::from(dest_path.clone());
    let cv_json = serde_json::to_string(&cv_clone).map_err(|e| e.to_string())?;
    fs::write(dir_path.join("cv_data.json"), cv_json).map_err(|e| e.to_string())?;
    Ok(version)
}

fn save_photo(file_src: &str, version: &str) -> Result<String, String> {
    let file_path = Path::new(file_src);
    let extension = file_path.extension().and_then(|ext| ext.to_str()).unwrap_or("");
    let new_file_name = format!("photo.{}", extension);
    let photo_path = path::app_data_dir(&Default::default())
        .ok_or("failed to get app dir".to_string())?
        .join("CVKawa")
        .join(version)
        .join(new_file_name.clone());
    let path_ref = photo_path.as_path();
    let src = file_path.to_str().ok_or("failed to get file path".to_string())?;
    let dest = path_ref.to_str().ok_or("failed to get file path".to_string())?;
    println!("src: {}, dest: {}", src, dest);
    fs::copy(src, dest).map_err(|e| e.to_string())?;
    Ok(dest.to_string())
}