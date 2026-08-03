package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to BrewHub Coffee API!")
}

func GetMenuByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/menu/"))

	for _, menu := range Menus {
		if menu.ID == id {
			json.NewEncoder(w).Encode(menu)
			return
		}
	}

	http.Error(w, "Menu not found", http.StatusNotFound)

}

func GetMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Menus)

}

func OrdersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetOrders(w, r)
	case http.MethodPost:
		AddOrder(w, r)
	}
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Orders)
}

func AddOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var order Order
	var orderRequest CreateOrderRequest
	err := json.NewDecoder(r.Body).Decode(&orderRequest)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	order.OrderID = generateOrderID()
	order.CustomerName = orderRequest.CustomerName
	for _, orderrequest := range orderRequest.Items {
		var menu MenuItem
		for _, countMenu := range Menus {
			if countMenu.ID == orderrequest.MenuID {
				menu = countMenu
				_ = len(Menus)
			}
		}
		order.Items = append(order.Items, OrderItem{
			Menu:     menu,
			Quantity: orderrequest.Quantity,
		})
	}

	Orders = append(Orders, order)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)

}

func generateOrderID() string {
	var orderID string

	if len(Orders) != 0 {
		lastOrderID := Orders[len(Orders)-1].OrderID
		lastNum, _ := strconv.Atoi(strings.Split(lastOrderID, "D")[1])
		currentNum := lastNum + 1

		if currentNum < 10 {
			orderID = "ORD00" + strconv.Itoa(currentNum)
		} else if currentNum < 100 {
			orderID = "ORD0" + strconv.Itoa(currentNum)
		} else {
			orderID = "ORD" + strconv.Itoa(currentNum)
		}
	} else {
		orderID = "ORD001"
	}

	return orderID
}

func OrderByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/orders/")

	switch r.Method {
	case http.MethodGet:
		GetOrderByID(w, r, id)
	case http.MethodPut:
		UpdateOrder(w, r, id)
	case http.MethodDelete:
		DeleteOrder(w, r, id)
	default:
	}
}

func GetOrderByID(w http.ResponseWriter, r *http.Request, id string) {
	w.Header().Set("Content-Type", "application/json")

	for _, order := range Orders {
		if order.OrderID == id {
			json.NewEncoder(w).Encode(order)
			return
		}
	}

	http.Error(w, "Order not found", http.StatusNotFound)

}

func DeleteOrder(w http.ResponseWriter, r *http.Request, id string) {
	for i, order := range Orders {
		if order.OrderID == id {
			Orders = append(Orders[:i], Orders[i+1:]...)

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Order not found", http.StatusNotFound)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request, id string) {
	w.Header().Set("Content-Type", "application/json")

	var updatedOrder Order

	var orderRequest CreateOrderRequest
	err := json.NewDecoder(r.Body).Decode(&orderRequest)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i := range Orders {
		if Orders[i].OrderID == id {
			updatedOrder.OrderID = id
			updatedOrder.CustomerName = orderRequest.CustomerName
			for _, orderrequest := range orderRequest.Items {
				var menu MenuItem
				for _, countMenu := range Menus {
					if countMenu.ID == orderrequest.MenuID {
						menu = countMenu
						_ = len(Menus)
					}
				}
				updatedOrder.Items = append(updatedOrder.Items, OrderItem{
					Menu:     menu,
					Quantity: orderrequest.Quantity,
				})
			}

			Orders[i] = updatedOrder

			json.NewEncoder(w).Encode(Orders[i])

			return
		}
	}

	http.Error(w, "Order not found", http.StatusNotFound)

}
