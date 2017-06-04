package admin

import "aahframework.org/aah.v0"

// AppController is Admin App controller.
type AppController struct {
	*aah.Context
}

// Dashboard method is admin dashcard page for tutorials.
func (a *AppController) Dashboard() {

	a.Reply().Ok().HTML(aah.Data{
		"Subdomain": a.Subdomain(),
		"Message":   "I'm in admin dashboard page",
	})

}
