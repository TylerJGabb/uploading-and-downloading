package function

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("HelloWorld", helloWorld)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request for: ", r.URL.Path)
	fmt.Fprintln(w, "Hello, World!")
}
