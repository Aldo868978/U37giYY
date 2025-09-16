// 代码生成时间: 2025-09-16 14:48:21
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "log"
    "github.com/spf13/cobra"
)

// Define the root command
var rootCmd = &cobra.Command{
    Use:   "folder-organizer",
    Short: "Folder Organizer helps to organize files within a specified directory.",
    Long:  `Folder Organizer is a utility for organizing files within a specified directory.
It can classify files based on their extensions and move them into subdirectories accordingly.`,
    Run: func(cmd *cobra.Command, args []string) {
        organize(cmd, args)
    },
}

// Define flags
var targetDir string
var organizeByExt bool

func init() {
    rootCmd.PersistentFlags().StringVarP(&targetDir, "directory", "d", ".", "Target directory to organize")
    rootCmd.PersistentFlags().BoolVarP(&organizeByExt, "by-extension", "e", false, "Organize files by extension")
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

// organize function to handle the logic of organizing files
func organize(cmd *cobra.Command, args []string) {
    if _, err := os.Stat(targetDir); os.IsNotExist(err) {
        log.Fatalf("Target directory does not exist: %s", targetDir)
    }

    if organizeByExt {
        organizeByExtension(targetDir)
    } else {
        fmt.Println("No organization method specified. Use --by-extension to organize by file extension.")
    }
}

// organizeByExtension function to organize files by their extensions
func organizeByExtension(directory string) {
    // Read all files and directories in the target directory
    files, err := os.ReadDir(directory)
    if err != nil {
        log.Fatalf("Error reading directory: %s", err)
    }

    for _, file := range files {
        if !file.IsDir() {
            ext := strings.TrimPrefix(filepath.Ext(file.Name()), ".")
            if ext == "" {
                continue // Skip files without an extension
            }

            // Create a subdirectory if it doesn't exist
            dirPath := filepath.Join(directory, ext)
            if _, err := os.Stat(dirPath); os.IsNotExist(err) {
                err := os.MkdirAll(dirPath, 0755)
                if err != nil {
                    log.Fatalf("Failed to create directory: %s", err)
                }
            }

            // Move the file to the corresponding subdirectory
            srcPath := filepath.Join(directory, file.Name())
            destPath := filepath.Join(dirPath, file.Name())
            if err := os.Rename(srcPath, destPath); err != nil {
                log.Fatalf("Failed to move file: %s", err)
            }
        }
    }
}
