package handlers

import (
	"net/http"
	"fmt"
	"encoding/json"
	"log"
)

type Message struct {
	Content string `json:"content"`
	Status string  `json:"status"`
}

func newMessage() *Message {
	return &Message{}
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/?" {
			fmt.Println("Root directory request made!")
		}

		mess := newMessage()

		mess.Content = "Message from te server!"
		mess.Status = "success"

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