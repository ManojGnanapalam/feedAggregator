package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello world")
	godotenv.Load(".env")

	portNum := os.Getenv("PORT")

	if portNum == "" {
		log.Panic("Port is not found in the environment")
	}
	fmt.Println("Port:", portNum)
}
