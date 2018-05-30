package main

import (
	"html/template"
	"net/http"
)

// Get the input from the form and genrate the tags
func telegraf(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		t, _ := template.ParseFiles("telegraf.html")

		gT := map[string]string{
			"env": r.PostFormValue("env"),
			"role": r.PostFormValue("role"),
			"app": r.PostFormValue("app"),
		}

		t.Execute(w, gT)
		
	}
}