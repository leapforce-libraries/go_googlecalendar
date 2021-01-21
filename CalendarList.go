package googlecalendar

import (
	"encoding/json"
	"fmt"

	errortools "github.com/leapforce-libraries/go_errortools"
	oauth2 "github.com/leapforce-libraries/go_oauth2"
)

type CalendarListResponse struct {
	Kind          string              `json:"kind"`
	Etag          string              `json:"etag"`
	NextPageToken string              `json:"nextPageToken"`
	NextSyncToken string              `json:"nextSyncToken"`
	Items         []CalendarListEntry `json:"items"`
}

type CalendarListEntry struct {
	Kind                 string             `json:"kind"`
	Etag                 string             `json:"etag"`
	ID                   string             `json:"id"`
	Summary              string             `json:"summary"`
	TimeZone             string             `json:"timeZone"`
	ColorID              string             `json:"colorId"`
	BackgroundColor      string             `json:"backgroundColor"`
	ForegroundColor      string             `json:"foregroundColor"`
	AccessRole           string             `json:"accessRole"`
	DefaultReminders     []json.RawMessage  `json:"defaultReminders"`
	ConferenceProperties ConferenceProperty `json:"conferenceProperties"`
}

type ConferenceProperty struct {
	AllowedConferenceSolutionTypes []string `json:"allowedConferenceSolutionTypes"`
}

func (service *Service) GetCalendarList() (*[]CalendarListEntry, *errortools.Error) {
	maxResults := 10
	pageToken := ""
	syncToken := ""

	calenderListEntries := []CalendarListEntry{}

	for syncToken == "" {
		queryPageToken := ""
		if pageToken != "" {
			queryPageToken = fmt.Sprintf("&pageToken=%s", pageToken)
		}

		calendarListReponse := CalendarListResponse{}

		requestConfig := oauth2.RequestConfig{
			URL:           service.url(fmt.Sprintf("users/me/calendarList?maxResults=%v%s", maxResults, queryPageToken)),
			ResponseModel: &calendarListReponse,
		}
		_, _, e := service.googleService.Get(&requestConfig)
		if e != nil {
			return nil, e
		}

		calenderListEntries = append(calenderListEntries, calendarListReponse.Items...)

		pageToken = calendarListReponse.NextPageToken
		syncToken = calendarListReponse.NextSyncToken
	}

	return &calenderListEntries, nil
}
