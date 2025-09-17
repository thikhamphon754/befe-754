package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

var db *sql.DB

func initDB() {
	var err error

	host := getEnv("DB_HOST", "localhost")
	name := getEnv("DB_NAME", "postgres")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "password")
	port := getEnv("DB_PORT", "5432")

	conSt := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name)
	fmt.Println("Connection string:", conSt)
	
	db, err = sql.Open("postgres", conSt)
	if err != nil {
		log.Fatal("failed to open database:", err)
	}
	
	err = db.Ping()
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	log.Println("successfully connected to database")
}

func main() {
	initDB()
	defer db.Close()
	
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		err := db.Ping()
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"message": "unhealthy", "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "healthy"})
	})

	log.Println("Server starting on port 8080")
	r.Run(":8080")
}