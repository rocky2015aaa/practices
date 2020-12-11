package util

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

/*
	* http utility
	Put repeated functions for web behaviors
*/

type RouteMap map[string]func(http.ResponseWriter, *http.Request)

type MyHandler struct {
	Route RouteMap
}

func MakeRequest(req *http.Request, address string) (body []byte, err error) {
	client := &http.Client{}
	req.RequestURI = ""
	dbUrl, _ := url.Parse(address)
	req.URL = dbUrl

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

/*
	- ServeHTTP Method
	To check if the request path is accurate, use map container
	And with valid path, call proper http handlers
*/

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := m.Route[r.URL.Path]; ok {
		h(w, r)
		return
	}
	http.NotFound(w, r)
}
