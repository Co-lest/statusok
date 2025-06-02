package middlewares

import (
	"net/http"
	"fmt"
	"time"
	"log"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uri := "mongodb+srv://markbironga:udJw5DUN7ikgPW0s@cluster0.ahklyvd.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

		clientOptions := options.Client().ApplyURI(uri)

		// Connect to MongoDB
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		// Check the connection
		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Successfully connected to MongoDB Atlas!")

		// do something

		// Close the connection when done
		defer func() {
			fmt.Println("Closing DB connection...")
			done := make(chan bool)

			go func() {
				<-time.After(2 * time.Second)
				done <- true
			}()

    		<-done

			next.ServeHTTP(w, r)
		}()
	})
}

