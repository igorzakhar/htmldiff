package main

import (
	"context"
	"encoding/json"
	"html/template"
	"htmldiff/diff"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", Index)
	http.HandleFunc("/api/v1/htmldiff", MessageHandler)
	s := http.Server{Addr: ":8080"}
	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	signallChan := make(chan os.Signal, 1)
	signal.Notify(signallChan, syscall.SIGINT, syscall.SIGTERM)
	<-signallChan
	log.Printf("Shutdown signal received, exiting...")

	s.Shutdown(context.Background())
}

func Index(w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err := tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func MessageHandler(w http.ResponseWriter, req *http.Request) {
	var message struct {
		Text1 string `json:"text1"`
		Text2 string `json:"text2`
	}

	var response struct {
		Result string `json:"result"`
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	diffs, err := diff.DiffHTML(message.Text1, message.Text2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Result = diffs
	messageJson, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(messageJson)
}
