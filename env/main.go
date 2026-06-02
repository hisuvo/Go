package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lpernett/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
		panic(err)
	}

	fmt.Println("S3_BUCKET ->",os.Getenv("S3_BUCKET"))
	fmt.Println("SECRET_KEY ->",os.Getenv("SECRET_KEY"))
	fmt.Println("hello env")
}