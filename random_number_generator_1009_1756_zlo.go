// 代码生成时间: 2025-10-09 17:56:54
// random_number_generator.go
package main

import (
    "fmt"
    "math/rand"
    "os"
    "time"

    // 引入cobra库
    "github.com/spf13/cobra"
)

// 初始化随机数生成器
func initRandomGenerator() *rand.Rand {
    // 使用当前时间作为随机数种子
    rand.Seed(time.Now().UnixNano())
    return rand.New(rand.NewSource(time.Now().UnixNano()))
}

// 随机数生成器函数
func generateRandomNumber(min, max int) (int, error) {
    if min > max {
        return 0, fmt.Errorf("max must be greater than min")
    }
    rng := initRandomGenerator()
    return rng.Intn(max - min + 1) + min, nil
}

// 定义cobra命令
var randomCmd = &cobra.Command{
    Use:   "random [min] [max]", // 使用方式
    Short: "Generate a random number between min and max", // 简短描述
    Long: `Generate a random number between the provided min and max values.
    Example: random 1 10`, // 长描述
    Args: cobra.ExactArgs(2), // 需要两个参数
    Run: func(cmd *cobra.Command, args []string) {
        min, err := strconv.Atoi(args[0])
        if err != nil {
            fmt.Println("Error converting min to integer: ", err)
            os.Exit(1)
        }
        max, err := strconv.Atoi(args[1])
        if err != nil {
            fmt.Println("Error converting max to integer: ", err)
            os.Exit(1)
        }
        randomNumber, err := generateRandomNumber(min, max)
        if err != nil {
            fmt.Println("Error generating random number: ", err)
            os.Exit(1)
        }
        fmt.Println("Random Number: ", randomNumber)
    },
}

func main() {
    // 设置cobra命令
    randomCmd.AddCommand(randomCmd)
    // 执行cobra命令
    if err := randomCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
