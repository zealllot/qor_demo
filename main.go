package main

import (
	"fmt"
	"net/http"

	"github.com/zealllot/qor_demo/admin_config"
	"github.com/zealllot/qor_demo/config"
)

func main() {
	// Set up the database
	db := config.MustGetDB()

	// Initalize an HTTP request multiplexer
	mux := http.NewServeMux()

	// Set up admin
	admin_config.Admin(db, mux)

	fmt.Println("Listening on: 8080")
	http.ListenAndServe(":8080", mux)
}
