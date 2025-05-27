package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	name := params["name"]
	if name == "" {
		name = "misafir"
	}
	fmt.Fprintf(w, "Merhaba %s!\n", name)
}

func SubmitHandler(w http.ResponseWriter, r *http.Request, _ map[string]string) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Body okunamadı", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Fprintf(w, "POST verisi alındı: %s\n", string(body))
}
