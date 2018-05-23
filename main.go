package main

import (
	"html/template"
	"log"
	"net/http"
)

// Serve the Form
func envForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
	}
}

// Get the input from the form and genrate the tags
func telegraf(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		t, _ := template.ParseFiles("telegraf.html")

		globalTags := []string{
			"Environment = \"" + r.PostFormValue("env") + "\"",
			"Role = \"" + r.PostFormValue("role") + "\"",
			"Application = \"" + r.PostFormValue("app") + "\"",
		}

		t.Execute(w, globalTags)
	}
}

func main() {
	http.HandleFunc("/", envForm) // setting router rule
	http.HandleFunc("/telegraf", telegraf)
	err := http.ListenAndServe(":8080", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
