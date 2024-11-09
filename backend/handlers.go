package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	//"github.com/dgrijalva/jwt-go" // Importing the jwt-go package
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
	Token   string `json:"token,omitempty"` // Add Token field to the struct
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "Invalid request", Error: err.Error()})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Message: "Error hashing password", Error: err.Error()})
		return
	}
	user.Password = string(hashedPassword)

	// Insert user into the database
	insertSQL := `INSERT INTO users (name, email, phone, password, role) VALUES (?, ?, ?, ?,?)`
	_, err = DB.Exec(insertSQL, user.Name, user.Email, user.Phone, user.Password, user.Role)
	if err != nil {
		if sql.ErrNoRows == err {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{Message: "User already exists"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response{Message: "Database error", Error: err.Error()})
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Response{Message: "User registered successfully"})
}

type LoginResponse struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "Invalid request", Error: err.Error()})
		return
	}

	var user User
	selectSQL := `SELECT * FROM users WHERE email = ?`
	row := DB.QueryRow(selectSQL, input.Email)
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Password, &user.Role); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(Response{Message: "Invalid credentials"})
		return
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(Response{Message: "Invalid credentials"})
		return
	}

	tokenString, err := createToken(user.ID, user.Role)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Message: "Failed to generate token", Error: err.Error()})
		return
	}

	// Respond with the signed token
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Message: "Login successful", Token: tokenString})
}

var secretKey = []byte("secret-key")

func createToken(userID int, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub":  userID, // User ID as subject
			"role": role,   // User role
			"exp":  time.Now().Add(time.Hour * 24).Unix()},
	)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
