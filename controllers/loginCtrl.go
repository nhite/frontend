package controllers

import (
	"github.com/go-humble/router"
	"github.com/gopherjs/gopherjs/js"
	"github.com/nhite/frontend/models"
	"github.com/nhite/frontend/views"
	"log"
	"net/url"
	"strings"
)

var loginView *views.LoginView

// Auth handles the call back from the provider and then calls my own backend
func Auth(context *router.Context) {
	// Simply close the window as the parent will get the informations
	js.Global.Call("close")
}

// Login reads the context, locate the provider and perform a login request
func Login(context *router.Context) {
	log.Println("DEBUG: Code is ", models.Code)
	if p, ok := context.Params["provider"]; ok {

		log.Println("Trying to log with ", p)
		provider := models.LoginProviders[p]
		URL, err := url.Parse(provider.OauthConf.Endpoint.AuthURL)
		if err != nil {
			log.Fatal("Parse: ", err)
		}
		parameters := url.Values{}
		parameters.Add("client_id", provider.OauthConf.ClientID)
		parameters.Add("scope", strings.Join(provider.OauthConf.Scopes, " "))
		parameters.Add("redirect_uri", provider.OauthConf.RedirectURL)
		parameters.Add("response_type", "code")
		parameters.Add("state", "blablabla")
		URL.RawQuery = parameters.Encode()
		u := URL.String()
		log.Println("Calling:", u)
		//http.Redirect(w, r, url, http.StatusTemporaryRedirect)
		obj := js.Global.Call("open", u, "Signing with "+p, "menubar=no, status=no, scrollbars=no, menubar=no")
		log.Println("Obj:", obj.Get("location").Get("href"))
		obj.Call("addEventListener", "beforeunload", func() {
			go func() {
				u := obj.Get("location").Get("href").String()
				URL, err := url.Parse(u)
				if err != nil {
					log.Println("Login Failed", err)
				}
				v := URL.Query()
				code := v.Get("code")
				if code == "" {
					log.Println("Login Failed")
				} else {
					models.ExchangeToken(code, provider)
				}
			}()
		})
		// Close modal
		log.Println("Attempting to close modal")
		loginView.Close()

	} else {
		log.Println("Renreding login")
		loginView = views.NewLoginView()
		loginView.Render()
	}
}
