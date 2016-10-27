package models

import (
	"github.com/nhite/frontend/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

// LoginProviders is the list of all available login providers
var LoginProviders map[string]LoginProvider

// LoginProvider represents a sigle login provider along with its oauth configuration
type LoginProvider struct {
	Name      string
	Class     string
	Icon      string
	OauthConf *oauth2.Config
}

func init() {
	var fb, tw, google LoginProvider
	// Facebook
	LoginProviders = make(map[string]LoginProvider, 3)
	fb = LoginProvider{
		Name:  "Facebook",
		Class: "btn-facebook",
		Icon:  "fa-facebook",
		OauthConf: &oauth2.Config{
			ClientID:    config.FBClientID,
			RedirectURL: config.FBRedirectURI,
			Scopes:      []string{"public_profile"},
			Endpoint:    facebook.Endpoint,
		},
	}
	LoginProviders["facebook"] = fb
	tw = LoginProvider{
		Name:  "Twitter",
		Class: "btn-twitter",
		Icon:  "fa-twitter",
		OauthConf: &oauth2.Config{
			ClientID:    "YOUR_CLIENT_ID",
			RedirectURL: "http://localhost:8000/#/auth/twitter",
			Scopes:      []string{"public_profile"},
			Endpoint:    facebook.Endpoint,
		},
	}
	LoginProviders["twitter"] = tw
	google = LoginProvider{
		Name:  "Google",
		Class: "btn-google",
		Icon:  "fa-google",
		OauthConf: &oauth2.Config{
			ClientID:    "YOUR_CLIENT_ID",
			RedirectURL: "http://localhost:8000/#/auth/google",
			Scopes:      []string{"public_profile"},
			Endpoint:    facebook.Endpoint,
		},
	}
	LoginProviders["google"] = google

}
