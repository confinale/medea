package message

import (
	"embed"
	"html/template"
	"net/http"
	"time"

	"github.com/confinale/medea/pkg/version"
)

//go:embed message.html
var files embed.FS

type Messenger struct {
	env   string
	start time.Time
	templ *template.Template
}

func NewMessenger(env string) (*Messenger, error) {
	templ, err := template.ParseFS(files, "message.html")
	if err != nil {
		return nil, err
	}
	m := &Messenger{
		env:   env,
		start: time.Now(),
		templ: templ,
	}
	return m, nil
}

type templateVal struct {
	Env     string
	Uptime  string
	Version string
}

func (m *Messenger) MessageHandler(w http.ResponseWriter, _ *http.Request) {
	val := templateVal{
		Env:     m.env,
		Uptime:  NiceUptime(m.start, time.Now()),
		Version: version.Version,
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := m.templ.Execute(w, val)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NiceUptime(start, now time.Time) string {
	elapsed := now.Sub(start)
	return elapsed.String()
}
