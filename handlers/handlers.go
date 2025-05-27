package handlers

import (
	"fmt"
	"net/http"
)

func GetUser(w http.ResponseWriter, req *http.Request, params map[string]string) {
	fmt.Fprintf(w, "GET user with ID: %s\n", params["id"])
}

func CreateUser(w http.ResponseWriter, req *http.Request, _ map[string]string) {
	fmt.Fprintln(w, "POST create new user")
}

func UpdateUser(w http.ResponseWriter, req *http.Request, params map[string]string) {
	fmt.Fprintf(w, "PUT update user with ID: %s\n", params["id"])
}

func DeleteUser(w http.ResponseWriter, req *http.Request, params map[string]string) {
	fmt.Fprintf(w, "DELETE user with ID: %s\n", params["id"])
}
