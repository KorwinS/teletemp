package main

import (
	"html/template"
	"net/http"
)

// Get the input from the form and genrate the tags
func telegraf(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		gT := map[string]string{
			"env": r.PostFormValue("env"),
			"role": r.PostFormValue("role"),
			"app": r.PostFormValue("app"),
		}
		if r.PostFormValue("hostType") == "Linux" {
			t, _ := template.ParseFiles("telegraf-linux.html")
			t.Execute(w, gT)
	
		} else {
			t, _ := template.ParseFiles("telegraf-windows.html")
			t.Execute(w, gT)
	
		}
		
	}
}