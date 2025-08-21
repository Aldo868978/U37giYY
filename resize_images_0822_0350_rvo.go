// 代码生成时间: 2025-08-22 03:50:37
package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "image"
    "image/jpeg"
    "image/png"
    "os"
    "path/filepath"
    "strings"
    "github.com/nfnt/resize"
)

const usage = `批量调整图片尺寸
用法:
  resize_images -source=<源文件夹> -destination=<目标文件夹> -width=<宽度> -height=<高度>
`

type Options struct {
    Source string
    Destination string
    Width int
    Height int
}

func main() {
    var opts Options
    flag.StringVar(&opts.Source, "source", "", "源文件夹路径")
    flag.StringVar(&opts.Destination, "destination", "", "目标文件夹路径")
    flag.IntVar(&opts.Width, "width", 0, "目标宽度")
    flag.IntVar(&opts.Height, "height", 0, "目标高度")
    flag.Parse()

    if opts.Source == "" || opts.Destination == "" || opts.Width <= 0 || opts.Height <= 0 {
        fmt.Println("参数不足")
        fmt.Println(usage)
        return
    }

    if err := ResizeImages(opts); err != nil {
        fmt.Printf("批量调整图片尺寸失败：%v
", err)
    }
}

// ResizeImages 批量调整图片尺寸
func ResizeImages(opts Options) error {
    files, err := ioutil.ReadDir(opts.Source)
    if err != nil {
        return err
    }

    for _, file := range files {
        if !file.IsDir() {
            srcPath := filepath.Join(opts.Source, file.Name())
            dstPath := filepath.Join(opts.Destination, file.Name())
            if err := ResizeImage(srcPath, dstPath, opts.Width, opts.Height); err != nil {
                return err
            }
        }
    }
    return nil
}

// ResizeImage 调整单张图片的尺寸
func ResizeImage(srcPath, dstPath string, width, height int) error {
    srcImage, err := loadImage(srcPath)
    if err != nil {
        return err
    }
    dstImage := resize.Resize(uint(width), uint(height), srcImage, resize.NearestNeighbor)
    if err := saveImage(dstPath, dstImage); err != nil {
        return err
    }
    return nil
}

// loadImage 加载图片
func loadImage(path string) (image.Image, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    image, _, err := image.Decode(file)
    if err != nil {
        return nil, err
    }
    return image, nil
}

// saveImage 保存图片
func saveImage(path string, image image.Image) error {
    ext := strings.ToLower(filepath.Ext(path))
    var encoder image.Encoder
    switch ext {
    case ".jpg", ".jpeg":
        encoder = jpeg.Encoder{
            Quality: 90,
        }
    case ".png":
        encoder = png.Encoder{
            // 可以添加更多选项
        }
    default:
        return fmt.Errorf("不支持的图片格式: %s", ext)
    }

    file, err := os.Create(path)
    if err != nil {
        return err
}
    defer file.Close()

    return encoder.Encode(file, image)
}
