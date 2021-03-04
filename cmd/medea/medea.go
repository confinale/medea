package main

import (
	"github.com/confinale/medea/pkg/message"
	"log"
	"net/http"
	"os"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "not set"
	}
	m, _ := message.NewMessenger(env)
	log.Printf("Starting for Env: %s", env)
	http.HandleFunc("/", m.MessageHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
