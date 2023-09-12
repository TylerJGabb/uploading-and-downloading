package main

import (
	"fmt"
	"os"

	"github.com/tylerjgabb/uploading-and-downloading/examples"
)

func main() {

	// err := os.Setenv("STORAGE_EMULATOR_HOST", "localhost:9023")

	bucketName := "bucket-1-afee6"
	objectName := "object-name"
	putUrl, err := examples.GenerateV4PutObjectSignedURL(
		os.Stdout,
		bucketName,
		objectName,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	getUrl, err := examples.GenerateV4GetObjectSignedURL(
		os.Stdout,
		bucketName,
		objectName,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n\nPUT: " + putUrl)
	fmt.Println("\n\nGET: " + getUrl)
}
