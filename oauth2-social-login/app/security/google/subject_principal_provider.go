package facebook

import (
	"fmt"

	"aahframe.work"
	"aahframe.work/config"
	"aahframe.work/essentials"
	"aahframe.work/security/authc"
	"aahframe.work/security/scheme"
	"golang.org/x/oauth2"
	google "google.golang.org/api/oauth2/v2"
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
	googleAuth := aah.App().SecurityManager().AuthScheme(keyName).(*scheme.OAuth2)
	googleService, err := google.New(googleAuth.Client(token))
	if err != nil {
		return nil, fmt.Errorf("unable to create google service")
	}

	userInfo, err := googleService.Userinfo.Get().Do()
	if err != nil || userInfo == nil || userInfo.Id == "" {
		return nil, fmt.Errorf("unable to get user info from google")
	}

	principals := make([]*authc.Principal, 0)
	principals = append(principals, &authc.Principal{Realm: "Google", IsPrimary: true, Claim: "Email", Value: userInfo.Email})
	principals = append(principals, &authc.Principal{Realm: "Google", Claim: "ID", Value: userInfo.Id})
	principals = append(principals, &authc.Principal{Realm: "Google", Claim: "Name", Value: userInfo.Name})

	return principals, nil
}
