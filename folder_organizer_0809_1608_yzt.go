// 代码生成时间: 2025-08-09 16:08:26
package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"

    "github.com/spf13/cobra"
)

// FolderOrganizer represents the folder organizer application
type FolderOrganizer struct {
    RootDirectory string
    Year         int
}

// NewFolderOrganizer creates a new FolderOrganizer instance
func NewFolderOrganizer(rootDirectory string, year int) *FolderOrganizer {
    return &FolderOrganizer{
        RootDirectory: rootDirectory,
        Year:         year,
    }
}

// Organize performs the folder organization based on the provided structure
func (f *FolderOrganizer) Organize() error {
    // Ensure the root directory exists
    if _, err := os.Stat(f.RootDirectory); os.IsNotExist(err) {
        return fmt.Errorf("root directory does not exist: %w", err)
    }

    // Create a new directory structure based on the current year
    newDir := fmt.Sprintf("%s/%d", f.RootDirectory, f.Year)
    if err := os.MkdirAll(newDir, 0755); err != nil {
        return fmt.Errorf("failed to create directory: %w", err)
    }

    // List all files and subdirectories in the root directory
    files, err := os.ReadDir(f.RootDirectory)
    if err != nil {
        return fmt.Errorf("failed to read root directory: %w", err)
    }

    for _, file := range files {
        srcPath := filepath.Join(f.RootDirectory, file.Name())
        destPath := filepath.Join(newDir, file.Name())

        // Skip directories and only move files
        if file.IsDir() {
            continue
        }

        // Move the file to the new directory
        if err := os.Rename(srcPath, destPath); err != nil {
            return fmt.Errorf("failed to move file %s: %w", file.Name(), err)
        }
    }

    return nil
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
    Use:   "folder-organizer",
    Short: "Folder Organizer CLI tool",
    Long: `Folder Organizer CLI tool helps to organize files in folders based on the year.`,
}

// Execute adds all child commands to the root command sets flags appropriately.
func Execute() {
    if err := RootCmd.Execute(); err != nil {
        log.Fatal(err)
    }
}

func main() {
    // Set up the root directory and year flags
    rootDir := RootCmd.PersistentFlags().StringP("root-dir", "r", ".", "The root directory to organize files")
    year := RootCmd.PersistentFlags().IntP("year", "y", time.Now().Year(), "The year to organize files into")

    // Add the root directory and year flags to the root command
    RootCmd.MarkPersistentFlagRequired("root-dir")
    RootCmd.MarkPersistentFlagRequired("year")

    // Run the folder organizer when the command is executed
    RootCmd.RunE = func(cmd *cobra.Command, args []string) error {
        fo := NewFolderOrganizer(*rootDir, *year)
        return fo.Organize()
    }

    Execute()
}