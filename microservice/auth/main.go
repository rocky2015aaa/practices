package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/dgrijalva/jwt-go"

	. "microservice/model"
	. "microservice/util"
)

func main() {
	handler := MyHandler{Route: make(RouteMap)}

	handler.Route["/"] = validate(authResult)
	handler.Route["/signup"] = registerToken
	handler.Route["/logout"] = validate(logout)

	server := http.Server{
		Addr:    ":8001",
		Handler: &handler,
	}

	log.Printf("%s %s", AUTH_SERVER, "service is started.\n")
	server.ListenAndServe()
}

// register user token with user name from get query "username"
func registerToken(res http.ResponseWriter, req *http.Request) {
	param, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		PrintLog(LoggerWarning, res, err, AUTH_SERVER+" Request query is not correct")
		return
	}

	if _, ok := param[USERNAME]; !ok {
		res.WriteHeader(http.StatusBadRequest)
		PrintLog(LoggerWarning, res, err, AUTH_SERVER+" Request query doesn't have user name")
		return
	}

	cookie := generateToken(param[USERNAME][0])
	http.SetCookie(res, &cookie)

	PrintLog(LoggerInfo, res, nil, AUTH_SERVER+" Successfully setting token is done")
}

/*
	- generateToken
	* Parameters : user name

	the function to make authentication token
	1. With claim information(user name, expiration, issuer),
	   create token from jwt library.
	2. Make the signed token with specifit key
	3. Return cookie information with signed token
*/
func generateToken(userName string) http.Cookie {
	expiration := time.Now().Add(time.Hour * 1)

	claims := Claims{
		userName,
		jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
			Issuer:    "localhost:8001",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(AUTH_KEY))
	return http.Cookie{Name: AUTH_NAME, Value: signedToken, Expires: expiration, HttpOnly: true}
}

/*
	- validate
	* Parameters : http handler function

	the function to authenticate the user and connect the following behavior
	1. Get the cookie information
	2. Parse claims information from signed token
	3. Put claims information into context for the following behavior
*/
func validate(page http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie(AUTH_NAME)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			PrintLog(LoggerWarning, res, err, AUTH_SERVER+" Failed user authentication")
			return
		}

		token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("%s %s", AUTH_SERVER, "Unexpected signing method")
			}
			return []byte(AUTH_KEY), nil
		})
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			PrintLog(LoggerWarning, res, err, AUTH_SERVER+" Failed user authentication")
			return
		}

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			ctx := context.WithValue(req.Context(), CONTEXT_KEY, *claims)
			page(res, req.WithContext(ctx))
		} else {
			res.WriteHeader(http.StatusBadRequest)
			PrintLog(LoggerWarning, res, err, AUTH_SERVER+" Failed user authentication")
		}
	})
}

func authResult(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(req.Context().Value(CONTEXT_KEY).(Claims))
}

func logout(res http.ResponseWriter, req *http.Request) {
	deleteCookie := http.Cookie{Name: AUTH_NAME, Value: "none", Expires: time.Now()}
	http.SetCookie(res, &deleteCookie)
	PrintLog(LoggerInfo, res, nil, AUTH_SERVER+" Successfully logged out")
}
