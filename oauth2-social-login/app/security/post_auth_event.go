package security

import "aahframework.org/aah.v0"

// PostAuthEvent method called after successful authentication and authorization.
func PostAuthEvent(e *aah.Event) {
	ctx := e.Data.(*aah.Context)
	ctx.Log().Info("Method security.PostAuthEvent called")

	sub := ctx.Subject()
	ctx.Session().Set("Realm", sub.PrimaryPrincipal().Realm)
	ctx.Session().Set("Username", sub.PrimaryPrincipal().Value)
	ctx.Session().Set("ID", sub.Principal("ID").Value)
	ctx.Session().Set("Name", sub.Principal("Name").Value)
}
