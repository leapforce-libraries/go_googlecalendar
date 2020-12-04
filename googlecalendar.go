package googlecalendar

import (
	"net/http"

	bigquerytools "github.com/leapforce-libraries/go_bigquerytools"
	errortools "github.com/leapforce-libraries/go_errortools"
	google "github.com/leapforce-libraries/go_google"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
)

const (
	apiName         string = "GoogleCalendar"
	apiURL          string = "https://www.googleapis.com/calendar/v3"
	authURL         string = "https://accounts.google.com/o/oauth2/v2/auth"
	tokenURL        string = "https://oauth2.googleapis.com/token"
	tokenHTTPMethod string = http.MethodPost
	redirectURL     string = "http://localhost:8080/oauth/redirect"
)

// GoogleDrive stores GoogleDrive configuration
//
type GoogleCalendar struct {
	oAuth2 *oauth2.OAuth2
}

// methods
//
func NewGoogleCalendar(clientID string, clientSecret string, scope string, bigQuery *bigquerytools.BigQuery) *GoogleCalendar {
	gd := GoogleCalendar{}
	config := oauth2.OAuth2Config{
		ApiName:         apiName,
		ClientID:        clientID,
		ClientSecret:    clientSecret,
		Scope:           scope,
		RedirectURL:     redirectURL,
		AuthURL:         authURL,
		TokenURL:        tokenURL,
		TokenHTTPMethod: tokenHTTPMethod,
	}
	gd.oAuth2 = oauth2.NewOAuth(config, bigQuery)
	return &gd
}

func (gc *GoogleCalendar) InitToken() *errortools.Error {
	return gc.oAuth2.InitToken()
}

func (gd *GoogleCalendar) Get(url string, model interface{}) (*http.Response, *errortools.Error) {
	err := google.ErrorResponse{}
	_, res, e := gd.oAuth2.Get(url, model, &err)

	if e != nil {
		if err.Error.Message != "" {
			e.SetMessage(err.Error.Message)
		}
		return nil, e
	}

	return res, nil
}
