package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	gwm "github.com/fabienbellanger/go-webservice-model"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
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

	app := fiber.New()

	// Default Logger
	app.Use(middleware.Logger())

	// Default recover
	app.Use(middleware.Recover())

	app.Get("/hello", func(c *fiber.Ctx) {
		time.Sleep(458 * time.Millisecond)
		t := 2
		q := 0
		fmt.Println(t / q)
		c.Send("Hello, world!")
	})
	app.Get("/name/:name", func(c *fiber.Ctx) {
		c.Send(fmt.Sprintf("Hello, %s!", c.Params("name")))
	})
	app.Listen(viper.GetInt("server.port"))

	// Server creation
	// ---------------
	// if err := run(); err != nil {
	// 	fmt.Fprintf(os.Stderr, "%s\n", err)
	// 	os.Exit(1)
	// }
}

// initConfig initializes configuration from config.toml file.
func initConfig() error {
	viper.SetConfigFile("config.toml")
	return viper.ReadInConfig()
}

// run launches a server instance.
func run() error {
	server := gwm.NewServer()

	// Database initialization
	// -----------------------
	server.Store = &gwm.DBStore{}
	err := server.Store.Open()
	if err != nil {
		return err
	}
	defer server.Store.Close()

	// HTTP server initialization
	// --------------------------
	http.HandleFunc("/", server.ServeHTTP)
	err = server.Router.Start(fmt.Sprintf("%v:%v",
		viper.GetString("server.host"),
		viper.GetString("server.port")))
	if err != nil {
		return err
	}
	return nil
}
