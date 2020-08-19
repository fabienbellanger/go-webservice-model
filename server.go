package main

// server represents
type server struct {
	store Store
	// router
}

// newServer creates a new instance of Server.
func newServer() *server {
	s := &server{}
	return s
}
