use tauri::api::path;
use tauri::Window;
use tauri::api::process::{Command, CommandEvent};

// TODO:i think language should be loaded from db.[next step]
#[tauri::command]
pub async fn cv_print(window: Window, version: String, language: String) {
     println!("{}", language);
    let file_path = path::app_data_dir(&Default::default())
        .ok_or("failed to get app dir".to_string())
        .unwrap()
        .join("CVKawa")
        .join(version)
        .join("cv_data.json");
    let json_path = match file_path.to_str().ok_or("failed to get json path".to_string()) {
        Ok(p) => p.to_string(),
        Err(e) => {
            eprintln!("failed to get json path: {}", e);
            return;
        }
    };
    let desktop_path = match path::desktop_dir().unwrap().join("cv.pdf").to_str().ok_or("failed to get desktop path".to_string()) {
        Ok(p) => p.to_string(),
        Err(e) => {
            eprintln!("failed to get desktop path: {}", e);
            return;
        }
    };
    let (mut rx, mut child) = Command::new_sidecar("jtp")
        .expect("failed to create `my-sidecar` binary command")
        .args(&["-f", &json_path, "-o", &desktop_path, "-l", &language])
        .spawn()
        .expect("Failed to spawn sidecar");

    tauri::async_runtime::spawn(async move {
        // read events such as stdout
        while let Some(event) = rx.recv().await {
            if let CommandEvent::Stdout(line) = event {
                window
                    .emit("message", Some(format!("'{}'", line)))
                    .expect("failed to emit event");
                // write to stdin
                child.write("message from Rust\n".as_bytes()).unwrap();
            }
        }
    });
}

#[tauri::command]
pub async fn cv_preview(handle: tauri::AppHandle, version: String) {
    let label = format!("Viewer-{}", version);
    let url = format!("/#/print_view?version={}", version);
    let doc_page = tauri::WindowBuilder::new(
        &handle,
        label, /* the unique window label */
        tauri::WindowUrl::App(url.into()),
    );
    doc_page.inner_size(595.0, 842.0).title("Preview").center().build().unwrap();
}