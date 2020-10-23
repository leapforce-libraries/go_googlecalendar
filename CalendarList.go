package googlecalendar

import (
	"fmt"
)

type CalendarListResponse struct {
	Kind          string              `json:"kind"`
	Etag          string              `json:"etag"`
	NextPageToken string              `json:"nextPageToken"`
	NextSyncToken string              `json:"nextSyncToken"`
	Items         []CalendarListEntry `json:"calendarList"`
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
	DefaultReminders     []string           `json:"defaultReminders"`
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

		url := fmt.Sprintf("%s/calendarList?maxResults=%v%%s", apiURL, maxResults)
		fmt.Println(url)

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
