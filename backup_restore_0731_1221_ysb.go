// 代码生成时间: 2025-07-31 12:21:52
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "log"
    "archive/zip"
    "io"
    "io/ioutil"
)

// BackupService provides methods for backing up and restoring data.
type BackupService struct {
    // Define any necessary fields for the backup service
}

// NewBackupService creates a new instance of BackupService.
func NewBackupService() *BackupService {
    return &BackupService{}
}

// Backup creates a backup of the specified directory and saves it to a zip file.
func (s *BackupService) Backup(srcDir, backupFile string) error {
    // Create a buffer to write our archive to.
    buf := new(bytes.Buffer)
    
    // Create a zip archive.
    zw := zip.NewWriter(buf)
    defer zw.Close()
    
    // Walk through the source directory.
    err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        
        // Create a file to store in the zip.
        f, err := zw.Create(path)
        if err != nil {
            return err
        }
        
        // If it's a file, copy its contents.
        if !info.IsDir() {
            file, err := os.Open(path)
            if err != nil {
                return err
            }
            defer file.Close()
            
            _, err = io.Copy(f, file)
            return err
        }
        
        return nil
    })
    
    if err != nil {
        return err
    }
    
    // Save the zip file to the backup location.
    return ioutil.WriteFile(backupFile, buf.Bytes(), 0644)
}

// Restore restores data from a zip file to a specified directory.
func (s *BackupService) Restore(backupFile, dstDir string) error {
    // Open the zip archive for reading.
    reader, err := zip.OpenReader(backupFile)
    if err != nil {
        return err
    }
    defer reader.Close()
    
    // Iterate through the files in the archive.
    for _, file := range reader.File {
        // Create the directory structure.
        target := filepath.Join(dstDir, file.Name)
        if file.FileInfo().IsDir() {
            os.MkdirAll(target, os.ModePerm)
            continue
        }
        
        // Create the target directory structure.
        if err := os.MkdirAll(filepath.Dir(target), os.ModePerm); err != nil {
            return err
        }
        
        // Write the file contents to the destination.
        rc, err := file.Open()
        if err != nil {
            return err
        }
        defer rc.Close()
        
        outFile, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
        if err != nil {
            return err
        }
        defer outFile.Close()
        
        _, err = io.Copy(outFile, rc)
        if err != nil {
            return err
        }
    }
    return nil
}

func main() {
    // Create a new backup service instance.
    service := NewBackupService()
    
    // Define the source directory and backup file location.
    srcDir := "./data"
    backupFile := "./data_backup.zip"
    dstDir := "./restored_data"
    
    // Backup data.
    if err := service.Backup(srcDir, backupFile); err != nil {
        log.Fatalf("Failed to backup data: %s", err)
    } else {
        fmt.Println("Backup successful.")
    }
    
    // Restore data.
    if err := service.Restore(backupFile, dstDir); err != nil {
        log.Fatalf("Failed to restore data: %s", err)
    } else {
        fmt.Println("Restore successful.")
    }
}
