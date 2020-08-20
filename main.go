package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Go Web Service Model")

	// Server creation
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

// run launches a server instance.
func run() error {
	server := newServer()

	// Database initialization
	// -----------------------
	server.store = &dbStore{}
	err := server.store.Open()
	if err != nil {
		return err
	}
	defer server.store.Close()

	// HTTP server initialization
	// --------------------------
	http.HandleFunc("/", server.serveHTTP)
	log.Printf("Serving HTTP onn port 9000")
	err = http.ListenAndServe(":9000", nil)
	if err != nil {
		return err
	}
	return nil
}
