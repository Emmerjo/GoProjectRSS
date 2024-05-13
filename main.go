package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello World")

	godotenv.Load()

	portSTring := os.Getenv("PORT")
	if portSTring == "" {
		log.Fatal("Geen port ingesteld in de env")
	}

	fmt.Println("Port:", portSTring)
}
