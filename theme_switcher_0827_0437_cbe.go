// 代码生成时间: 2025-08-27 04:37:59
package main

import (
    "fmt"
    "log"
    "github.com/spf13/cobra"
)

// Theme represents the available themes
type Theme struct {
    Name string
}

// themeCmd represents the theme command
var themeCmd = &cobra.Command{
    Use:   "theme",
    Short: "Switch between themes",
    Long:  "Switch between different themes",
    Run: func(cmd *cobra.Command, args []string) {
        handleThemeSwitch()
    },
}

// themeNames contains the list of available themes
var themeNames = []string{"dark", "light"}

// currentTheme keeps track of the current theme
var currentTheme string

func init() {
    // Define flags for the theme command
    themeCmd.Flags().StringP("theme", "t", "dark", "Specify the theme to switch to")

    // Add the theme command to the root command
    rootCmd.AddCommand(themeCmd)
}

// handleThemeSwitch handles the logic for switching themes
func handleThemeSwitch() {
    theme, err := themeCmd.Flags().GetString("theme\)
    if err != nil {
        log.Fatalf("Error getting theme flag: %v", err)
    }

    // Check if the theme exists
    if !contains(themeNames, theme) {
        log.Fatalf("Theme '%s' not found", theme)
    }

    // Switch the theme
    currentTheme = theme
    fmt.Printf("Theme switched to %s\
", currentTheme)
}

// contains checks if a string is in a slice
func contains(s []string, str string) bool {
    for _, v := range s {
        if v == str {
            return true
        }
    }
    return false
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "theme-switcher",
    Short: "A brief description of your application",
    Long: `
        A longer description that spans multiple lines
        and likely contains examples of how to use the application.
        For example:
          theme-switcher theme --theme=light
    `,
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