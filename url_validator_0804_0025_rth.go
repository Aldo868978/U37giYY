// 代码生成时间: 2025-08-04 00:25:38
package main

import (
    "fmt"
    "net/http"
    "net/url"
    "os"
    "strings"

    "github.com/spf13/cobra"
)

// URLValidator is the command to validate URL
var URLValidator = &cobra.Command{
    Use:   "validate <URL>", // cobra pattern for the command
    Short: "Validates the given URL for its effectiveness",
    Long:  `Validates the given URL for its effectiveness.
  Example usage:
    go run . --url https://www.example.com`,
    Args:  cobra.ExactArgs(1), // exactly one argument required
    RunE: func(cmd *cobra.Command, args []string) error {
        u, err := url.ParseRequestURI(args[0])
        if err != nil {
            return fmt.Errorf("invalid URL: %w", err)
        }
        if u.Scheme == "" || u.Host == "" {
            return fmt.Errorf("URL must have a scheme and a host")
        }
        resp, err := http.Head(args[0])
        if err != nil {
            return fmt.Errorf("failed to reach URL: %w", err)
        }
        defer resp.Body.Close()
        if resp.StatusCode != http.StatusOK {
            return fmt.Errorf("URL is not effective, status code: %d", resp.StatusCode)
        }
        fmt.Printf("URL is valid and effective: %s
", args[0])
        return nil
    },
}

func main() {
    rootCmd := &cobra.Command{
        Use:   "url-validator",
        Short: "A brief description of your application",
        Long:  `A longer description that spans multiple lines and likely contains
  examples and usage of using your application.`,
    }

    rootCmd.AddCommand(URLValidator)

    // Execute the root command
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}