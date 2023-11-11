package main

import (
	"log"

	"github.com/MichalPitr/proglog/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
