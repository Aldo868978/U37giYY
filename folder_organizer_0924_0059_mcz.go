// 代码生成时间: 2025-09-24 00:59:30
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "log"
    "flag"

    "github.com/spf13/cobra"
)

// Define constants for folder types
const (
    Audio = "audio"
    Video = "video"
    Image = "image"
    Document = "document"
)

// FolderOrganizer represents the application state
type FolderOrganizer struct {
    srcPath string
    folderMap map[string]string
}

// NewFolderOrganizer creates a new FolderOrganizer instance
func NewFolderOrganizer(srcPath string) *FolderOrganizer {
    return &FolderOrganizer{
        srcPath: srcPath,
        folderMap: map[string]string{
            Audio:     "audio",
            Video:     "video",
            Image:     "image",
            Document:  "document",
        },
    }
}

// Organize folders based on file extensions
func (f *FolderOrganizer) Organize() error {
    // Read the source directory
    files, err := os.ReadDir(f.srcPath)
    if err != nil {
        return fmt.Errorf("failed to read source directory: %w", err)
    }

    for _, file := range files {
        if !file.IsDir() {
            // Get the file extension
            ext := strings.TrimPrefix(filepath.Ext(file.Name()), ".")
            // Map the extension to the corresponding folder
            folder, ok := f.folderMap[ext]
            if !ok {
                continue
            }
            // Create the destination path
            dstPath := filepath.Join(f.srcPath, folder, file.Name())
            // Move the file to the new location
            if err := os.Rename(filepath.Join(f.srcPath, file.Name()), dstPath); err != nil {
                return fmt.Errorf("failed to move file %s: %w", file.Name(), err)
            }
        }
    }
    return nil
}

// main function to execute the program
func main() {
    var srcPath string
    flag.StringVarP(&srcPath, "source", "s", ".", "source directory path")
    flag.Parse()

    // Create a new FolderOrganizer instance
    organizer := NewFolderOrganizer(srcPath)

    // Check if the source directory exists
    if _, err := os.Stat(srcPath); os.IsNotExist(err) {
        log.Fatalf("source directory does not exist: %s", srcPath)
    }

    // Organize the folders
    if err := organizer.Organize(); err != nil {
        log.Fatalf("error organizing folders: %s", err)
    }
    fmt.Println("Folders organized successfully")
}
