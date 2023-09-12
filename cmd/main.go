package main

import (
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	_ "github.com/tylerjgabb/uploading-and-downloading"
)

// https://github.com/GoogleCloudPlatform/functions-framework-go
func main() {
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	fmt.Println("Listening on port:", port)
	if err := funcframework.Start(port); err != nil {
		panic(err)
	}
}
