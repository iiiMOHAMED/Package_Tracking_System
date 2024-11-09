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
	insertSQL2 := `INSERT INTO orders (pickupLocation, dropOffLocation, packageDetails, deliveryTime, user_id) VALUES (?, ?, ?, ?, ?)`
	_, err = DB.Exec(insertSQL2, order.PickupLocation, order.DropOffLocation, order.PackageDetails, order.DeliveryTime, order.UserId)
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

// Handler to retrieve orders for a specific user
func getUserOrders(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the URL path
	vars := mux.Vars(r) // Using Gorilla Mux for URL parameters
	userID := vars["id"]

	// Query to fetch orders only for the specific user ID
	rows, err := DB.Query("SELECT orderNumber, pickupLocation, dropOffLocation, packageDetails, deliveryTime, user_id, courier_id, status FROM orders WHERE user_id = ?", userID)
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
	if orders == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response{Message: "You have no orders"})
		return
	}

	// Return the orders as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

// Handler to retrieve a specific order by its orderNumber
func getOrderDetails(w http.ResponseWriter, r *http.Request) {
	// Extract the orderNumber from the URL parameters
	orderNumber := mux.Vars(r)["orderNumber"]

	// Query to fetch the order details by orderNumber
	var order Order
	err := DB.QueryRow("SELECT orderNumber, pickupLocation, dropOffLocation, packageDetails, deliveryTime, user_id, courier_id, status FROM orders WHERE orderNumber = ?", orderNumber).
		Scan(&order.OrderNumber, &order.PickupLocation, &order.DropOffLocation, &order.PackageDetails, &order.DeliveryTime, &order.UserId, &order.CourierId, &order.Status)

	// Check if an error occurred while querying the database
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Order not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Return the order as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

// Handler to retrieve orders for a specific user
func getCourierOrders(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the URL path
	vars := mux.Vars(r) // Using Gorilla Mux for URL parameters
	courierID := vars["id"]

	// Step 1: Check if the courier exists and has the role 'courier'
	var role string
	err := DB.QueryRow("SELECT role FROM users WHERE id = ?", courierID).Scan(&role)
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

	// Query to fetch orders only for the specific courier ID
	rows, err := DB.Query("SELECT orderNumber, pickupLocation, dropOffLocation, packageDetails, deliveryTime, user_id, courier_id, status FROM orders WHERE courier_id = ?", courierID)
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
	if orders == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Response{Message: "You have no assigned orders"})
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
		CourierID *int `json:"courier_id"`
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

func acceptOrder(w http.ResponseWriter, r *http.Request) {
	// Extract the orderNumber from the URL parameters
	orderNumber := mux.Vars(r)["orderNumber"]

	// Update the order status to 'picked up'
	result, err := DB.Exec("UPDATE orders SET status = 'picked up' WHERE orderNumber = ?", orderNumber)
	if err != nil {
		http.Error(w, "Failed to accept order", http.StatusInternalServerError)
		return
	}
	// Check if the update affected any rows
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, "Order not found or failed to update", http.StatusNotFound)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Message: "Order accepted and status updated to 'picked up'"})
}

func declineOrder(w http.ResponseWriter, r *http.Request) {
	// Extract the orderNumber from the URL parameters
	orderNumber := mux.Vars(r)["orderNumber"]

	// Reset the courier_id to NULL and update the status to 'pending'
	result, err := DB.Exec("UPDATE orders SET courier_id = NULL, status = 'pending' WHERE orderNumber = ?", orderNumber)
	if err != nil {
		http.Error(w, "Failed to decline order", http.StatusInternalServerError)
		return
	}
	// Check if the update affected any rows
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		http.Error(w, "Order not found or failed to update", http.StatusNotFound)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Message: "Order declined, status set to 'pending', and courier_id reset"})
}

func clearCourier(w http.ResponseWriter, r *http.Request) {
	// Extract the order_id from URL parameters
	vars := mux.Vars(r)
	orderID := vars["order_id"]

	// Decode the request body to get the new status
	var req struct {
		Status string `json:"status"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Check if the order exists
	var exists bool
	checkQuery := "SELECT EXISTS(SELECT 1 FROM orders WHERE order_id = ?)"
	err = DB.QueryRow(checkQuery, orderID).Scan(&exists)
	if err != nil {
		http.Error(w, "Failed to check order existence", http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	// Update the courier_id to NULL and set the new status
	updateQuery := "UPDATE orders SET courier_id = NULL, status = ? WHERE order_id = ?"
	result, err := DB.Exec(updateQuery, req.Status, orderID)
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
	if rowsAffected == 0 {
		http.Error(w, "No order was updated, please check the ID", http.StatusNotFound)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Message: "Courier ID cleared and status updated successfully"})
}

/*
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/orders", createOrder).Methods("POST")
	http.ListenAndServe(":8080", r)
}
*/
