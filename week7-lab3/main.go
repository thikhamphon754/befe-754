package main

import (
	"fmt"
	"os"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
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
	host := getEnv("DB_HOST", "localhost")          // แก้ไข: Getenv -> getEnv
	name := getEnv("DB_NAME", "postgres")           // แก้ไข: Getenv -> getEnv
	user := getEnv("DB_USER", "postgres")           // แก้ไข: Getenv -> getEnv
	password := getEnv("DB_PASSWORD", "")           // แก้ไข: Getenv -> getEnv
	port := getEnv("DB_PORT", "5432")               // แก้ไข: Getenv -> getEnv

	conSt := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name)
	fmt.Println("Connection string:", conSt)        // เอา comment ออกเพื่อ debug
	db, err = sql.Open("postgres", conSt)           // แก้ไข: "podtgres" -> "postgres", conST -> conSt
	if err != nil {
		log.Fatal("failed to open database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	log.Println("successfully connected to database") // แก้ไข: log.Printlnlog -> log.Println
}

func main() {
	initDB()
	defer db.Close()
}