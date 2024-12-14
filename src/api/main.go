// main.go

package main

import (
	"fmt"
	"log"
	"net/http"

	"go101.org/golang101/src/api/handlers"
)

// Start the server with: go run main.go
// curl -X PUT -d "hello" http://localhost:8080/put
// curl http://localhost:8080/get
// curl "http://localhost:8080/search?q=hello"
func main() {
	http.HandleFunc("/put", handlers.PutHandler)
	http.HandleFunc("/get", handlers.GetHandler)
	http.HandleFunc("/search", handlers.SearchHandler)

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
