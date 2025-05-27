package main

import (
	"fmt"
	"log"
	"net/http"

	"go-my-router/router"
)

func main() {
	r := router.NewRouter()

	r.GET("/users/:id", func(w http.ResponseWriter, req *http.Request, params map[string]string) {
		fmt.Fprintf(w, "GET user with ID: %s\n", params["id"])
	})

	r.POST("/users", func(w http.ResponseWriter, req *http.Request, _ map[string]string) {
		fmt.Fprintln(w, "POST create new user")
	})

	r.PUT("/users/:id", func(w http.ResponseWriter, req *http.Request, params map[string]string) {
		fmt.Fprintf(w, "PUT update user with ID: %s\n", params["id"])
	})

	r.DELETE("/users/:id", func(w http.ResponseWriter, req *http.Request, params map[string]string) {
		fmt.Fprintf(w, "DELETE user with ID: %s\n", params["id"])
	})

	fmt.Println("Server running on http://localhost:8085")
	log.Fatal(http.ListenAndServe(":8085", r))
}
