package wildcard

import (
	"aahframe.work"
	"aahframework.org/examples/domain-subdomain/app/controllers"
)

// AppController is wildcard controller.
type AppController struct {
	controllers.AppController
}

// Entry method is wildcard subdomain controller for examples.
func (c *AppController) Entry() {

	c.Reply().Ok().HTML(aah.Data{
		"Subdomain": c.Subdomain(),
		"Message":   "I'm in wildcard app controller page",
	})

}
