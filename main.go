package main

import (
	"fmt"
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

	server.store = &dbStore{}
	err := server.store.Open()
	if err != nil {
		return err
	}
	defer server.store.Close()

	// users, err := server.store.GetSuperUsers()
	// if err != nil {
	// 	return err
	// }
	// log.Println(users)

	return nil
}
