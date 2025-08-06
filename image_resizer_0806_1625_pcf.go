// 代码生成时间: 2025-08-06 16:25:06
package main

import (
    "fmt"
# NOTE: 重要实现细节
    "image"
    "image/jpeg"
# FIXME: 处理边界情况
    "image/png"
# 增强安全性
    "os"
    "path/filepath"
    "github.com/spf13/cobra"
)
# NOTE: 重要实现细节

// 定义图片尺寸调整器的配置结构体
type ResizerConfig struct {
# 扩展功能模块
    Width, Height int
    OutputFolder string
    Quality       int
}

// 定义图片尺寸调整器命令
var resizeCmd = &cobra.Command{
    Use:   "resize <path>",
# NOTE: 重要实现细节
    Short: "Resize images to specified dimensions",
    Long:  `Resize images in a directory to specified width and height.`,
    Args:  cobra.ExactArgs(1), // 需要一个参数，即图片路径
    Run:   resizeImages,
}

// main函数入口
func main() {
# TODO: 优化性能
    cobra.OnInitialize(initConfig)
    var width, height int
    var outputFolder, qualityStr string
    resizeCmd.Flags().IntVar(&width, "width", 0, "Width of the resized image")
    resizeCmd.Flags().IntVar(&height, "height", 0, "Height of the resized image")
    resizeCmd.Flags().StringVar(&outputFolder, "output", "./resized", "Output folder for resized images")
    resizeCmd.Flags().StringVar(&qualityStr, "quality", "90", "JPEG quality of the resized image")
    resizeCmd.MarkFlagRequired("width")
    resizeCmd.MarkFlagRequired("height")
    quality, _ := strconv.Atoi(qualityStr)
    err := resizeCmd.Execute()
    if err != nil {
# 扩展功能模块
        fmt.Println(err)
        os.Exit(1)
    }
}

// initConfig初始化配置
func initConfig() {
    // 在这里可以初始化配置
# NOTE: 重要实现细节
}

// resizeImages是resizeCmd的执行函数
func resizeImages(cmd *cobra.Command, args []string) {
# 改进用户体验
    path := args[0]
    config := ResizerConfig{
        Width:        cmd.Flag("width").GetInt(""),
        Height:       cmd.Flag("height").GetInt(""),
        OutputFolder: cmd.Flag("output").GetString(""),
        Quality:      cmd.Flag("quality").GetInt("90"),
    }
    err := resizeImagesInPath(path, config)
    if err != nil {
# 增强安全性
        fmt.Println("Error resizing images: ", err)
        os.Exit(1)
    }
# 添加错误处理
}

// resizeImagesInPath遍历给定路径的所有图片文件，并进行尺寸调整
func resizeImagesInPath(path string, config ResizerConfig) error {
    files, err := os.ReadDir(path)
    if err != nil {
        return fmt.Errorf("failed to list files in %s: %w", path, err)
# NOTE: 重要实现细节
    }
    for _, file := range files {
        if file.IsDir() {
# 添加错误处理
            continue
        }
        if !isImageFile(file.Name()) {
            continue
        }
        err := resizeImage(filepath.Join(path, file.Name()), config)
        if err != nil {
            return fmt.Errorf("failed to resize image %s: %w", file.Name(), err)
        }
    }
    return nil
}

// isImageFile检查文件是否为图片文件
func isImageFile(filename string) bool {
    ext := strings.ToLower(filepath.Ext(filename))
    return ext == ".jpg" || ext == ".jpeg" || ext == ".png"
}

// resizeImage调整单个图片文件的尺寸
func resizeImage(filePath string, config ResizerConfig) error {
    src, err := os.Open(filePath)
    if err != nil {
        return fmt.Errorf("failed to open image %s: %w", filePath, err)
    }
    defer src.Close()
    img, _, err := image.Decode(src)
    if err != nil {
        return fmt.Errorf("failed to decode image %s: %w", filePath, err)
    }
    
    resizedImg := resizeImageToDim(img, config.Width, config.Height)
    outputFile := filepath.Join(config.OutputFolder, filepath.Base(filePath))
    output, err := os.Create(outputFile)
    if err != nil {
        return fmt.Errorf("failed to create resized image file %s: %w", outputFile, err)
    }
    defer output.Close()
    
    if strings.HasSuffix(outputFile, ".png") {
# 添加错误处理
        err = png.Encode(output, resizedImg)
    } else {
        err = jpeg.Encode(output, resizedImg, &jpeg.Options{Quality: config.Quality})
    }
    if err != nil {
        return fmt.Errorf("failed to encode resized image %s: %w", outputFile, err)
    }
    fmt.Printf("Resized image saved to %s
", outputFile)
    return nil
# 改进用户体验
}

// resizeImageToDim调整图片尺寸到指定的宽度和高度
# FIXME: 处理边界情况
func resizeImageToDim(img image.Image, width, height int) *image.NRGBA {
    // 创建一个新的NRGBA图像，具有指定的宽度和高度
    m := image.NewNRGBA(image.Rect(0, 0, width, height))
    // 将原始图像缩放到新图像的大小
    m = resize.Resize(width, height, img, resize.Lanczos3)
    return m
}
