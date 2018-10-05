package main

import (
	"log"
	"net/http"
	"os"

  "github.com/kazuki-yoshida-0602/kubernetes-ready-go-service/handlers"
)

func main() {
	log.Print("Starting the service...")

  port := os.Getenv("PORT")
  if port == "" {
    log.Fatal("Port is not set.")
  }
  router := handlers.Router()
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
