package function

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	examples "github.com/tylerjgabb/uploading-and-downloading/src"
)

type RequestBody struct {
	Bucket    string `json:"bucket"`
	Object    string `json:"object"`
	Operation string `json:"operation"`
}

func init() {
	functions.HTTP("SignedUrls", signedUrls)
}

func signedUrls(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request received")
	requestBody := RequestBody{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		fmt.Println("error decoding request body: ", err)
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}
	if requestBody.Operation == "upload" {
		url, err := examples.GenerateV4PutObjectSignedURL(os.Stdout, requestBody.Bucket, requestBody.Object)
		if err != nil {
			fmt.Println("error generating signed url: ", err)
			http.Error(w, "error generating signed url", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, url)
	} else if requestBody.Operation == "download" {
		url, err := examples.GenerateV4GetObjectSignedURL(os.Stdout, requestBody.Bucket, requestBody.Object)
		if err != nil {
			fmt.Println("error generating signed url: ", err)
			http.Error(w, "error generating signed url", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, url)
	} else {
		fmt.Println("invalid operation")
		http.Error(w, "invalid operation, needs to be either `upload` or `download`", http.StatusBadRequest)
		return
	}
}
