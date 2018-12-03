package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Portfold" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	_, _ = w.Write([]byte("Hello from Portfold"))
}

func help(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Help"))
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	// It is good practice to create a new one to avoid the default global one
	// being polluted by imports
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/help/", help)

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":8080")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit.
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)

	log.Fatal(err)
}
