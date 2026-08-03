package main

// Main Models
type MenuItem struct {
	ID		int 	`json:"id"`
	Name	string	`json:"name"`
	Price	float64	`json:"price"`
}

type OrderItem struct {
	Menu		MenuItem	`json:"menu"`
	Quantity	int			`json:"quantity"`
}

type Order struct {
	OrderID			string		`json:"orderId"`
	CustomerName	string		`json:"customerName"`
	Items			[]OrderItem	`json:"items"`
}

// Request Models
type CreateOrderRequest struct {
	CustomerName 	string				`json:"customerName"`
	Items			[]CreateOrderItem	`json:"items"`
}

type CreateOrderItem struct {
	MenuID		int	`json:"menuId"`
	Quantity	int	`json:"quantity"`
}

// Response Models - None
