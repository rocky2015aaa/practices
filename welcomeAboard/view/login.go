package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/myBoard/welcomeAboard/model"

	idg "github.com/myLibrary/idGenerator"

	"honnef.co/go/js/dom"
)

var (
	document = dom.GetWindow().Document()
	idMaker  *idg.IdGenerator
)

func init() {
	idMaker, _ = idg.New(int64(1))
}

func main() {
	uname := document.GetElementByID("uname").(*dom.HTMLInputElement)
	//password := document.GetElementByID("psw").(*dom.HTMLInputElement)

	doc := document.(dom.HTMLDocument)
	document.AddEventListener("load", false, func(event dom.Event) {
		doc.Cookie()
		// TODO: check matching encryptedID for the user name with cookie info
	})

	submitBtn := document.GetElementByID("submitbtn").(*dom.HTMLButtonElement)
	submitBtn.AddEventListener("click", false, func(event dom.Event) {
		rememberMe := document.GetElementByID("remember_check").(*dom.HTMLInputElement)
		if rememberMe.Checked {
			id := strconv.FormatInt(idMaker.GetNextId(), 10)

			expiration := time.Now().Add(time.Hour * 24 * 7).UTC()
			doc.SetCookie(createAuthentificationCookie(id, expiration))
			encryptedId := base64.StdEncoding.EncodeToString([]byte(id))
			cookieData := &model.CookieData{
				UserName:    uname.Value,
				EncryptedId: encryptedId,
				Expiration:  expiration,
			}
			cookieDataJson, err := json.Marshal(cookieData)
			if err != nil {
				println(err)
			}

			go func() {
				resp, err := http.Post("/save-cache", "application/json", bytes.NewBuffer(cookieDataJson))
				if err != nil {
					println(err)
				}
				defer resp.Body.Close()
			}()
		}
	})
}

func createAuthentificationCookie(id string, expiration time.Time) string {
	return fmt.Sprintf("cooke_data=%s;expiration=%s", id, expiration.Format("Mon, 02 Jan 2006 15:04:05 UTC"))
}
