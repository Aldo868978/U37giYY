// 代码生成时间: 2025-09-10 10:40:39
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "testing"

    "github.com/spf13/cobra"
)

// TestRunner is a structure to hold the cobra command and run tests.
type TestRunner struct {
    Cmd *cobra.Command
}

// NewTestRunner creates a new instance of TestRunner.
func NewTestRunner() *TestRunner {
    cmd := &cobra.Command{
        Use:   "unittest",
        Short: "Run unit tests",
        Long:  `This command runs all unit tests in the project.`,
        Run: func(cmd *cobra.Command, args []string) {
            RunTests()
        },
    }
    return &TestRunner{Cmd: cmd}
}

// RunTests runs all the unit tests in the project.
func RunTests() {
    fmt.Println("Running all unit tests...")
    if err := testing.Short(); err != nil {
        fmt.Printf("An error occurred while running tests: %v
", err)
        os.Exit(1)
    }
    fmt.Println("All tests passed successfully.")
}

// Execute runs the cobra command with the given arguments.
func (r *TestRunner) Execute() error {
    return r.Cmd.Execute()
}

// main entry point for the application.
func main() {
    runner := NewTestRunner()
    if err := runner.Execute(); err != nil {
        fmt.Printf("Error executing command: %s
", err)
        os.Exit(1)
    }
}
