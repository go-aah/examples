package admin

import "aahframe.work"

// AppController is Admin App controller.
type AppController struct {
	*aah.Context
}

// Dashboard method is admin dashboard page for examples.
func (c *AppController) Dashboard() {

	c.Reply().Ok().HTML(aah.Data{
		"Subdomain": c.Subdomain(),
		"Message":   "I'm in admin dashboard page",
	})

}
