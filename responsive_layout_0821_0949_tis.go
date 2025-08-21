// 代码生成时间: 2025-08-21 09:49:34
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"

    "github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
    Use:   "responsive-layout",
    Short: "A brief description of your application",
    Long: `A longer description that spans multiple lines and likely contains
examples of using your application. For example:
Cobra is a CLI library for Go that enables creation of a Cobra
command without having to handle all of the work required for
arg parsing.`,
    Run: func(cmd *cobra.Command, args []string) {
        serveHttp()
    },
}

// serveHttp sets up the HTTP server with routes for serving HTML and CSS files
func serveHttp() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        serveLayout(w, r)
    })
    
    port := "8080"
    fmt.Printf("Starting server on :%s
", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatalf("Error starting server: %s", err)
    }
}

// serveLayout serves the layout HTML file with CSS for responsive design
func serveLayout(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    
    // Check for different device types or screen sizes to adjust layout
    isMobile := strings.Contains(r.UserAgent(), "Mobile") || strings.Contains(r.UserAgent(), "Android") || strings.Contains(r.UserAgent(), "iPhone")
    if isMobile {
        http.ServeFile(w, r, "mobile_layout.html")
    } else {
        http.ServeFile(w, r, "desktop_layout.html")
    }
}

func main() {
    if err := RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
