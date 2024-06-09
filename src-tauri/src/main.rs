// Prevents additional console window on Windows in release, DO NOT REMOVE!!
#![cfg_attr(not(debug_assertions), windows_subsystem = "windows")]

use cv_kawa_system::*;
// use cv_kawa_database::*;
use cv_kawa_main::*;
mod cv_kawa_database;
mod cv_kawa_system;
mod cv_kawa_main;


fn main() {
    tauri::Builder::default()
        .invoke_handler(tauri::generate_handler![
            init_system,
            version_list,
            load_cv,
            soft_del,
            save_cv,
            cv_print,
            cv_preview,
            get_config,
            set_config
        ])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
