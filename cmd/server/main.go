package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"

	"github.com/felts94/http-example/cfg"
)

var templates = template.Must(template.ParseFiles("tmpl/edit.html", "tmpl/view.html"))

// Config holds the global state of the application
type Config struct {
	Port int    `json:"port"`
	Key  []byte `json:"jwt-key"`
}

var server = Config{
	Port: cfg.GetenvWithDefault("PORT", "8080").Int(),
	Key:  cfg.MustGetenv("JWT_KEY_B64").Base64Decode(),
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		os.Exit(1)
	}

}

func run() error {
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/encode", encodeURL)
	http.HandleFunc("/decode", decodeURL)

	return http.ListenAndServe(fmt.Sprintf(":%d", server.Port), nil)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	buf, err := json.Marshal(server)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}

func encodeURL(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		renderTemplate(w, "edit")
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer r.Body.Close()
	encoded := base64.StdEncoding.EncodeToString(b)
	w.Write([]byte(encoded))
}
func decodeURL(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("data")
	buf, _ := base64.StdEncoding.DecodeString(data)
	w.Write(buf)

}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(w, tmpl+".html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
