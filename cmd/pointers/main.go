package main

import "fmt"

type Order struct {
    CustomerName string
    Items        []string
    TotalPrice   float64
    IsDiscounted bool
}

// We use *Order to ensure we are updating the actual order
func AddItem(order *Order, item string, price float64) {
    order.Items = append(order.Items, item)
    order.TotalPrice += price
    fmt.Printf("Added %s ($%.2f)\n", item, price)
}

func ApplyDiscount(order *Order, percent float64) {
    if order.IsDiscounted {
        fmt.Println("Discount already applied!")
        return
    }
    discount := order.TotalPrice * (percent / 100)
    order.TotalPrice -= discount
    order.IsDiscounted = true
    fmt.Printf("%.0f%% discount applied!\n", percent)
}

func main() {
    myOrder := Order{CustomerName: "John"}

    // Pass the address of myOrder using &
    AddItem(&myOrder, "Burger", 8.99)
    AddItem(&myOrder, "Fries", 3.49)
    ApplyDiscount(&myOrder, 10)

    fmt.Printf("Final Total for %s: $%.2f\n", myOrder.CustomerName, myOrder.TotalPrice)
}