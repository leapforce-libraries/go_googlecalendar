package googlecalendar

import (
	"net/http"

	bigquerytools "github.com/Leapforce-nl/go_bigquerytools"

	oauth2 "github.com/Leapforce-nl/go_oauth2"
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
func NewGoogleCalendar(division int, clientID string, clientSecret string, bigQuery *bigquerytools.BigQuery, isLive bool) (*GoogleCalendar, error) {
	gd := GoogleCalendar{}
	config := oauth2.OAuth2Config{
		ApiName:         apiName,
		ClientID:        clientID,
		ClientSecret:    clientSecret,
		RedirectURL:     redirectURL,
		AuthURL:         authURL,
		TokenURL:        tokenURL,
		TokenHTTPMethod: tokenHTTPMethod,
	}
	gd.oAuth2 = oauth2.NewOAuth(config, bigQuery, isLive)
	return &gd, nil
}

func (gd *GoogleCalendar) Get(url string, model interface{}) (*http.Response, error) {
	res, err := gd.oAuth2.Get(url, model)

	if err != nil {
		return nil, err
	}

	return res, nil
}
