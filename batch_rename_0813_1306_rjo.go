// 代码生成时间: 2025-08-13 13:06:30
package main

import (
    "errors"
    "fmt"
    "log"
    "os"
    "path/filepath"

    "github.com/spf13/cobra"
)

// 定义全局变量
var sourceDir string
var targetDir string
var prefix string

// rootCmd 是 cobra.Command 的根命令
var rootCmd = &cobra.Command{
    Use:   "batch-rename",
    Short: "A tool to batch rename files",
    Long:  `A tool to batch rename files in a specific directory with a given prefix.`,
    Run:   run,
}

// init 初始化命令行参数
func init() {
    rootCmd.Flags().StringVarP(&sourceDir, "source", "s", "", "source directory path")
    rootCmd.Flags().StringVarP(&targetDir, "target", "t", "", "target directory path")
    rootCmd.Flags().StringVarP(&prefix, "prefix", "p", "", "prefix for the new file names")
    rootCmd.MarkFlagRequired("source")
    rootCmd.MarkFlagRequired("target")
    rootCmd.MarkFlagRequired("prefix")
}

// run 执行命令行操作
func run(cmd *cobra.Command, args []string) {
    err := renameFiles(sourceDir, targetDir, prefix)
    if err != nil {
        log.Fatalf("Error during file rename: %v", err)
    }
    fmt.Println("Files have been successfully renamed.")
}

// renameFiles 重命名文件
func renameFiles(source, target, prefix string) error {
    sourceFiles, err := os.ReadDir(source)
    if err != nil {
        return fmt.Errorf("failed to read source directory: %w", err)
    }

    for _, file := range sourceFiles {
        srcPath := filepath.Join(source, file.Name())
        destPath := filepath.Join(target, prefix+file.Name())

        // 检查是否为文件
        if !file.IsDir() {
            err := os.Rename(srcPath, destPath)
            if err != nil {
                return fmt.Errorf("failed to rename file %s to %s: %w", srcPath, destPath, err)
            }
        }
    }
    return nil
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}