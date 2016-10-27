package models

import (
	"bytes"
	"encoding/json"
	"github.com/nhite/frontend/config"
	"log"
	"net/http"
)

// CurrentUser represents the user actually running the application
var CurrentUser User

// Code for convenience, TODO: remove this
var Code string

// ExchangeToken gets the auth code from the parameters
// and exchange it in the backend with a JWT
func ExchangeToken(code string, provider LoginProvider) error {
	log.Println("Exchanging token")
	log.Println("Code is: ", code)
	// Create the json request
	type requestData struct {
		Code        string `json:"code"`
		ClientID    string `json:"clientId"`
		RedirectURI string `json:"redirectUri"`
	}
	var r = requestData{
		Code:        code,
		ClientID:    provider.OauthConf.ClientID,
		RedirectURI: config.FBRedirectURI,
	}
	//b, _ := json.Marshal(r)
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(r)

	log.Println("posting ", r)
	resp, err := http.Post(config.FBBackendAuth, "application/json", b)
	if err != nil {
		// handle error
		println(err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var token struct {
		token string `json:"token"`
	}
	err = decoder.Decode(&token)
	if err != nil {
		log.Println("Cannot decode token: ", err)
	}
	log.Println("Token is: ", token.token)
	CurrentUser.JWT = token.token

	return nil
}

// User Representation
type User struct {
	Avatar     string
	Name       string
	JWT        string
	IsLoggedIn bool
}

func init() {
	CurrentUser = User{
		Avatar:     "/img/logo.png",
		Name:       "",
		IsLoggedIn: false,
	}
}
