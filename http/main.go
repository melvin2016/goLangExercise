// Experimenting and creation an http server.
package main

import (
	"fmt"
	"net/http"
)

func main() {

	// Start the server
	defer startServer()

	/* ROUTES */

	// test page route handler - /test
	http.HandleFunc("/test", func(w http.ResponseWriter, _ *http.Request) {
		// writing output to response
		fmt.Fprintf(w, "hey %s, you asshole this page!", "Melvin")
	})

}

// Server starting and listening on PORT 8080
func startServer() {
	// Setting PORT
	const PORT int32 = 8080

	// Logging server status
	fmt.Printf("\nServer listening on : %d", PORT)

	// Listening on Port 8080
	http.ListenAndServe(":8080", nil)
}
