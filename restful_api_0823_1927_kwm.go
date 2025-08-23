// 代码生成时间: 2025-08-23 19:27:37
package main

import (
    "encoding/json"
    "net/http"
    "github.com/spf13/cobra"
)

// 定义 RESTful API 的结构体，包含 Cobra 的根命令
var rootCmd = &cobra.Command{
    Use:   "restful-api",
    Short: "A brief description of your application",
    Long:  `A longer description that spans multiple lines and likely contains
examples of how to use the application.`,
    Run: func(cmd *cobra.Command, args []string) {
        // 启动 HTTP 服务器
        startServer()
    },
}

// 定义 HTTP 路由及处理函数
func startServer() {
    http.HandleFunc("/api/hello", helloHandler)
    http.ListenAndServe(":8080", nil)
}

// helloHandler 处理 GET 请求，返回一个简单的 Hello World 响应
func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
        return
    }

    response := map[string]string{
        "message": "Hello, World!",
    }
    json.NewEncoder(w).Encode(response)
}

func main() {
    // 配置 Cobra 命令
    rootCmd.Example = `  # 启动 HTTP 服务器
  restful-api
  # 通过浏览器或 curl 发送 GET 请求到 http://localhost:8080/api/hello`
    rootCmd.AddCommand(versionCmd)
    if err := rootCmd.Execute(); err != nil {
        panic(err)
    }
}

// 定义版本命令
var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "Print the version number of the application",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("restful-api version 1.0")
    },
}
