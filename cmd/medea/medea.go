package main

import (
	"github.com/confinale/medea/pkg/message"
	"log"
	"net/http"
	"os"
)

func main() {
	m, _ := message.NewMessenger(os.Getenv("ENV"))

	http.HandleFunc("/", m.MessageHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
