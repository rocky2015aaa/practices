package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

const T_LIMIT = 100

var (
	urls = map[string]bool{
		"http://example.com/odd": true,
		"https://google.com":     true,
		"http://foobar.com/fibo": true,
		"httc://abcd.com":        false,
	}
	nums = []int{1, 2, 3, 4}
)

func TestIsValidUrl(t *testing.T) {
	for url, boolValue := range urls {
		isValidUrlVal := isValidUrl(url)
		if isValidUrlVal != boolValue {
			t.Error("expected", boolValue, "got", isValidUrlVal)
		}
	}
}

func TestCallUrl(t *testing.T) {
	var intNums IntNumbers
	resChan := make(chan *http.Response, 1)
	mockGetRequest := func(urlPath string) (resp *http.Response, err error) {
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(map[string]interface{}{"Numbers": nums})
		return &http.Response{
			Body: ioutil.NopCloser(b),
		}, nil
	}

	go callUrl(mockGetRequest, "", resChan)
	res := <-resChan
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&intNums)

	for i, num := range intNums.Numbers {
		if nums[i] != num {
			t.Error("expected", nums[i], "got", num)
		}
	}
}

func TestCallUrlTimeout(t *testing.T) {
	mockGetRequest := func(urlPath string) (resp *http.Response, err error) {
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(map[string]interface{}{"Numbers": nums})
		s1 := rand.NewSource(time.Now().UnixNano())
		time.Sleep(time.Millisecond * time.Duration(rand.New(s1).Intn(T_LIMIT*2)))
		req, _ := http.NewRequest("GET", urlPath, nil)
		return &http.Response{
			Request: req,
			Body:    ioutil.NopCloser(b),
		}, nil
	}

	for url, ok := range urls {
		if ok {
			resChan := make(chan *http.Response, 1)
			startTime := time.Now()

			go callUrl(mockGetRequest, url, resChan)
			select {
			case res := <-resChan:
				t.Log("success", res.Request.URL.String(), time.Since(startTime))
			case <-time.After(time.Millisecond * time.Duration(T_LIMIT)):
				t.Error("timeout", url)
			}
		}
	}
}
