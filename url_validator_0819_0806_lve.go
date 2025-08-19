// 代码生成时间: 2025-08-19 08:06:00
package main

import (
    "fmt"
    "net/url"
    "strings"
    "os"
    "github.com/spf13/cobra"
)

// URLValidatorCmd is a command to validate a URL
var URLValidatorCmd = &cobra.Command{
    Use:   "validate <url>",
    Short: "Validates a given URL",
    Long:  `Validates whether a given URL is valid or not.`,
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        validateURL(args[0])
    },
}

// validateURL function checks if the given URL is valid
func validateURL(inputURL string) {
    // Parse the URL
    parsedURL, err := url.ParseRequestURI(inputURL)
    if err != nil {
        fmt.Printf("Error parsing URL: %s
", err)
        fmt.Println("Invalid URL")
        os.Exit(1)
    }

    // Check if scheme and host are not empty
    if parsedURL.Scheme == "" || parsedURL.Host == "" {
        fmt.Println("Invalid URL")
        os.Exit(1)
    }

    // Check if scheme is valid (HTTP or HTTPS)
    if !strings.HasPrefix(parsedURL.Scheme, "http") {
        fmt.Println("Invalid URL scheme. Only HTTP and HTTPS are allowed.")
        os.Exit(1)
    }

    fmt.Println("Valid URL")
}

func main() {
    // Setup Cobra command
    rootCmd := &cobra.Command{
        Use:   "urlvalidator",
        Short: "URL Validity Checker",
        Long:  `This tool checks whether a given URL is valid or not.`,
    }

    rootCmd.AddCommand(URLValidatorCmd)

    // Execute Cobra command
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}