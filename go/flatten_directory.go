// flatten_directory.go
package main

import (
    "fmt"
    "io"
    "os"
    "path/filepath"
)

func flattenDirectory(sourceDir, targetDir string) {
    if _, err := os.Stat(sourceDir); os.IsNotExist(err) {
        fmt.Println("Error: Source directory does not exist.")
        os.Exit(1)
    }

    os.MkdirAll(targetDir, os.ModePerm)

    filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            targetPath := filepath.Join(targetDir, info.Name())
            copyFile(path, targetPath)
        }
        return nil
    })
    fmt.Printf("All files have been extracted to %s without preserving folder structure.\n", targetDir)
}

func copyFile(source, target string) {
    sourceFile, err := os.Open(source)
    if err != nil {
        fmt.Printf("Error opening source file %s: %v\n", source, err)
        return
    }
    defer sourceFile.Close()

    targetFile, err := os.Create(target)
    if err != nil {
        fmt.Printf("Error creating target file %s: %v\n", target, err)
        return
    }
    defer targetFile.Close()

    _, err = io.Copy(targetFile, sourceFile)
    if err != nil {
        fmt.Printf("Error copying file %s to %s: %v\n", source, target, err)
    }
}

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage: flatten_directory <source_directory> <target_directory>")
        os.Exit(1)
    }
    flattenDirectory(os.Args[1], os.Args[2])
}
