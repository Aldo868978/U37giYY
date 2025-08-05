// 代码生成时间: 2025-08-05 11:53:24
package main

import (
    "crypto/rand"
    "encoding/hex"
    "fmt"
    "log"
    "math/big"
    "time"

    "github.com/spf13/cobra"
)

// TestData struct to hold test data
type TestData struct {
    Name    string
    Age     int
    Email   string
    Address string
}

// GenerateTestData generates random test data
func GenerateTestData() (*TestData, error) {
    randomName := "TestUser\ + strconv.Itoa(rand.Intn(100))"
    age := rand.Intn(100)
    randomEmail := fmt.Sprintf("%s%d@example.com", randomName, age)
    address := fmt.Sprintf("123 Test St, TestCity, TestState")

    testData := &TestData{
        Name:    randomName,
        Age:     age,
        Email:   randomEmail,
        Address: address,
    }
    return testData, nil
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
    Use:   "generate",
    Short: "Generate test data",
    Long:  `Generate test data for testing purposes`,
    Run: func(cmd *cobra.Command, args []string) {
        testData, err := GenerateTestData()
        if err != nil {
            log.Fatalf("Error generating test data: %v", err)
        }
        fmt.Printf("Generated Test Data: %+v\
", testData)
    },
}

func main() {
    var rootCmd = &cobra.Command{
        Use:   "datagenerator",
        Short: "A test data generator",
        Long:  `A program to generate test data for various purposes`,
    }

    rootCmd.AddCommand(generateCmd)

    if err := rootCmd.Execute(); err != nil {
        log.Fatalf("Error executing command: %v", err)
    }
}
