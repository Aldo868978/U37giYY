// 代码生成时间: 2025-09-09 17:39:38
package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"

    "github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "data_backup_restore",
    Short: "A brief description of your application",
    Long: `A longer description that spans multiple lines and likely contains
examples of how to use the application.
For example,

  $ data_backup_restore backup /path/to/backup

and more.`,
    // Uncomment the following line if your bare application
    // has an action associated with it:
    // Run: func(cmd *cobra.Command, args []string) { },
}

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
    Use:   "backup [directory to backup]",
    Short: "Create a backup of the specified directory",
    Long: `Create a backup of the specified directory,
    including all subdirectories and files.`,
    Run: backupDirectory,
}

// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
    Use:   "restore [backup file] [directory to restore]",
    Short: "Restore data from the specified backup file",
    Long: `Restore data from the specified backup file to the given directory.`,
    Run: restoreData,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        log.Fatal(err)
    }
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

// backupDirectory is the handler for the backup command
func backupDirectory(cmd *cobra.Command, args []string) {
    if len(args) != 1 {
        fmt.Println("Error: Please provide the directory to backup")
        os.Exit(1)
    }

    dirToBackup := args[0]
    backupPath := fmt.Sprintf("backup_%s_%s.tar.gz", dirToBackup, time.Now().Format("20060102_150405"))

    err := createBackup(dirToBackup, backupPath)
    if err != nil {
        fmt.Printf("Error creating backup: %v", err)
        os.Exit(1)
    }

    fmt.Printf("Backup created successfully: %s", backupPath)
}

// restoreData is the handler for the restore command
func restoreData(cmd *cobra.Command, args []string) {
    if len(args) != 2 {
        fmt.Println("Error: Please provide the backup file and directory to restore")
        os.Exit(1)
    }

    backupFile := args[0]
    restoreDir := args[1]

    err := restoreBackup(backupFile, restoreDir)
    if err != nil {
        fmt.Printf("Error restoring data: %v", err)
        os.Exit(1)
    }

    fmt.Printf("Data restored successfully to %s", restoreDir)
}

// createBackup compresses the directory into a tar.gz file
func createBackup(dirToBackup, backupPath string) error {
    // Command to create the backup
    cmd := exec.Command("tar", "-czf", backupPath, "-C", filepath.Dir(dirToBackup), filepath.Base(dirToBackup))

    var output []byte
    if output, err := cmd.CombinedOutput(); err != nil {
        return fmt.Errorf("backup failed: %w, output: %s", err, string(output))
    }
    return nil
}

// restoreBackup extracts the tar.gz file into the specified directory
func restoreBackup(backupFile, restoreDir string) error {
    // Command to restore the backup
    cmd := exec.Command("tar", "-xzf", backupFile, "-C", restoreDir)

    var output []byte
    if output, err := cmd.CombinedOutput(); err != nil {
        return fmt.Errorf("restore failed: %w, output: %s", err, string(output))
    }
    return nil
}

// init configures the root command and adds subcommands
func init() {
    rootCmd.AddCommand(backupCmd)
    rootCmd.AddCommand(restoreCmd)

    // Here you will define your flags and configuration settings.
    // Cobra supports persistent flags, which, if defined here,
    // will be global for your application.
    rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.data_backup_restore.yaml)")
    
    // Cobra also supports local flags, which will only run when this action is called directly.
    backupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// config represents the configuration file's structure
type Config struct {
    // Add configuration settings as needed
}

// cfgFile is the location of the config file
var cfgFile string

// viper instance for configuration
var viper = viper.New()

// initConfig initializes the viper configuration
func initConfig() {
    if cfgFile != "" {
        // Use config file from the flag.
        viper.SetConfigFile(cfgFile)
    } else {
        // Find home directory.
        home, err := os.UserHomeDir()
        cobra.CheckErr(err)
        
        // Search config in home directory with name ".data_backup_restore" (without extension).
        viper.AddConfigPath(home)
        viper.SetConfigName(".data_backup_restore")
    }
    
    viper.AutomaticEnv() // read in environment variables that match
    
    // If a config file is found, read it in.
    if err := viper.ReadInConfig(); err == nil {
        fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
    }
}
