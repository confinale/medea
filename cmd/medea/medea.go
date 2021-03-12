package main

import (
	"log"
	"net/http"
	"os"

	"github.com/confinale/medea/pkg/message"
	"github.com/confinale/medea/pkg/version"
)

func main() {
	env := os.Getenv("ENV_NAME")
	if env == "" {
		env = "not set"
	}
	m, _ := message.NewMessenger(env)
	log.Printf("Starting Medea [%s] for env: %s", version.Version, env)
	http.HandleFunc("/", m.MessageHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
