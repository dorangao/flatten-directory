# flatten_directory.py
import os
import shutil
import sys

def flatten_directory(source_dir: str, target_dir: str):
    if not os.path.exists(source_dir):
        print("Error: Source directory does not exist.")
        sys.exit(1)

    os.makedirs(target_dir, exist_ok=True)

    for root, _, files in os.walk(source_dir):
        for file in files:
            shutil.copy2(os.path.join(root, file), os.path.join(target_dir, file))

    print(f"All files have been extracted to {target_dir} without preserving folder structure.")

if __name__ == "__main__":
    if len(sys.argv) != 3:
        print("Usage: python flatten_directory.py <source_directory> <target_directory>")
        sys.exit(1)

    flatten_directory(sys.argv[1], sys.argv[2])
