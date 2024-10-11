// flatten_directory.rs
use std::fs;
use std::path::Path;
use std::process;

fn flatten_directory(source_dir: &str, target_dir: &str) {
    if !Path::new(source_dir).exists() {
        eprintln!("Error: Source directory does not exist.");
        process::exit(1);
    }

    fs::create_dir_all(target_dir).unwrap();

    for entry in fs::read_dir(source_dir).unwrap() {
        let entry = entry.unwrap();
        let path = entry.path();
        if path.is_file() {
            let target_path = Path::new(target_dir).join(entry.file_name());
            fs::copy(&path, &target_path).unwrap();
        } else if path.is_dir() {
            flatten_directory(path.to_str().unwrap(), target_dir);
        }
    }
}

fn main() {
    let args: Vec<String> = std::env::args().collect();
    if args.len() != 3 {
        eprintln!("Usage: flatten_directory <source_directory> <target_directory>");
        process::exit(1);
    }
    flatten_directory(&args[1], &args[2]);
    println!("All files have been extracted to {} without preserving folder structure.", &args[2]);
}
