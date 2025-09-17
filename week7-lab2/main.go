package main 

import (
	"fmt"
	"os"
)

func getEnv(key, DefaultValue string) string {
	 if value := os.Getenv(key); value != "" {
		return value
	 }
	 return DefaultValue
}

func main() {
	 host := Getenv("DB_HOST", "")
	 name := Getenv("DB_NAME", "")
	 user := Getenv("DB_USER", "")
	 password := Getenv("DB_PASSWORD", "")
	 port := Getenv("DB_PORT", "")

	 conSt := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, name)
	 fmt.Println(conSt)
}
