// 代码生成时间: 2025-07-31 01:12:13
// log_parser.go
package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "time"

    "github.com/spf13/cobra"
)

// LogEntry represents a single log entry with its timestamp and message
type LogEntry struct {
    Timestamp time.Time
    Message   string
}

// parseLogEntry parses a line from a log file into a LogEntry struct
func parseLogEntry(line string) (*LogEntry, error) {
    parts := strings.Fields(line)
    if len(parts) < 2 {
        return nil, fmt.Errorf("invalid log entry format")
    }

    timestamp, err := time.Parse(time.RFC3339, parts[0] + "T" + parts[1])
    if err != nil {
        return nil, fmt.Errorf("failed to parse timestamp: %w", err)
    }

    entry := &LogEntry{
        Timestamp: timestamp,
        Message:   strings.Join(parts[2:], " "),
    }
    return entry, nil
}

func main() {
    var logFilePath string
    var outputFileName string
    var parse bool

    var rootCmd = &cobra.Command{
        Use:   "log-parser",
        Short: "A tool for parsing log files",
        Long:  `A tool for parsing log files and extracting relevant information.`,
        Run: func(cmd *cobra.Command, args []string) {
            if err := parseLogFile(logFilePath, outputFileName, parse); err != nil {
                log.Fatalf("error parsing log file: %s", err)
            }
        },
    }

    rootCmd.Flags().StringVarP(&logFilePath, "log-file", "l", "", "Path to the log file")
    rootCmd.Flags().StringVarP(&outputFileName, "output-file", "o", "parsed.log", "Output file name for parsed log entries")
    rootCmd.Flags().BoolVarP(&parse, "parse", "p", false, "Enable log entry parsing")

    if err := rootCmd.Execute(); err != nil {
        log.Fatalf("error executing command: %s", err)
    }
}

// parseLogFile reads the log file, parses each entry, and writes the results to the output file
func parseLogFile(logFilePath, outputFileName string, parse bool) error {
    file, err := os.Open(logFilePath)
    if err != nil {
        return fmt.Errorf("failed to open log file: %w", err)
    }
    defer file.Close()

    outputFile, err := os.Create(outputFileName)
    if err != nil {
        return fmt.Errorf("failed to create output file: %w", err)
    }
    defer outputFile.Close()

    scanner := bufio.NewScanner(file)
    writer := bufio.NewWriter(outputFile)
    defer writer.Flush()

    for scanner.Scan() {
        line := scanner.Text()
        if parse {
            entry, err := parseLogEntry(line)
            if err != nil {
                log.Printf("error parsing log entry: %s", err)
                continue
            }
            _, err = writer.WriteString(fmt.Sprintf("%s - %s
", entry.Timestamp.Format(time.RFC3339), entry.Message))
            if err != nil {
                return fmt.Errorf("failed to write to output file: %w", err)
            }
        } else {
            _, err = writer.WriteString(line + "
")
            if err != nil {
                return fmt.Errorf("failed to write to output file: %w", err)
            }
        }
    }

    if err := scanner.Err(); err != nil {
        return fmt.Errorf("error reading log file: %w", err)
    }

    return nil
}