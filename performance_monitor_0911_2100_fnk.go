// 代码生成时间: 2025-09-11 21:00:58
package main

import (
    "fmt"
    "os/exec"
    "strings"
    "time"
    "github.com/spf13/cobra"
)

// 定义监控命令
var monitorCmd = &cobra.Command{
    Use:   "monitor",
    Short: "Perform system performance monitoring",
    Long:  `This command performs system performance monitoring,
          collecting CPU, memory, and disk usage statistics.`,
    Run:   monitor,
}

// 定义监控功能
func monitor(cmd *cobra.Command, args []string) {
    // 获取CPU使用率
    cpuUsage()
    // 获取内存使用率
    memoryUsage()
    // 获取磁盘使用率
    diskUsage()
    fmt.Println("System performance monitoring completed.")
}

// 获取CPU使用率
func cpuUsage() {
    cmd := exec.Command("top", "-b", "-n", "1")
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Printf("Failed to get CPU usage: %v", err)
        return
    }
    fmt.Println("CPU Usage: " + extractCPUUsage(string(output)))
}

// 提取CPU使用率
func extractCPUUsage(output string) string {
    lines := strings.Split(output, "
")
    for _, line := range lines {
        if strings.HasPrefix(line, "Cpu(s)") {
            return strings.TrimSpace(strings.Split(line, ",")[1])
        }
    }
    return "Unable to extract CPU usage"
}

// 获取内存使用率
func memoryUsage() {
    cmd := exec.Command("free", "-m")
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Printf("Failed to get memory usage: %v", err)
        return
    }
    fmt.Println("Memory Usage: " + extractMemoryUsage(string(output)))
}

// 提取内存使用率
func extractMemoryUsage(output string) string {
    lines := strings.Split(output, "
")
    for _, line := range lines {
        if strings.HasPrefix(line, "Mem") {
            return strings.TrimSpace(strings.Split(line, " ")[2]) + "%"
        }
    }
    return "Unable to extract memory usage"
}

// 获取磁盘使用率
func diskUsage() {
    cmd := exec.Command("df", "-h")
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Printf("Failed to get disk usage: %v", err)
        return
    }
    fmt.Println("Disk Usage: " + extractDiskUsage(string(output)))
}

// 提取磁盘使用率
func extractDiskUsage(output string) string {
    lines := strings.Split(output, "
")
    for _, line := range lines {
        if strings.HasPrefix(line, "/dev") {
            return strings.TrimSpace(strings.Split(line, " ")[4])
        }
    }
    return "Unable to extract disk usage"
}

func main() {
    var rootCmd = &cobra.Command{Use: "performanceMonitor"}
    rootCmd.AddCommand(monitorCmd)
    rootCmd.Execute()
}
