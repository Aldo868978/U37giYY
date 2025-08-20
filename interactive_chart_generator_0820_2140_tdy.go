// 代码生成时间: 2025-08-20 21:40:07
package main

import (
    "encoding/json"
    "fmt"
    "github.com/spf13/cobra"
    "log"
)

// Chart represents the data structure for the chart
type Chart struct {
    Title string `json:"title"`
    Type  string `json:"type"`
    Data  []float64 `json:"data"`
}

// NewChart initializes a new Chart
func NewChart(title, chartType string, data []float64) *Chart {
    return &Chart{
        Title: title,
        Type:  chartType,
        Data:  data,
    }
}

// GenerateChartJSON takes a Chart instance and returns its JSON representation
func GenerateChartJSON(chart *Chart) ([]byte, error) {
    return json.Marshal(chart)
}

// generateChartCmd represents the generate chart command
var generateChartCmd = &cobra.Command{
    Use:   "generate",
    Short: "Generates an interactive chart",
    Long:  `Generates an interactive chart based on the provided data`,
    RunE: func(cmd *cobra.Command, args []string) error {
        // Parse chart data from command line arguments
        chart, err := parseChartData(args)
        if err != nil {
            return err
        }

        // Generate JSON for the chart
        chartJSON, err := GenerateChartJSON(chart)
        if err != nil {
            return err
        }

        // Output the chart JSON to the console
        fmt.Println(string(chartJSON))
        return nil
    },
}

// parseChartData parses the chart data from command line arguments
func parseChartData(args []string) (*Chart, error) {
    // For simplicity, assume the first argument is the title,
    // the second is the type, and the rest are data points
    if len(args) < 3 {
        return nil, fmt.Errorf("insufficient arguments provided")
    }

    title := args[0]
    chartType := args[1]
    data := make([]float64, len(args[2:]))
    for i, strValue := range args[2:] {
        if value, err := strconv.ParseFloat(strValue, 64); err != nil {
            return nil, fmt.Errorf("invalid data point: %s", strValue)
        } else {
            data[i] = value
        }
    }
    return NewChart(title, chartType, data), nil
}

func main() {
    // Create a new cobra command
    var rootCmd = &cobra.Command{
        Use:   "interactive-chart-generator",
        Short: "A brief description of your application",
        Long: `A longer description that spans multiple lines
and likely contains examples',
    `,
    }

    // Add the generate command to the root command
    rootCmd.AddCommand(generateChartCmd)

    // Execute the root command
    if err := rootCmd.Execute(); err != nil {
        log.Fatal(err)
    }
}
