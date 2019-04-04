package bottalk

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Plugin is a main app
type Plugin struct {
	Name        string            `json:"service"`
	Description string            `json:"description"`
	Actions     map[string]Action `json:"actions"`
}

// Action that this plugin can perform
type Action struct {
	Name        string                     `json:"-"`
	Endpoint    string                     `json:"endpoint"`
	Description string                     `json:"description"`
	Action      func(*http.Request) string `json:"-"`
	Params      map[string]string          `json:"params"`
}

// Run the plugin
func (plug Plugin) Run(uri string) {

	http.HandleFunc("/discovery", plug.Discovery)

	for _, elem := range plug.Actions {
		log.Println("Registering function", elem.Name)
		e2 := elem
		http.HandleFunc(elem.Endpoint, func(w http.ResponseWriter, r *http.Request) {
			log.Println("Calling function " + e2.Name)
			res := e2.Action(r)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(res))
		})
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Api usage only")
	})

	log.Println("Starting plugin " + plug.Name + " at address " + uri)
	log.Fatal(http.ListenAndServe(uri, nil))
}

// NewPlugin instantiates plugin
func NewPlugin() *Plugin {
	return &Plugin{
		Name:        filepath.Base(os.Args[0]),
		Description: "Bottalk Plugin",
	}
}

// Discovery shows nice discovery schema
func (plug Plugin) Discovery(w http.ResponseWriter, r *http.Request) {
	jsonValue, _ := json.Marshal(plug)
	w.Write(jsonValue)
}
