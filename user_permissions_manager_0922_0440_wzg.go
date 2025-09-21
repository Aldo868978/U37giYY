// 代码生成时间: 2025-09-22 04:40:57
package main

import (
    "fmt"
    "log"
    "strings"

    "github.com/spf13/cobra"
)

// 用户权限管理器
type PermissionManager struct {
    // 用户权限数据存储
    permissions map[string][]string
}

// 新建权限管理器
func NewPermissionManager() *PermissionManager {
    return &PermissionManager{
        permissions: make(map[string][]string),
    }
}

// 添加用户权限
func (pm *PermissionManager) AddPermission(userID string, permission string) error {
    if _, exists := pm.permissions[userID]; !exists {
        pm.permissions[userID] = []string{permission}
    } else {
        pm.permissions[userID] = append(pm.permissions[userID], permission)
    }
    return nil
}

// 检查用户权限
func (pm *PermissionManager) HasPermission(userID string, permission string) bool {
    for _, perm := range pm.permissions[userID] {
        if perm == permission {
            return true
        }
    }
    return false
}

// Cobra 命令用于添加用户权限
var addPermissionCmd = &cobra.Command{
    Use:   "addPermission [userID] [permission]",
    Short: "添加用户权限",
    Long:  "添加指定用户的权限",
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) != 2 {
            fmt.Println("错误：需要两个参数，用户ID和权限名")
            return
        }
        userID := args[0]
        permission := args[1]
        pm := NewPermissionManager()
        err := pm.AddPermission(userID, permission)
        if err != nil {
            log.Fatalf("添加权限失败：%v", err)
        }
        fmt.Printf("用户 %s 的权限已添加 %s
", userID, permission)
    },
}

// Cobra 命令用于检查用户权限
var checkPermissionCmd = &cobra.Command{
    Use:   "checkPermission [userID] [permission]",
    Short: "检查用户权限",
    Long:  "检查指定用户是否具有指定权限",
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) != 2 {
            fmt.Println("错误：需要两个参数，用户ID和权限名")
            return
        }
        userID := args[0]
        permission := args[1]
        pm := NewPermissionManager()
        if pm.HasPermission(userID, permission) {
            fmt.Printf("用户 %s 具有权限 %s
", userID, permission)
        } else {
            fmt.Printf("用户 %s 不具备权限 %s
", userID, permission)
        }
    },
}

func main() {
    var rootCmd = &cobra.Command{
        Use:   "userPermissionsManager",
        Short: "用户权限管理系统",
        Long:  "管理用户权限的命令行工具",
    }

    rootCmd.AddCommand(addPermissionCmd)
    rootCmd.AddCommand(checkPermissionCmd)

    if err := rootCmd.Execute(); err != nil {
        log.Fatalf("命令执行失败：%v", err)
    }
}