// 代码生成时间: 2025-08-09 03:45:14
package main

import (
    "flag"
    "fmt"
    "image"
    "image/jpeg"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"

    "github.com/nfnt/resize"
    "github.com/spf13/cobra"
)

// ImageResizerCmd is the cobra command for resizing images
var ImageResizerCmd = &cobra.Command{
    Use:   "image-resizer",
    Short: "Resizes images to a specified width and height",
    Long:  `Resizes images to a new width and height.`,
    Run:   resizeImages,
}

// flags
var widthFlag int
var heightFlag int
var srcDirFlag string
var destDirFlag string

func init() {
    ImageResizerCmd.Flags().IntVarP(&widthFlag, "width", "w", 0, "Width of the resized images")
    ImageResizerCmd.Flags().IntVarP(&heightFlag, "height", "h", 0, "Height of the resized images")
    ImageResizerCmd.Flags().StringVarP(&srcDirFlag, "source", "s", "./", "Source directory of images to resize")
    ImageResizerCmd.Flags().StringVarP(&destDirFlag, "destination", "d", "./", "Destination directory for resized images")
}

// resizeImages is the function that performs the image resizing
func resizeImages(cmd *cobra.Command, args []string) {
    if widthFlag <= 0 || heightFlag <= 0 {
        fmt.Println("Error: Width and height must be greater than 0")
        os.Exit(1)
    }

    // Check if the source directory exists
    if _, err := os.Stat(srcDirFlag); os.IsNotExist(err) {
        fmt.Printf("Error: Source directory %s does not exist
", srcDirFlag)
        os.Exit(1)
    }

    // Check if the destination directory exists, create if not
    if _, err := os.Stat(destDirFlag); os.IsNotExist(err) {
        if err := os.MkdirAll(destDirFlag, 0755); err != nil {
            fmt.Printf("Error: Unable to create destination directory %s
", destDirFlag)
            os.Exit(1)
        }
    }

    // Walk through the source directory
    err := filepath.Walk(srcDirFlag, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.IsDir() {
            return nil
        }

        // Check if the file is an image (only JPEG for simplicity)
        if strings.HasSuffix(info.Name(), ".jpg") || strings.HasSuffix(info.Name(), ".jpeg") {
            resizeImage(path)
        }
        return nil
    })
    if err != nil {
        fmt.Printf("Error walking through source directory: %v
", err)
        os.Exit(1)
    }
}

// resizeImage resizes a single image file
func resizeImage(path string) {
    // Open the image file
    file, err := os.Open(path)
    if err != nil {
        fmt.Printf("Error opening image file %s: %v
", path, err)
        return
    }
    defer file.Close()

    // Decode the image
    img, _, err := image.Decode(file)
    if err != nil {
        fmt.Printf("Error decoding image %s: %v
", path, err)
        return
    }

    // Resize the image
    img = resize.Resize(uint(widthFlag), uint(heightFlag), img, resize.Lanczos3)

    // Create the destination path
    destPath := filepath.Join(destDirFlag, filepath.Base(path))

    // Create the image file in the destination directory
    outFile, err := os.Create(destPath)
    if err != nil {
        fmt.Printf("Error creating resized image file %s: %v
", destPath, err)
        return
    }
    defer outFile.Close()

    // Save the resized image to the destination directory
    if err := jpeg.Encode(outFile, img, nil); err != nil {
        fmt.Printf("Error saving resized image file %s: %v
", destPath, err)
        return
    }
    fmt.Printf("Resized image saved to %s
", destPath)
}

func main() {
    cmd := &cobra.Command{
        Use:    "image-resizer",
        Short:  "An image resizer tool",
        Long:   `A utility tool for resizing images in a directory.`,
        Args:   cobra.MinimumNArgs(0),
        Run:    ImageResizerCmd.Run,
    }

    cmd.Execute()
}