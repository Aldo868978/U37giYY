// 代码生成时间: 2025-08-14 10:36:14
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

// 定义版本号常量
const Version string = "1.0.0"

// 定义重命名的参数结构体
type RenameOptions struct {
    Path   string
    Prefix string
}

// 初始化rootCmd
var rootCmd = &cobra.Command{
    Use:   "batch-rename [options]", // 命令的使用说明
    Short: "A batch file rename tool using Cobra.",
    Long:  `This tool can rename files in batch with a specified prefix.`,
    Run:   handleRename,
}

// 初始化RenameOptions
var opts RenameOptions

// init函数初始化cobra命令行参数
func init() {
    cobra.OnInitialize() // 初始化cobra
    rootCmd.Flags().StringVarP(&opts.Path, "path", "p", ".", "Directory path to rename files in") // 设置--path参数，默认当前目录
    rootCmd.Flags().StringVarP(&opts.Prefix, "prefix", "r", "", "Prefix to add to each file name") // 设置--prefix参数
}

// handleRename处理文件重命名
func handleRename(cmd *cobra.Command, args []string) {
    path := opts.Path
    if !strings.HasSuffix(path, string(os.PathSeparator)) {
        path += string(os.PathSeparator)
    }
    
    // 获取指定目录下的所有文件
    files, err := os.ReadDir(path)
    if err != nil {
        log.Fatalf("Failed to read directory: %v", err)
    }

    for _, file := range files {
        if !file.IsDir() { // 只处理文件
            newName := fmt.Sprintf("%s%s", opts.Prefix, file.Name()) // 添加前缀到文件名
            oldPath := filepath.Join(path, file.Name())
            newPath := filepath.Join(path, newName)
            
            // 重命名文件
            if err := os.Rename(oldPath, newPath); err != nil {
                log.Printf("Failed to rename file %s to %s: %v", oldPath, newPath, err)
            } else {
                fmt.Printf("Renamed %s to %s", oldPath, newPath)
            }
        }
    }
}

// main函数是程序的入口点
func main() {
    rootCmd.Version = Version // 设置cobra的版本号
    if err := rootCmd.Execute(); err != nil {
        log.Fatal(err)
    }
}
