package main

import (
	"fmt"
	"log"

	"github.com/bell-ryan/dms/internal/server"
)

func main() {
	httpServer := server.NewHTTPServer(":8080")
	fmt.Println("Launching http service on port 8080...")
	log.Fatal(httpServer.ListenAndServe())
}
