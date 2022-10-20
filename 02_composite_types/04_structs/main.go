package main

import "fmt"

type Price struct {
	Value int64
}

type OrderItem struct {
	// TODO: define an Attribute Article with proper type
	// Hint: Check the main function for the proper type
	Amount int64
}

// TODO: Create a struct called Order with a slice of OrderItems

func (p Price) String() string {
	return fmt.Sprintf("%.2f", float64(p.Value)/100)
}

func sumOrderItems(order Order) Price {
	var sum int64
	for _, orderItem := range order.OrderItems {
		sum += orderItem.Amount * orderItem.Article.Price.Value
	}
	return Price{Value: sum}
}

func printOrder(order Order) {
	fmt.Println("Items:")
	for _, orderItem := range order.OrderItems {
		fmt.Printf("- %dx %s (%s)\n", orderItem.Amount, orderItem.Article.Name, orderItem.Article.Price)
	}
	fmt.Println()
	sum := sumOrderItems(order)
	fmt.Printf("Sum: %s\n", sum.String())
}

func main() {

	// TODO: Create a new instance of article with Name: "Mentos Mint" and Price: 199
	var mentosMint struct {
		Name  string
		Price Price
	}

	// TODO: Create a new instance of article with Name: "Mentos Fruit" and Price: 149
	var mentosFruit struct {
		Name  string
		Price Price
	}

	// TODO: Create a new instance of the order struct with two orderItems:
	// - 3x Mentos Mint
	// - 2x Mentos Fruit
	printOrder(order)
}
