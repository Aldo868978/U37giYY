// 代码生成时间: 2025-09-01 19:46:34
package main

import (
    "flag"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "time"
    "strings"
)

// Version is the version of the backup tool
var Version string

func main() {
    // Create a new instance of the Cobra command
    var rootCmd = &cobra.Command{
        Use:   "file-backup-sync",
        Short: "A tool for file backup and synchronization",
        Long:  `A command line tool for backing up and synchronizing files across directories.`,
    }

    // Define flags
    sourcePath := flag.StringP("source", "s", "", "The source directory path to backup")
    destPath := flag.StringP("destination\, "d", "", "The destination directory path for synchronization")
    version := flag.BoolP("version", "v", false, "Print the version of the file-backup-sync tool")

    // Bind flags to the root command
    rootCmd.Flags().StringVarP(sourcePath, "source", "s", *sourcePath, "The source directory path to backup")
    rootCmd.Flags().StringVarP(destPath, "destination\, "d", *destPath, "The destination directory path for synchronization")
    rootCmd.Flags().BoolVar(version, "version", *version, "Print the version of the file-backup-sync tool")

    // Execute the command
    err := rootCmd.Execute()
    if err != nil {
        log.Fatalf("Failed to execute command: %v", err)
    }
}

// backupSourceToDestination copies the source directory to the destination directory
func backupSourceToDestination(sourcePath, destPath string) error {
    // Ensure source path exists
    if _, err := os.Stat(sourcePath); os.IsNotExist(err) {
        return fmt.Errorf("source path does not exist: %s", sourcePath)
    }

    // Create destination directory if it doesn't exist
    if _, err := os.Stat(destPath); os.IsNotExist(err) {
        err := os.MkdirAll(destPath, 0755)
        if err != nil {
            return fmt.Errorf("failed to create destination directory: %v", err)
        }
    }

    // Walk through the source path
    err := filepath.Walk(sourcePath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // Construct the relative path
        relativePath, err := filepath.Rel(sourcePath, path)
        if err != nil {
            return err
        }

        // Create the destination file path
        destFilePath := filepath.Join(destPath, relativePath)

        // Handle directories
        if info.IsDir() {
            // Create subdirectory if it doesn't exist
            if _, err := os.Stat(destFilePath); os.IsNotExist(err) {
                err := os.MkdirAll(destFilePath, 0755)
                if err != nil {
                    return fmt.Errorf("failed to create subdirectory: %v", err)
                }
            }
            return nil
        }

        // Handle files
        srcFile, err := os.Open(path)
        if err != nil {
            return fmt.Errorf("failed to open source file: %v", err)
        }
        defer srcFile.Close()

        destFile, err := os.Create(destFilePath)
        if err != nil {
            return fmt.Errorf("failed to create destination file: %v", err)
        }
        defer destFile.Close()

        _, err = io.Copy(destFile, srcFile)
        if err != nil {
            return fmt.Errorf("failed to copy file: %v", err)
        }

        // Preserve file permissions
        err = os.Chmod(destFilePath, info.Mode())
        if err != nil {
            return fmt.Errorf("failed to set file permissions: %v", err)
        }

        return nil
    })

    return err
}

// getVersion prints the version of the backup tool
func getVersion() {
    if Version == "" {
        fmt.Println("Version not set")
    } else {
        fmt.Printf("file-backup-sync version: %s\
", Version)
    }
}
