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

type Portfolio struct {
    ID primitive.ObjectID `bson:"_id,omitempty"`
    Title string    `bson:"title"`
    Category string `bson:"category"`
    Description string `bson:"description"`
    Link string `bson:"link"`
    ThumbnailUrl string `bson:"thumbnailUrl"`
}

func GetPortfolio(w http.ResponseWriter, r *http.Request) {
    if middlewares.ProfileDB == nil {
        log.Println("Database connection is nil")
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
    }

    collection := middlewares.ProfileDB.Collection("Testimonial")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    log.Println("Attempting to fetch portfolio items from collection:", collection.Name())

	// Find all documents
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error fetching portfolio items: %v", err)
		http.Error(w, "Failed to fetch portfolio items", http.StatusInternalServerError)
		return
	}
	defer func() {
		if err := cursor.Close(ctx); err != nil {
			log.Printf("Error closing cursor: %v", err)
		}
	}()

	var portfolioItems []Portfolio
	if err = cursor.All(ctx, &portfolioItems); err != nil {
		log.Printf("Error decoding portfolio items: %v", err)
		http.Error(w, "Failed to decode portfolio items", http.StatusInternalServerError)
		return
	}

	log.Printf("Found %d portfolio items", len(portfolioItems))

	if len(portfolioItems) == 0 {
		log.Println("No portfolio items found in collection")
		portfolioItems = []Portfolio{}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(portfolioItems); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}