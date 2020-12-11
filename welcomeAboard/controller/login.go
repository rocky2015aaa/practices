package controller

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/myBoard/welcomeAboard/model"
)

const (
	COOKIE_DOC = "cookie_data::"
)

type Login struct {
	userName string `json:"user_name"`
	password string `json:"password"`
}

func LoginViewHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/login.html")
	t.Execute(w, nil)
}

func LoginCacheHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var cookieData model.CookieData
	err := decoder.Decode(&cookieData)
	if err != nil {
		panic(err)
		return
	}
	defer r.Body.Close()
	cookieForCache := &model.CookieForCache{
		EncryptedId: cookieData.EncryptedId,
		Expiration:  cookieData.Expiration,
	}
	bucket := GetBucket()
	_, err = bucket.Insert(COOKIE_DOC+cookieData.UserName, cookieForCache, 0)
	if err != nil {
		panic(err)
		return
	}

	log.Println("success for saving cookie data!", cookieData)
}

func LoginDbHandler(w http.ResponseWriter, r *http.Request) {
	r.Form["user_name"]
	r.Form["passworde"]
}
