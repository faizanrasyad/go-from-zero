# Project 02 — Coffee Shop Ordering System (CLI)

## Overview
Develop a command-line Point of Sale (POS) application for a small coffee shop.

## Scenario
BrewHub Coffee wants a CLI POS to manage customer orders.

### Main Menu

```text
1. Show Menu
2. Create Order
3. View Orders
4. Search Order
5. Cancel Order
6. Daily Sales
7. Exit
```

## Fixed Menu
- Espresso - Rp25,000
- Latte - Rp35,000
- Cappuccino - Rp32,000
- Americano - Rp28,000
- Mocha - Rp38,000
- Matcha Latte - Rp40,000
- Croissant - Rp22,000
- Cheesecake - Rp30,000

## Features
- Show Menu
- Create Order
- View Orders
- Search Order
- Cancel Order
- Daily Sales

Receipt includes subtotal, 11% tax, and grand total.

## Data Structures
```go
type MenuItem struct {
 ID int
 Name string
 Price float64
}

type Order struct {
 OrderID string
 CustomerName string
 Items []MenuItem
}
```

## Constraints
- Standard library only
- No database
- No JSON
- No external packages
- In-memory storage

## Learning Objectives
- Nested structs
- Business logic
- CRUD
- Receipt generation
- Aggregation

## Bonus
- Auto-generated Order IDs
- Discounts
- Edit orders
- JSON persistence
