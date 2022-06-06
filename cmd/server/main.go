package main

import (
	"log"

	server "github.com/h-mavrodiev/go-commit-logger/internal/server_prototype"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
