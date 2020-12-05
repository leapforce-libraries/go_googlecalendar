package googlecalendar

import (
	bigquerytools "github.com/leapforce-libraries/go_bigquerytools"
	google "github.com/leapforce-libraries/go_google"
)

const (
	apiName string = "GoogleCalendar"
	apiURL  string = "https://www.googleapis.com/calendar/v3"
)

// GoogleCalendar stores GoogleCalendar configuration
//
type GoogleCalendar struct {
	Client *google.GoogleClient
}

// methods
//
func NewGoogleCalendar(clientID string, clientSecret string, scope string, bigQuery *bigquerytools.BigQuery) *GoogleCalendar {
	config := google.GoogleClientConfig{
		APIName:      apiName,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scope:        scope,
	}

	googleClient := google.NewGoogleClient(config, bigQuery)

	return &GoogleCalendar{googleClient}
}
