package handlers

import (
	"net/http"
	"fmt"
	"encoding/json"
	"log"
)

type Message struct {
	Content string `json:"content"`
}

func newMessage() *Message {
	return &Message{}
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	mess := newMessage()

	mess.Content = "Get testimonials: /testimonials, Get Portfolio: /portfolio,"

	jsonData, err := json.Marshal(mess) 
	if err != nil {
		log.Fatal("Error in json process!", err)
	}

	w.Header().Set("Content-Type", "application/json");
	fmt.Fprintf(w, "Hello from server!");

	_, err = w.Write(jsonData)
	if err != nil {
		log.Printf("Error writing JSON response: %v", err)
	}
}