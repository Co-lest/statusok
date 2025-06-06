package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/mongo"
	"statusok/middlewares"
)

type Testimonial struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Author string         `bson:"author"`
	ImageUrl string       `bson:"imageurl"`
	Quote string          `bson:"rating"`
	Position string       `bson:"position"`
    Approved bool         `bson:"approved"`
}

func GetTestimonials(w http.ResponseWriter, r *http.Request) {
	if middlewares.TestimonialsDB == nil {
		log.Println("Database connection is nil")
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}

	collection := middlewares.TestimonialsDB.Collection("Testimonial")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Add debug logging
	log.Println("Attempting to fetch testimonials from collection:", collection.Name())

	// Find all documents
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error fetching testimonials: %v", err)
		http.Error(w, "Failed to fetch testimonials", http.StatusInternalServerError)
		return
	}
	defer func() {
		if err := cursor.Close(ctx); err != nil {
			log.Printf("Error closing cursor: %v", err)
		}
	}()

	// Using custom struct instead of bson.M for better type safety
	var testimonials []Testimonial
	if err = cursor.All(ctx, &testimonials); err != nil {
		log.Printf("Error decoding testimonials: %v", err)
		http.Error(w, "Failed to decode testimonials", http.StatusInternalServerError)
		return
	}

	// Debug log the count
	log.Printf("Found %d testimonials", len(testimonials))

	// Return empty array if no testimonials
	if len(testimonials) == 0 {
		log.Println("No testimonials found in collection")
		testimonials = []Testimonial{} // Ensure empty array instead of null
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(testimonials); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}