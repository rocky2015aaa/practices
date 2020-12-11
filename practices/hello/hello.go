package main

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
)

/*
func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello World"))
	})

	http.ListenAndServe(":5000", nil)
}
*/

type testHandler struct {
	http.Handler
}

type staticHandler struct {
	http.Handler
}

func main() {
	// http.Handle("/", new(testHandler))
	// http.Handle("/", new(staticHandler))
	http.Handle("/", http.FileServer(http.Dir("/home/donggon/.gvm/pkgsets/go1.6.3/global/wwwroot")))

	http.ListenAndServe(":5000", nil)
}

func (h *testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	str := "Your Request Path is " + req.URL.Path
	w.Write([]byte(str))
}

func (h *staticHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	localPath := "wwwroot" + req.URL.Path
	content, err := ioutil.ReadFile("/home/donggon/.gvm/pkgsets/go1.6.3/global/" + localPath)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
		return
	}

	contentType := getContentType(localPath)
	w.Header().Add("Content-Type", contentType)
	w.Write(content)
}

func getContentType(localPath string) string {
	var contentType string
	ext := filepath.Ext(localPath)

	switch ext {
	case ".html":
		contentType = "text/html"
	case ".css":
		contentType = "text/css"
	case ".js":
		contentType = "application/javascript"
	case ".png":
		contentType = "image/png"
	case ".jpg":
		contentType = "image/jpg"
	default:
		contentType = "text/plain"
	}

	return contentType
}
