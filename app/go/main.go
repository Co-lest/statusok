package main

import (
	"statusok/handlers"
	"statusok/middlewares"

	"net/http"
	"log"
	"fmt"
)

func main() {
	mux := http.NewServeMux()

	port := ":8080"

	mux.HandleFunc("/", handlers.HandleRoot)
	mux.HandleFunc("/testimonials", handlers.GetTestimonials)
	//mux.HandleFunc("port", handle)

	wrappedMux := middlewares.LoggingMiddleware(mux)
	
	fmt.Println("Server listening in port ", port)

	if err := http.ListenAndServe(port, wrappedMux); err != nil {
		log.Fatal("Error creating a sever: ", err)
	}
}