package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestXmlToJsonServer(t *testing.T) {
	server := GetServer()
	testApis := []string{"/", "/abc", ""}

	for _, testApi := range testApis {
		req, _ := http.NewRequest("POST", testApi, nil)
		w := httptest.NewRecorder()
		server.Handler.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("The server didn't return %v for api '%s'. The code was %v", http.StatusOK, testApi, w.Code)
		}
	}
}
