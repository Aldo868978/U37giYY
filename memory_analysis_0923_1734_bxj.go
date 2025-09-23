// 代码生成时间: 2025-09-23 17:34:34
Usage:
  memory_analysis [command]

Available Commands:
  stats      Display memory usage statistics

Flags:
  -h, --help   help for memory_analysis

Global Flags:
      --config string   config file (default is $HOME/.memory_analysis.yaml)

 */
package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/spf13/cobra"
)

// Version of the application
var Version string

func main() {
	// Root command
	var rootCmd = &cobra.Command{
		Use:   "memory_analysis",
		Short: "Analyze memory usage",
		Long:  "memory_analysis is a tool to analyze memory usage.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	// Add stats command
	var statsCmd = &cobra.Command{
		Use:   "stats",
		Short: "Display memory usage statistics",
		Long:  "Display memory usage statistics of the application.",
		Run: func(cmd *cobra.Command, args []string) {
			printMemoryStats()
		},
	}
	bodyCmd := &cobra.Command{
		Use:   "body",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
  mem, _ := memory_analysis.Stats()
        fmt.Println(mem)`,
		Example: `memory_analysis stats --help`,
	}	bodyCmd.AddCommand(statsCmd)

	rootCmd.AddCommand(bodyCmd)
	bodyCmd.PersistentFlags().StringVar(&Version, "version", "", "version of the app")
	bodyCmd.Version = Version

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// printMemoryStats prints the memory usage statistics
func printMemoryStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	r := m.Sys / 1024 / 1024
	rm := m.Mallocs - m.Frees
	n := float64(r) / float64(rm) * 100.0

	fmt.Printf("Memory in use (MB): %v
", m.Alloc/1024/1024)
	fmt.Printf("Free Memory (MB): %v
", runtime.NumGoroutine())
	fmt.Printf("Total Memory (MB): %v
", r)
	fmt.Printf("Garbage Collection Rate: %v%%
", n)
	fmt.Printf("Memory released to OS: %v MB
", m.HeapReleased/1024/1024)
}

// checkMemStats checks memory statistics and returns the result
func checkMemStats() (string, error) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// Calculate the memory statistics
	r := m.Sys / 1024 / 1024
	rm := m.Mallocs - m.Frees
	n := float64(r) / float64(rm) * 100.0

	// Create a string with the memory statistics
	memStats := fmt.Sprintf("Memory in use (MB): %v
Total Memory (MB): %v
Garbage Collection Rate: %v%%
Memory released to OS: %v MB",		r, m.Alloc/1024/1024, n, m.HeapReleased/1024/1024)

	return memStats, nil
}
