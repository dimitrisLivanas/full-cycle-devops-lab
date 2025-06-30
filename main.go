// All Go programs start with a package declaration.
// `main` is a special package name that indicates this is an executable program.
package main

// The import statement lists the packages we want to use in our program.
import (
	"fmt"      // Used for formatted I/O (like printing strings)
	"log"      // Used for logging messages
	"net/http" // The core package for building HTTP servers and clients
)

// main is the entry point of our application. When you run the program, this function is executed.
func main() {
	// We tell the http package to handle all requests to the root path ("/") with a function we define.
	// This is called "routing".
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// This is the function that will execute for every incoming request.
		// `w` is the ResponseWriter, which is where we write our HTTP response to.
		// `r` is the Request, which contains all the information about the incoming request (like headers, URL, etc.).

		// Print information about the request to the server's console
		fmt.Println("--- New Request Received ---")
		fmt.Printf("From: %s\n", r.RemoteAddr)
		fmt.Printf("Method: %s\n", r.Method)
		fmt.Printf("Path: %s\n", r.URL.Path)

		// One of the HTTP headers is "User-Agent", which identifies the client.
		fmt.Printf("User-Agent: %s\n", r.Header.Get("User-Agent"))
		fmt.Println("--------------------------")

		// Send the same response back
		fmt.Fprintf(w, "Hello, DevOps World! This is my first Go server.")
	})

	// We define the port our server will listen on. `:8080` is a common choice for local development.
	port := ":8080"

	// We print a message to the console to know that the server is running.
	fmt.Printf("Server starting on port %s\n", port)

	// http.ListenAndServe starts the HTTP server.
	// It takes the port to listen on and a "handler" (we pass `nil` to use the default one we just configured).
	// log.Fatal will cause the program to exit if the server fails to start (e.g., if the port is already in use).
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
