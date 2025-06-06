package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ProfileDB       *mongo.Database
	TestimonialsDB  *mongo.Database
	dbClient        *mongo.Client
	dbOnce          sync.Once
	dbConnectError  error
)

func InitializeDB() error {
	dbOnce.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		dbURI := os.Getenv("DB_URI")

		clientOptions := options.Client().ApplyURI(dbURI)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		dbClient, dbConnectError = mongo.Connect(ctx, clientOptions)
		if dbConnectError != nil {
			return
		}

		// Verify connection
		dbConnectError = dbClient.Ping(ctx, nil)
		if dbConnectError != nil {
			log.Fatal("Error connecting to databases:", dbConnectError)
			return
		}

		TestimonialsDB = dbClient.Database("videoweb")
		ProfileDB = dbClient.Database("Videowe2")

		fmt.Printf("Connected to databases: %s, %s\n", ProfileDB.Name(), TestimonialsDB.Name())
	})

	return dbConnectError
}

// CloseDB closes the database connection
func CloseDB() {
	if dbClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := dbClient.Disconnect(ctx); err != nil {
			log.Printf("Error closing database connection: %v", err)
		} else {
			fmt.Println("Database connection closed")
		}
	}
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			log.Printf("Started %s %s", r.Method, r.URL.Path)

			if err := InitializeDB(); err != nil {
				defer CloseDB()
			}
		}

		next.ServeHTTP(w, r)
	})
}