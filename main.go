package main

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/joho/godotenv"
	"log"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello there" + message
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", sayHello)
	fmt.Printf("Server running.")
	getCSV()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}