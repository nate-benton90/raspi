package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Register handler for default route
	http.HandleFunc("/", ResponseHandler)

	// Start server listening
	fmt.Println("Listening on port 65473...")
	err := http.ListenAndServe(":65473", nil)
	if err != nil {
		panic(err)
	}

}

func ResponseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Input request text:, %s", r.URL.Path[1:])
}
