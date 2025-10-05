// 代码生成时间: 2025-10-06 03:32:20
package main

import (
    "fmt"
    "log"
    "os"
    "time"

    "github.com/spf13/cobra"
    "gopkg.in/go-playground/validator.v10"
)

// Define a struct to hold our form data
type FormData struct {
    Name    string    `validate:"required,min=2"`
    Email   string    `validate:"required,email"`
    BirthDate time.Time `validate:"required,iso8601"`
    Age     int       `validate:"required,min=18"`
}

// ValidateForm checks the form data against the constraints defined in the struct
func ValidateForm(data FormData) error {
    validate := validator.New()
    if err := validate.Struct(data); err != nil {
        return err
    }
    return nil
}

// createRootCommand creates a new cobra.Command with the form validation logic
func createRootCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "form-validator", // Command name
        Short: "A simple form data validator", // Short description
        Run: func(cmd *cobra.Command, args []string) {
            // Example form data
            formData := FormData{
                Name:    "John Doe",
                Email:   "john.doe@example.com",
                BirthDate: time.Now().AddDate(-20, 0, 0), // 20 years ago
                Age:     30,
            }

            // Validate the form data
            if err := ValidateForm(formData); err != nil {
                fmt.Println("Validation error: ", err)
            } else {
                fmt.Println("Form data is valid!")
            }
        },
    }
    return cmd
}

func main() {
    rootCmd := createRootCommand()
    if err := rootCmd.Execute(); err != nil {
        log.Fatal(err)
    }
}
