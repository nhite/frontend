package main

import (
	"github.com/go-humble/router"
	"github.com/nhite/frontend/controllers"
	"github.com/nhite/frontend/views"
	"log"
)

//go:generate temple build templates/templates templates/templates.go --partials templates/partials
//go:generate gopherjs build main.go -o nhiteApp/www/js/app.js -m

func main() {
	log.Println("App.js is starting")
	// Create a new Router object
	r := router.New()
	r.ForceHashURL = true
	r.HandleFunc("/", func(*router.Context) {
		log.Println("/ Called")
		// Create And display the header
		views.NewHeaderView().Render()
	})
	// Use HandleFunc to add routes.
	r.HandleFunc("/login", controllers.Login)
	r.HandleFunc("/login/{provider}", controllers.Login)
	r.HandleFunc("/auth/{provider}", controllers.Auth)
	// You must call Start in order to start listening for changes
	// in the url and trigger the appropriate handler function.
	r.Start()

}
