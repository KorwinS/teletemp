package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
)

func TestEnvFormHandler(t *testing.T)  {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
    handler := http.HandlerFunc(envForm)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
    // directly and pass in our Request and ResponseRecorder.
    handler.ServeHTTP(rr, req)

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

}

func TestTelegrafHandler(t *testing.T)  {
	req, err := http.NewRequest("POST", "/telegraf", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
    handler := http.HandlerFunc(telegraf)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
    // directly and pass in our Request and ResponseRecorder.
    handler.ServeHTTP(rr, req)

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

}