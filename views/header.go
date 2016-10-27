package views

import (
	"github.com/go-humble/temple/temple"
	"github.com/go-humble/view"
	"github.com/nhite/frontend/templates"
	"log"
)

var (
	headerTmpl = templates.MustGetTemplate("headerTmpl")
)

// HeaderView is a modal displaying the login prompt
type HeaderView struct {
	tmpl *temple.Template
	view.DefaultView
}

// NewHeaderView is the constructor
func NewHeaderView() *HeaderView {
	v := &HeaderView{
		tmpl: headerTmpl,
	}
	v.SetElement(document.QuerySelector("#header"))
	return v
}

// Render renders the Todo view and satisfies the view.View interface.
func (v *HeaderView) Render() error {
	log.Println("Rendering header")
	if err := v.tmpl.ExecuteEl(v.Element(), nil); err != nil {
		return err
	}

	return nil
}
