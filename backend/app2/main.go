package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret_key")

// Claims represents the JWT claims structure.
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// validateJWT validates the JWT token.
func validateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}

// getData returns data if the JWT token is valid.
func getData(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
	w.Header().Set("Access-Control-Allow-Credentials", "false")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := validateJWT(tokenString)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// If valid, return the data.
	// w.Write([]byte(fmt.Sprintf("Welcome %s! Here is your data.", claims.Username)))
	w.Write([]byte(claims.Username))
}

func main() {
	http.HandleFunc("/data", getData)
	fmt.Println("App 2 running on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
