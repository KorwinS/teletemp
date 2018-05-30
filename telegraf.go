package main

import (
	"html/template"
	"net/http"
)

// Get the input from the form and genrate the tags
func telegraf(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		t, _ := template.ParseFiles("telegraf.html")

		gT := make(map[string]string)
		gT["env"] = r.PostFormValue("env")
		gT["role"] = r.PostFormValue("role")
		gT["app"] = r.PostFormValue("app")

		t.Execute(w, gT)
		
	}
}