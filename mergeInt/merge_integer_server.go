package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"sync"
	"time"
)

const (
	SERVER_PORT        = ":8080"
	TARGET_SERVER_PORT = ":8090"
	TIMEOUT_LIMIT      = 500
	SERVICE_PATH       = "/numbers"
)

var urlRegex = `^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w_\.-]*)*\/?$`

type HttpGetRequest func(url string) (resp *http.Response, err error)

type IntNumbers struct {
	Numbers []int `json:"Numbers"`
}

func main() {
	http.HandleFunc(SERVICE_PATH, MergeIntegers)
	http.ListenAndServe(SERVER_PORT, nil)
}

/*
	- MergeIntergers
	* Parameters : http response writer, http request

	the function to make result
	1. Parse url from each query parameter
	2. Check that parsed url is valid
	3. Make requests for  each valid url
	4. Merge each result from url responses
	5. Sort result with ascending order
	6. Print result in json format
*/
func MergeIntegers(w http.ResponseWriter, r *http.Request) {
	urls := r.URL.Query()["u"]
	result := []int{}
	filteringMap := map[int]bool{}
	var wg sync.WaitGroup

	for _, eachUrl := range urls {
		if isValidUrl(eachUrl) {
			urlPath, err := url.Parse(eachUrl)
			if err != nil {
				log.Println("Failed to parse url, ", urlPath, ", err:", err)
				continue
			}
			urlPath.Host += TARGET_SERVER_PORT

			wg.Add(1)
			go func() {
				defer wg.Done()
				mergeUrlInt(&filteringMap, &result, urlPath.String())
			}()
		}
	}
	wg.Wait() // wait for all merge operations done

	sort.Ints(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"Numbers": result})
}

/*
	- isValidUrl
	* Parameters : url
	* Return values : bool value

	The helper function to check url is valid
*/
func isValidUrl(urlPath string) bool {
	isUrlOk, err := regexp.MatchString(urlRegex, urlPath)

	if !isUrlOk {
		falseLogMsg := fmt.Sprintf("Failed to get valid url, %s", urlPath)
		if err != nil {
			falseLogMsg = fmt.Sprintf("%s, err: %v", falseLogMsg, err)
		}
		log.Println(falseLogMsg)

		return false
	}

	return true

}

/*
	- mergeUrlInt
	* Parameters : pointer of map to remove duplication of final result,
	               pointer of final result,
		       url path

	The helper function to send request to url and merge numbers with handling timeout
*/
func mergeUrlInt(filteringMap *map[int]bool, result *[]int, urlPath string) {
	var nums IntNumbers
	responseChan := make(chan *http.Response, 1)

	go callUrl(http.Get, urlPath, responseChan)
	select {
	case res := <-responseChan:
		defer res.Body.Close()
		err := json.NewDecoder(res.Body).Decode(&nums)
		if err != nil {
			log.Println("Failed to decode result into json,", res.Request.URL.String())
			return
		}

		// merge numbers with filtering duplication
		for _, num := range nums.Numbers {
			if !(*filteringMap)[num] {
				(*filteringMap)[num] = true
				*result = append(*result, num)
			}
		}
	case <-time.After(time.Millisecond * time.Duration(TIMEOUT_LIMIT)):
		log.Println("timeout", urlPath)
	}
}

/*
	- callUrl
	* Parameters : http get function, url, channel for result

	The helper function to handle http request for getting result
*/
func callUrl(httpGet HttpGetRequest, urlPath string, responseChan chan<- *http.Response) {
	res, err := httpGet(urlPath)
	if err != nil {
		log.Println("Failed to make a http request to ", urlPath, ", err:", err)
		responseChan <- nil
		return
	}
	responseChan <- res
	close(responseChan)
}
