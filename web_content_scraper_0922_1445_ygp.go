// 代码生成时间: 2025-09-22 14:45:28
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
    "time"

    "github.com/spf13/cobra"
)

// 定义命令
var rootCmd = &cobra.Command{
    Use:   "webContentScraper",
    Short: "A tool to scrape web content",
    Long:  "webContentScraper is a CLI tool for scraping web content",
    Run:   scrapeWebContent,
}

// scrapeWebContent 执行网页内容抓取操作
func scrapeWebContent(cmd *cobra.Command, args []string) {
    if len(args) == 0 {
        fmt.Println("Please provide a URL to scrape")
        return
    }

    url := args[0]
    resp, err := http.Get(url)
    if err != nil {
        fmt.Printf("Error fetching URL: %s
", err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Printf("Failed to fetch URL: %d
", resp.StatusCode)
        return
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("Error reading response body: %s
", err)
        return
    }

    // 简单处理，移除HTML标签
    content := strings.ReplaceAll(string(body), "<[^>]*>", "")
    fmt.Println(content)
}

// initCobra 初始化Cobra命令
func initCobra() {
    rootCmd.PersistentFlags().StringP("url", "u", "", "URL to scrape")
    rootCmd.MarkFlagRequired("url")
}

func main() {
    initCobra()
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
    time.Sleep(100 * time.Millisecond) // 等待命令行输出
    }
}