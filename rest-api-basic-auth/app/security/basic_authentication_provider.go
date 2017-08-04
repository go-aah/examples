package security

import (
	"aahframework.org/config.v0"
	"aahframework.org/security.v0/authc"
)

var _ authc.Authenticator = (*BasicAuthenticationProvider)(nil)

// BasicAuthenticationProvider struct implements `authc.Authenticator` interface.
type BasicAuthenticationProvider struct {
	// for demo purpose in-memory subject (aka user) info's
	// Typically you will be using Database, API calls, LDAP, etc to get the Authentication
	// Information.
	users map[string]authc.AuthenticationInfo
}

// Init method initializes the BasicAuthenticationProvider, this method gets called
// during server start up.
func (fa *BasicAuthenticationProvider) Init(cfg *config.Config) error {

	// NOTE: for demo purpose I'm creating set users in the map.
	// Typically you will be using Database, API calls, LDAP, etc to get the Authentication
	// Information.

	fa.users = make(map[string]authc.AuthenticationInfo)

	// Subject 1
	authcInfo1 := authc.NewAuthenticationInfo()
	authcInfo1.Principals = append(authcInfo1.Principals,
		&authc.Principal{Value: "user1@example.com", IsPrimary: true, Realm: "inmemory"})
	authcInfo1.Credential = []byte(`$2y$10$2A4GsJ6SmLAMvDe8XmTam.MSkKojdobBVJfIU7GiyoM.lWt.XV3H6`) // welcome123
	fa.users["user1@example.com"] = *authcInfo1

	// Subject 2
	authcInfo2 := authc.NewAuthenticationInfo()
	authcInfo2.Principals = append(authcInfo2.Principals,
		&authc.Principal{Value: "user2@example.com", IsPrimary: true, Realm: "inmemory"})
	authcInfo2.Credential = []byte(`$2y$10$2A4GsJ6SmLAMvDe8XmTam.MSkKojdobBVJfIU7GiyoM.lWt.XV3H6`) // welcome123
	fa.users["user2@example.com"] = *authcInfo2

	// Subject 3 - user is locked state
	authcInfo3 := authc.NewAuthenticationInfo()
	authcInfo3.Principals = append(authcInfo3.Principals,
		&authc.Principal{Value: "user3@example.com", IsPrimary: true, Realm: "inmemory"})
	authcInfo3.Credential = []byte(`$2y$10$2A4GsJ6SmLAMvDe8XmTam.MSkKojdobBVJfIU7GiyoM.lWt.XV3H6`) // welcome123
	authcInfo3.IsLocked = true
	fa.users["user3@example.com"] = *authcInfo3

	return nil
}

// GetAuthenticationInfo method is `authc.Authenticator` interface
func (fa *BasicAuthenticationProvider) GetAuthenticationInfo(authcToken *authc.AuthenticationToken) (*authc.AuthenticationInfo, error) {

	if ai, found := fa.users[authcToken.Identity]; found {
		return &ai, nil
	}

	// No subject exists, return nil and error
	return nil, authc.ErrSubjectNotExists
}
