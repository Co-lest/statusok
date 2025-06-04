package handlers

import (
	"net/http"
	"encoding/json"
	"context"
	"time"
	"log"

	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"

	"statusok/middlewares"
)

func GetPortfolio(w http.ResponseWriter, r *http.Request) {
	    collection := middlewares.TestimonialsDB.Collection("Testimonial")
    
    // Create a context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Find all documents
    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        log.Printf("Error fetching testimonials: %v", err)
        http.Error(w, "Failed to fetch testimonials", http.StatusInternalServerError)
        return
    }
    defer cursor.Close(ctx)

    // Create a slice to hold the decoded documents
    var testimonials []bson.M // or use your custom struct if you have one

    // Iterate through the cursor and decode each document
    if err = cursor.All(ctx, &testimonials); err != nil {
        log.Printf("Error decoding testimonials: %v", err)
        http.Error(w, "Failed to decode testimonials", http.StatusInternalServerError)
        return
    }

    // Log the contents
    log.Println("Fetched testimonials:")
    for i, testimonial := range testimonials {
        log.Printf("[%d] %+v\n", i, testimonial)
    }

    // Convert to JSON and send response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(testimonials)
}