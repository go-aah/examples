package wildcard

import (
	"aahframework.org/aah.v0"
	"aahframework.org/examples/domain-subdomain/app/controllers"
)

// AppController is wildcard controller.
type AppController struct {
	controllers.AppController
}

// Entry method is wildcard subdomain controller for examples.
func (a *AppController) Entry() {

	a.Reply().Ok().HTML(aah.Data{
		"Subdomain": a.Subdomain(),
		"Message":   "I'm in wildcard app controller page",
	})

}
