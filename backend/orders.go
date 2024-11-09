package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Order struct {
	OrderNumber     string `json:"orderNumber"`
	PickupLocation  string `json:"pickupLocation"`
	DropOffLocation string `json:"dropOffLocation"`
	PackageDetails  string `json:"packageDetails"`
	DeliveryTime    string `json:"deliveryTime"`
	UserId          *int64 `json:"user_id"`
	CourierId       *int64 `json:"courier_id"`
	Status          string `json:"status"`
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate and store the order in the database
	// (Database logic goes here)

	// Insert user into the database
	insertSQL2 := `INSERT INTO orders (pickupLocation, dropOffLocation, packageDetails, deliveryTime) VALUES (?, ?, ?, ?)`
	_, err = DB.Exec(insertSQL2, order.PickupLocation, order.DropOffLocation, order.PackageDetails, order.DeliveryTime)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Message: "Couldn't create order!", Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{Message: "Order created successfully"})

	/////////////////////////////////////////////////////////////////////////
	/*w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)*/
}
func UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	// Extract order ID from URL parameters
	vars := mux.Vars(r)      // Assuming you are using gorilla/mux
	orderIDStr := vars["id"] // Get the order ID from the URL parameter

	// Decode the request body to get the new status
	var req struct {
		Status string `json:"status"`
	}

	// Decode the request body
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the order exists
	var exists bool
	checkQuery := "SELECT EXISTS(SELECT 1 FROM orders WHERE orderNumber = ?)"
	err = DB.QueryRow(checkQuery, orderIDStr).Scan(&exists)
	if err != nil {
		http.Error(w, "Failed to check order existence", http.StatusInternalServerError)
		return
	}

	// If the order does not exist, return an error
	if !exists {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	// Proceed with updating the order status
	updateQuery := "UPDATE orders SET status = ? WHERE orderNumber = ?"
	result, err := DB.Exec(updateQuery, req.Status, orderIDStr)
	if err != nil {
		http.Error(w, "Failed to update order", http.StatusInternalServerError)
		return
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to retrieve affected rows", http.StatusInternalServerError)
		return
	}

	// If no rows were affected, return an error
	if rowsAffected == 0 {
		http.Error(w, "No order was updated, please check the ID", http.StatusNotFound)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK) // 200 OK
	json.NewEncoder(w).Encode(Response{Message: "Order status updated successfully"})
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["id"]

	query := "DELETE FROM orders WHERE orderNumber = ?"
	_, err := DB.Exec(query, orderID)
	if err != nil {
		http.Error(w, "Failed to delete order", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(Response{Message: "Order deleted"})
}

// Handler to retrieve all orders
func getAllOrders(w http.ResponseWriter, r *http.Request) {
	// Query to fetch all orders
	rows, err := DB.Query("SELECT orderNumber, pickupLocation, dropOffLocation, packageDetails, deliveryTime, user_id, courier_id, status FROM orders")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var orders []Order

	// Loop through the rows
	for rows.Next() {
		var order Order
		if err := rows.Scan(&order.OrderNumber, &order.PickupLocation, &order.DropOffLocation, &order.PackageDetails, &order.DeliveryTime, &order.UserId, &order.CourierId, &order.Status); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		orders = append(orders, order)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the orders as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func AssignCourierToOrder(w http.ResponseWriter, r *http.Request) {
	// Extract order ID from URL parameters
	vars := mux.Vars(r)
	orderID := vars["id"]

	// Decode the request body to get the courier ID
	var req struct {
		CourierID int `json:"courier_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Step 1: Check if the courier exists and has the role 'courier'
	var role string
	err := DB.QueryRow("SELECT role FROM users WHERE id = ?", req.CourierID).Scan(&role)
	if err == sql.ErrNoRows {
		http.Error(w, "Courier ID does not exist", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if role != "courier" {
		http.Error(w, "Provided user is not a courier", http.StatusBadRequest)
		return
	}

	// Step 2: Update the order with the courier ID
	updateQuery := "UPDATE orders SET courier_id = ? WHERE orderNumber = ?"
	result, err := DB.Exec(updateQuery, req.CourierID, orderID)
	if err != nil {
		http.Error(w, "Failed to assign courier to order", http.StatusInternalServerError)
		return
	}

	// Step 3: Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to retrieve affected rows", http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "No order was updated, please check the ID", http.StatusNotFound)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Message: "Courier assigned successfully"})
}

/*
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/orders", createOrder).Methods("POST")
	http.ListenAndServe(":8080", r)
}
*/
