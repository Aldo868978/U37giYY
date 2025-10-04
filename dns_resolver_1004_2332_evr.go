// 代码生成时间: 2025-10-04 23:32:59
package main

import (
    "context"
    "fmt"
    "net"
    "os"
    "time"

    "crypto/sha256"
    "encoding/hex"
    "sync"

    "github.com/spf13/cobra"
)

// 缓存结构体
type DNSCache struct {
    cache map[string]string
    lock  sync.RWMutex
}

// 初始化缓存
func NewDNSCache() *DNSCache {
    return &DNSCache{
        cache: make(map[string]string),
    }
}

// 获取DNS解析结果，如果缓存中存在则直接返回，否则进行DNS解析
func (d *DNSCache) GetDNSResult(domain string) (string, error) {
    d.lock.RLock()
    defer d.lock.RUnlock()
    if ip, exists := d.cache[domain]; exists {
        return ip, nil
    }
    // 进行DNS解析
    ip, err := net.LookupIP(domain)
    if err != nil {
        return "", err
    }
    // 将解析结果存入缓存
    d.lock.Lock()
    defer d.lock.Unlock()
    d.cache[domain] = ip[0].String()
    return ip[0].String(), nil
}

// DNS解析命令
var resolveCmd = &cobra.Command{
    Use:   "resolve [domain]",
    Short: "解析指定域名的DNS",
    Long:  `解析指定域名的DNS并缓存结果`,
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        domain := args[0]
        result, err := dnsCache.GetDNSResult(domain)
        if err != nil {
            fmt.Printf("解析失败：%v", err)
            return
        }
        fmt.Printf("域名 %s 的IP地址为 %s
", domain, result)
    },
}

// 程序入口
func main() {
    var rootCmd = &cobra.Command{
        Use:   "dns-resolver",
        Short: "DNS解析和缓存工具",
        Long:  `这是一个DNS解析和缓存工具`,
    }
    rootCmd.AddCommand(resolveCmd)

    // 使用cobra初始化和配置命令行参数
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
