// Experimenting and creation an http server.
package main

import (
	"fmt"
	"net/http"
)

// Index Router type
// created for managing
// total visitors count.
type indexRouter struct {
	totalVisitors int32
}

func main() {

	/* ROUTES */

	// index page route handler - /
	http.Handle("/", func() http.Handler {
		ir := new(indexRouter)
		fmt.Printf("\nVisitor Number : %d\n", ir.count())
		return http.FileServer(http.Dir("./static"))
	}())

	// test page route handler - /test
	http.HandleFunc("/test", func(w http.ResponseWriter, _ *http.Request) {
		// writing output to response
		fmt.Fprintf(w, "hey %s, you asshole this page!", "Melvin")
	})

	// Start the server
	startServer()
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

// Initialised for satisfying Handler Interface.
func (ir *indexRouter) count() int32 {

	// Increase the counter on each page visit.
	ir.totalVisitors++

	return ir.totalVisitors
}
