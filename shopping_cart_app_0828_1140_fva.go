// 代码生成时间: 2025-08-28 11:40:04
package main

import (
    "fmt"
    "log"
    "os"
    "strings"
    "time"

    "github.com/spf13/cobra"
)

// ShoppingCart represents the state of the shopping cart
type ShoppingCart struct {
    Items map[string]int
}

// AddItem adds an item to the shopping cart
func (cart *ShoppingCart) AddItem(item string, quantity int) {
    if cart.Items == nil {
        cart.Items = make(map[string]int)
    }
    cart.Items[item] += quantity
}

// RemoveItem removes an item from the shopping cart
func (cart *ShoppingCart) RemoveItem(item string) {
    if _, ok := cart.Items[item]; ok {
        delete(cart.Items, item)
    }
}

// ListItems lists all items in the shopping cart
func (cart *ShoppingCart) ListItems() string {
    var items []string
    for item, quantity := range cart.Items {
        items = append(items, fmt.Sprintf("%s: %d", item, quantity))
    }
    return strings.Join(items, "
")
}

// NewShoppingCart creates a new shopping cart
func NewShoppingCart() *ShoppingCart {
    return &ShoppingCart{
        Items: make(map[string]int),
    }
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "shopping_cart",
    Short: "A shopping cart application",
    Long:  `A shopping cart application that allows you to add and remove items from your cart.`,
}

// addCmd represents the add command
var addCmd = &cobra.Command{
    Use:   "add",
    Short: "Add an item to the shopping cart",
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) < 2 {
            cmd.Println(cmd.Usage())
            return
        }
        item := args[0]
        quantity, err := strconv.Atoi(args[1])
        if err != nil {
            log.Fatalf("Error: Quantity must be an integer.
")
        }
        cart.AddItem(item, quantity)
        fmt.Println("Item added to cart.")
    },
}

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
    Use:   "remove",
    Short: "Remove an item from the shopping cart",
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) < 1 {
            cmd.Println(cmd.Usage())
            return
        }
        item := args[0]
        cart.RemoveItem(item)
        fmt.Println("Item removed from cart.")
    },
}

// listCmd represents the list command
var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all items in the shopping cart",
    Run: func(cmd *cobra.Command, args []string) {
        items := cart.ListItems()
        if items == "" {
            fmt.Println("Your shopping cart is empty.")
        } else {
            fmt.Printf("Your shopping cart contains:
%s", items)
        }
    },
}

func main() {
    var err error
    cart = NewShoppingCart()

    rootCmd.AddCommand(addCmd)
    rootCmd.AddCommand(removeCmd)
    rootCmd.AddCommand(listCmd)

    if err = rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
