package googlecalendar

import (
	"encoding/json"
	"fmt"
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

func (gd *GoogleCalendar) GetCalendarList() (*[]CalendarListEntry, error) {
	maxResults := 10
	pageToken := ""
	syncToken := ""

	calenderListEntries := []CalendarListEntry{}

	for syncToken == "" {
		queryPageToken := ""
		if pageToken != "" {
			queryPageToken = fmt.Sprintf("&pageToken=", pageToken)
		}
		url := fmt.Sprintf("%s/users/me/calendarList?maxResults=%v%s", apiURL, maxResults, queryPageToken)
		//fmt.Println(url)

		calendarListReponse := CalendarListResponse{}

		_, err := gd.Get(url, &calendarListReponse)
		if err != nil {
			return nil, err
		}

		calenderListEntries = append(calenderListEntries, calendarListReponse.Items...)

		pageToken = calendarListReponse.NextPageToken
		syncToken = calendarListReponse.NextSyncToken
	}

	return &calenderListEntries, nil
}
