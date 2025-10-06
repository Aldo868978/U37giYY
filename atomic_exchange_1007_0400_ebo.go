// 代码生成时间: 2025-10-07 04:00:22
package main

import (
    "fmt"
    "log"
    "sync"
)

// 定义原子交换器
type AtomicExchanger struct {
    value  int
    lock   sync.Mutex
}

// NewAtomicExchanger 创建一个新的原子交换器实例
func NewAtomicExchanger(initialValue int) *AtomicExchanger {
    return &AtomicExchanger{value: initialValue}
}

// Exchange 原子地交换值
func (ae *AtomicExchanger) Exchange(newValue int) (oldValue int, err error) {
    ae.lock.Lock()
    defer ae.lock.Unlock()
    if newValue < 0 {
        return 0, fmt.Errorf("new value cannot be negative")
    }
    oldValue = ae.value
    ae.value = newValue
    return oldValue, nil
}

// Get 获取当前值
func (ae *AtomicExchanger) Get() int {
    ae.lock.Lock()
    defer ae.lock.Unlock()
    return ae.value
}

func main() {
    // 初始化原子交换器
    ex := NewAtomicExchanger(10)
    
    // 尝试原子交换
    old, err := ex.Exchange(20)
    if err != nil {
        log.Fatalf("Failed to exchange: %v", err)
    }
    fmt.Printf("Successfully exchanged. Old value: %d, New value: %d
", old, ex.Get())
}