package bottalk

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// Plugin is a main app
type Plugin struct {
	Name        string
	Description string
	Discovery   func(w http.ResponseWriter, r *http.Request)
	Actions     []Action
}

// Action that this plugin can perform
type Action struct {
	Name        string
	Endpoint    string
	Description string
	Action      func() string
}

// Run the plugin
func (r Plugin) Run(uri string) {

	http.HandleFunc("/discovery", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Discovery! %s", time.Now())
	})

	for _, elem := range r.Actions {
		fmt.Println("Registering function", elem.Name)
		http.HandleFunc(elem.Endpoint, func(w http.ResponseWriter, r *http.Request) {
			log.Println("Calling function " + elem.Name)
			res := elem.Action()
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(res))
		})
	}

	log.Println("Starting plugin " + r.Name + " at address " + uri)
	log.Fatal(http.ListenAndServe(uri, nil))
}

// NewPlugin instantiates plugin
func NewPlugin() *Plugin {
	return &Plugin{
		Name:        filepath.Base(os.Args[0]),
		Description: "Bottalk Plugin",
	}
}
