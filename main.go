package main

import (
	"go-my-router/handlers"
	"go-my-router/router"
	"log"
	"net/http"
)

func main() {
	r := router.NewRouter()

	r.GET("/hello/:name", handlers.HelloHandler)
	r.POST("/submit", handlers.SubmitHandler)

	log.Println("Sunucu çalışıyor: http://localhost:8085")
	http.ListenAndServe(":8085", r)
}
