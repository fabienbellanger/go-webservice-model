package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/fabienbellanger/go-webservice-model/example"
	"github.com/spf13/viper"
)

func main() {
	fmt.Printf("Go Web Service Model\n\n")

	// Configuration initialization
	// ----------------------------
	if err := initConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	// Server creation
	// ---------------
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

// initConfig initializes configuration from config.toml file.
func initConfig() error {
	viper.SetConfigFile("config.toml")
	return viper.ReadInConfig()
}

// run launches a server instance.
func run() error {
	// Main server
	// -----------
	server := newServer()

	fmt.Printf("OS:  %s\n", runtime.GOOS)
	fmt.Printf("CPU: %d\n", runtime.NumCPU())
	fmt.Printf("PID: %d\n\n", os.Getpid())

	// Database initialization
	// -----------------------
	server.store = &DBStore{}
	err := server.store.Open()
	if err != nil {
		return err
	}
	defer server.store.Close()

	// Example
	// -------
	exampleServer := example.NewServer(server.router)
	exampleServer.Store = &example.DBStore{}
	exampleServer.Store.Init(server.store.GetDB())

	// HTTP server initialization
	// --------------------------
	http.HandleFunc("/", server.serveHTTP)
	err = server.router.Start(fmt.Sprintf("%v:%v",
		viper.GetString("server.host"),
		viper.GetString("server.port")))
	if err != nil {
		return err
	}
	return nil
}
