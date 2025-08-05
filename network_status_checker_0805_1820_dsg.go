// 代码生成时间: 2025-08-05 18:20:53
package main

import (
    "fmt"
    "net"
    "os"
    "time"

    "github.com/spf13/cobra"
)

// NetworkStatusChecker defines a structure to hold command line arguments
type NetworkStatusChecker struct {
    Host string
    Port int
}

// NewNetworkStatusChecker creates a new NetworkStatusChecker instance
func NewNetworkStatusChecker(host string, port int) *NetworkStatusChecker {
    return &NetworkStatusChecker{
        Host: host,
        Port: port,
    }
}

// CheckStatus checks the network connection status to the provided host and port
func (n *NetworkStatusChecker) CheckStatus() error {
    address := fmt.Sprintf("%s:%d", n.Host, n.Port)
    conn, err := net.DialTimeout("tcp", address, 5*time.Second)
    if err != nil {
        return fmt.Errorf("failed to connect to %s: %w", address, err)
    }
    defer conn.Close()
    return nil
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "network-status-checker",
    Short: "Check network connection status",
    Long:  `Check network connection status to a specified host and port`,
    RunE: func(cmd *cobra.Command, args []string) error {
        if len(args) != 2 {
            return fmt.Errorf("requires two arguments: host and port")
        }
        host := args[0]
        port, err := strconv.Atoi(args[1])
        if err != nil {
            return fmt.Errorf("invalid port: %w", err)
        }

        checker := NewNetworkStatusChecker(host, port)
        if err := checker.CheckStatus(); err != nil {
            return fmt.Errorf("network status check failed: %w", err)
        }
        fmt.Println("Network connection is established")
        return nil
    },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main().
// It only needs to happen once to the rootCmd.
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func main() {
    Execute()
}
