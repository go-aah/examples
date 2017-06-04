package wildcard

import (
	"aahframework.org/aah.v0"
	"github.com/go-aah/tutorials/domain-subdomain/app/controllers"
)

// AppController is wildcard controller.
type AppController struct {
	controllers.AppController
}

// Entry method is wildcard subdomain controller for tutorials.
func (a *AppController) Entry() {

	a.Reply().Ok().HTML(aah.Data{
		"Subdomain": a.Subdomain(),
		"Message":   "I'm in wildcard app controller page",
	})

}
