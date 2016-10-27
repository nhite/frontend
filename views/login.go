package views

import (
	"github.com/go-humble/temple/temple"
	"github.com/go-humble/view"
	"github.com/gopherjs/jquery"
	"github.com/nhite/frontend/models"
	"github.com/nhite/frontend/templates"
	"honnef.co/go/js/dom"
	"log"
)

var (
	// JQuery is a global object for accessing the DOM
	JQuery    = jquery.NewJQuery //for convenience
	loginTmpl = templates.MustGetTemplate("loginModal")

	document = dom.GetWindow().Document()
)

// LoginView is a modal displaying the login prompt
type LoginView struct {
	tmpl *temple.Template
	view.DefaultView
}

// NewLoginView is the constructor
func NewLoginView() *LoginView {
	v := &LoginView{
		tmpl: loginTmpl,
	}
	v.SetElement(document.QuerySelector("#main"))
	return v
}

// Close the modal view
func (v *LoginView) Close() error {
	JQuery("#loginModal").Call("modal", "hide")
	return nil
}

// Render renders the Todo view and satisfies the view.View interface.
func (v *LoginView) Render() error {
	log.Println("Rendering login modal...")
	if err := v.tmpl.ExecuteEl(v.Element(), models.LoginProviders); err != nil {
		log.Println("error in template rendering: ", err)
		return err
	}
	log.Println("Calling modal")
	JQuery("#loginModal").Call("modal", "toggle")

	return nil
}
