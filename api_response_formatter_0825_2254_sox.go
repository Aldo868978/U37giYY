// 代码生成时间: 2025-08-25 22:54:13
package main

import (
    "encoding/json"
    "fmt"
    "log"

    "github.com/spf13/cobra"
)

// ApiResponse represents the structure of an API response
type ApiResponse struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}

// formatApiResponse formats the given data into an API response
func formatApiResponse(status, message string, data interface{}) (string, error) {
    response := ApiResponse{
        Status:  status,
        Message: message,
        Data:    data,
    }

    jsonResponse, err := json.MarshalIndent(response, "", "  ")
    if err != nil {
        return "", fmt.Errorf("error formatting API response: %w", err)
    }

    return string(jsonResponse), nil
}

func main() {
    var rootCmd = &cobra.Command{
        Use:   "api-response-formatter",
        Short: "A tool to format API responses",
        Long:  `A simple CLI tool to format data into an API response format.`,
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Welcome to the API Response Formatter!")
        },
    }

    // Define a sub-command to format API responses
    var formatCmd = &cobra.Command{
        Use:   "format",
        Short: "Format data into an API response",
        Long:  `This command formats data into an API response format.`,
        Run: func(cmd *cobra.Command, args []string) {
            if len(args) < 3 {
                fmt.Println("Please provide at least three arguments: status, message, and data.")
                return
            }

            status := args[0]
            message := args[1]
            data := args[2]

            formattedResponse, err := formatApiResponse(status, message, data)
            if err != nil {
                log.Fatalf("Error formatting API response: %s", err)
            }

            fmt.Println(formattedResponse)
        },
    }

    rootCmd.AddCommand(formatCmd)

    err := rootCmd.Execute()
    if err != nil {
        log.Fatalf("Error executing command: %s", err)
    }
}
