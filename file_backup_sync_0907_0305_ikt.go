// 代码生成时间: 2025-09-07 03:05:33
package main

import (
    "crypto/sha1"
    "encoding/hex"
    "flag"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"
    "sync"
)

// FileBackupSyncTool 文件备份和同步工具
type FileBackupSyncTool struct {
    SourceDir  string
    DestinationDir string
    Concurrency int
}

// NewFileBackupSyncTool 创建一个新的文件备份和同步工具实例
func NewFileBackupSyncTool(sourceDir, destinationDir string, concurrency int) *FileBackupSyncTool {
    return &FileBackupSyncTool{
        SourceDir:  sourceDir,
        DestinationDir: destinationDir,
        Concurrency: concurrency,
    }
}

// BackupAndSync 执行文件备份和同步操作
func (tool *FileBackupSyncTool) BackupAndSync() error {
    sourceFiles, err := ioutil.ReadDir(tool.SourceDir)
    if err != nil {
        return fmt.Errorf("failed to read source directory: %w", err)
    }

    var wg sync.WaitGroup
    errCh := make(chan error, tool.Concurrency)

    for _, file := range sourceFiles {
        if file.IsDir() {
            continue
        }

        wg.Add(1)
        go func(f os.FileInfo) {
            defer wg.Done()

            sourcePath := filepath.Join(tool.SourceDir, f.Name())
            destinationPath := filepath.Join(tool.DestinationDir, f.Name())

            // 计算文件哈希值
            sourceHash, err := calculateFileHash(sourcePath)
            if err != nil {
                errCh <- fmt.Errorf("failed to calculate hash for %s: %w", sourcePath, err)
                return
            }

            // 检查目标目录中是否有相同哈希值的文件
            destHash, err := calculateFileHash(destinationPath)
            if os.IsNotExist(err) || destHash != sourceHash {
                // 复制文件
                if err := copyFile(sourcePath, destinationPath); err != nil {
                    errCh <- fmt.Errorf("failed to copy file from %s to %s: %w", sourcePath, destinationPath, err)
                    return
                }
            }
        }(file)
    }

    wg.Wait()
    close(errCh)

    // 检查是否有错误发生
    for err := range errCh {
        if err != nil {
            return err
        }
    }

    return nil
}

// calculateFileHash 计算文件的SHA1哈希值
func calculateFileHash(filePath string) (string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return "", err
    }
    defer file.Close()

    hash := sha1.New()
    if _, err := io.Copy(hash, file); err != nil {
        return "", err
    }
    return hex.EncodeToString(hash.Sum(nil)), nil
}

// copyFile 复制文件
func copyFile(sourcePath, destinationPath string) error {
    src, err := os.Open(sourcePath)
    if err != nil {
        return err
    }
    defer src.Close()

    dst, err := os.Create(destinationPath)
    if err != nil {
        return err
    }
    defer dst.Close()

    if _, err := io.Copy(dst, src); err != nil {
        return err
    }
    return nil
}

func main() {
    var sourceDir, destinationDir string
    var concurrency int

    flag.StringVar(&sourceDir, "source", "", "source directory")
    flag.StringVar(&destinationDir, "destination", "", "destination directory")
    flag.IntVar(&concurrency, "concurrency", 10, "concurrency level")
    flag.Parse()

    if sourceDir == "" || destinationDir == "" {
        log.Fatalf("source and destination directories must be provided")
    }

    tool := NewFileBackupSyncTool(sourceDir, destinationDir, concurrency)
    if err := tool.BackupAndSync(); err != nil {
        log.Fatalf("backup and sync failed: %v", err)
    }
    fmt.Println("Backup and sync completed successfully")
}
