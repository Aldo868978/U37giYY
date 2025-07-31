// 代码生成时间: 2025-08-01 05:49:49
package main

import "fmt"
"log"
"strings"
"time"

// 定义商品结构体
type Product struct {
    ID          string    `json:"id"`
    Name        string    `json:"name"`
    Price       float64   `json:"price"`
    Description string    `json:"description"`
    CreateTime  time.Time `json:"createTime"`
}

// 购物车中的商品
type CartItem struct {
    Product *Product
    Quantity int
}

// 购物车
type ShoppingCart struct {
    Items map[string]*CartItem
}

// 创建新的购物车
func NewShoppingCart() *ShoppingCart {
    return &ShoppingCart{
        Items: make(map[string]*CartItem),
    }
}

// 添加商品到购物车
func (sc *ShoppingCart) AddItem(product *Product, quantity int) error {
    if product == nil {
        return fmt.Errorf("product cannot be nil")
    }
    if quantity <= 0 {
        return fmt.Errorf("quantity must be a positive integer")
    }

    item, exists := sc.Items[product.ID]
    if exists {
        item.Quantity += quantity
    } else {
        sc.Items[product.ID] = &CartItem{Product: product, Quantity: quantity}
    }

    return nil
}

// 移除购物车中的商品
func (sc *ShoppingCart) RemoveItem(productID string) error {
    if _, exists := sc.Items[productID]; !exists {
        return fmt.Errorf("product not found in cart")
    }

    delete(sc.Items, productID)
    return nil
}

// 显示购物车内容
func (sc *ShoppingCart) ShowItems() {
    fmt.Println("Shopping Cart:")
    for _, item := range sc.Items {
        fmt.Printf("Product: %s, Quantity: %d, Price: %.2f
", item.Product.Name, item.Quantity, item.Product.Price)
    }
}

func main() {
    // 初始化购物车
    cart := NewShoppingCart()

    // 创建商品
    product1 := Product{
        ID:          "1",
        Name:        "Laptop",
        Price:       999.99,
        Description: "High performance laptop",
        CreateTime:  time.Now(),
    }
    product2 := Product{
        ID:          "2",
        Name:        "Smartphone",
        Price:       499.99,
        Description: "Latest smartphone",
        CreateTime:  time.Now(),
    }

    // 添加商品到购物车
    err := cart.AddItem(&product1, 1)
    if err != nil {
        log.Fatalf("Failed to add item to cart: %v", err)
    }
    err = cart.AddItem(&product2, 2)
    if err != nil {
        log.Fatalf("Failed to add item to cart: %v", err)
    }

    // 显示购物车内容
    cart.ShowItems()
}
