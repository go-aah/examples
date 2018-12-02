package facebook

import (
	"context"
	"fmt"
	"net/http"

	"aahframe.work"
	"aahframe.work/config"
	"aahframe.work/essentials"
	"aahframe.work/security/authc"
	"aahframe.work/security/scheme"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var _ authc.PrincipalProvider = (*SubjectPrincipalProvider)(nil)

// SubjectPrincipalProvider struct implements interface authc.PrincipalProvider.
type SubjectPrincipalProvider struct{}

// Init method initializes the SubjectPrincipalProvider, Init method gets called
// during aah server start up.
func (a *SubjectPrincipalProvider) Init(appCfg *config.Config) error {
	// NOTE: Init is called on application startup
	return nil
}

// Principals method is called to get Subject's Principals information.
func (a *SubjectPrincipalProvider) Principal(keyName string, v ess.Valuer) ([]*authc.Principal, error) {
	token := v.Get(aah.KeyOAuth2Token).(*oauth2.Token)
	githubAuth := aah.App().SecurityManager().AuthScheme(keyName).(*scheme.OAuth2)
	githubClient := github.NewClient(githubAuth.Client(token))

	userInfo, resp, err := githubClient.Users.Get(context.Background(), "")
	if err != nil || resp.StatusCode != http.StatusOK || userInfo == nil || userInfo.ID == nil {
		return nil, fmt.Errorf("unable to get user info from google")
	}

	principals := make([]*authc.Principal, 0)
	principals = append(principals, &authc.Principal{Realm: "GitHub", IsPrimary: true, Claim: "Email", Value: userInfo.GetEmail()})
	principals = append(principals, &authc.Principal{Realm: "GitHub", Claim: "ID", Value: fmt.Sprintf("%v", userInfo.GetID())})
	principals = append(principals, &authc.Principal{Realm: "GitHub", Claim: "Name", Value: userInfo.GetName()})

	return principals, nil
}
