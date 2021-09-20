package googlecalendar

import (
	"encoding/json"
	"fmt"
	"time"

	"cloud.google.com/go/civil"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type EventsResponse struct {
	Kind             string            `json:"kind"`
	Etag             string            `json:"etag"`
	Summary          string            `json:"summary"`
	Updated          string            `json:"updated"`
	TimeZone         string            `json:"timeZone"`
	AccessRole       string            `json:"accessRole"`
	DefaultReminders []json.RawMessage `json:"defaultReminders"`
	NextPageToken    string            `json:"nextPageToken"`
	NextSyncToken    string            `json:"nextSyncToken"`
	Items            []Event           `json:"items"`
}

type Event struct {
	Kind               string          `json:"kind"`
	Etag               string          `json:"etag"`
	ID                 string          `json:"id"`
	Status             string          `json:"status"`
	HTMLLink           string          `json:"htmlLink"`
	Created            *time.Time      `json:"created"`
	Updated            *time.Time      `json:"updated"`
	Summary            string          `json:"summary"`
	Description        string          `json:"description"`
	Location           string          `json:"location"`
	ColorID            string          `json:"colorId"`
	Creator            Attendee        `json:"creator"`
	Organizer          Attendee        `json:"organizer"`
	Start              StartEndTime    `json:"start"`
	End                StartEndTime    `json:"end"`
	EndTimeUnspecified bool            `json:"endTimeUnspecified"`
	RecurringEventId   json.RawMessage `json:"recurringEventId"`
	OriginalStartTime  StartEndTime    `json:"originalStartTime"`
	Transparency       string          `json:"transparency"`
	Visibility         string          `json:"visibility"`
	ICalUID            string          `json:"iCalUID"`
	Sequence           int64           `json:"sequence"`
	Attendees          []Attendee      `json:"attendees"`
	AttendeesOmitted   bool            `json:"attendeesOmitted"`
	ExtendedProperties struct {
		Private map[string]string `json:"private"`
		Shared  map[string]string `json:"shared"`
	} `json:"extendedProperties"`
	HangoutLink string `json:"hangoutLink"`
	//ConferenceData string `json:"conferenceData"`
	//Gadget string `json:"gadget"`
	AnyoneCanAddSelf        bool         `json:"anyoneCanAddSelf"`
	GuestsCanInviteOthers   bool         `json:"guestsCanInviteOthers"`
	GuestsCanModify         bool         `json:"guestsCanModify"`
	GuestsCanSeeOtherGuests bool         `json:"guestsCanSeeOtherGuests"`
	PrivateCopy             bool         `json:"privateCopy"`
	Locked                  bool         `json:"locked"`
	Reminders               Reminders    `json:"reminders"`
	Source                  Source       `json:"source"`
	Attachments             []Attachment `json:"attachments"`
}

type Attendee struct {
	ID               string `json:"id"`
	Email            string `json:"email"`
	DisplayName      string `json:"displayName"`
	Organizer        bool   `json:"organizer"`
	Self             bool   `json:"self"`
	Resource         bool   `json:"resource"`
	Optional         bool   `json:"optional"`
	ResponseStatus   string `json:"responseStatus"`
	Comment          string `json:"comment"`
	AdditionalGuests int    `json:"additionalGuests"`
}

type StartEndTime struct {
	Date     string     `json:"date"`
	DateTime *time.Time `json:"dateTime"`
	TimeZone string     `json:"timeZone"`
}

type Reminders struct {
	UseDefault bool       `json:"useDefault"`
	Overrides  []Reminder `json:"overrides"`
}

type Reminder struct {
	Method  string `json:"method"`
	Minutes int    `json:"minutes"`
}

type Source struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

type Attachment struct {
	FileURL  string `json:"fileUrl"`
	Title    string `json:"title"`
	MimeType string `json:"mimeType"`
	IconLink string `json:"iconLink"`
	FileID   string `json:"fileId"`
}

func (service *Service) GetEvents(calendarID string, timeMin *civil.Date) (*[]Event, *errortools.Error) {
	maxResults := 10
	pageToken := ""
	syncToken := ""

	events := []Event{}

	for syncToken == "" {
		queryPageToken := ""
		if pageToken != "" {
			queryPageToken = fmt.Sprintf("&pageToken=%s", pageToken)
		}
		timeMin_ := ""
		if timeMin != nil {
			timeMin_ = fmt.Sprintf("&timeMin=%s%s", timeMin.String(), "T00:00:00.000Z")
		}

		eventsReponse := EventsResponse{}

		requestConfig := go_http.RequestConfig{
			URL:           service.url(fmt.Sprintf("calendars/%s/events?maxResults=%v%s%s", calendarID, maxResults, queryPageToken, timeMin_)),
			ResponseModel: &eventsReponse,
		}
		_, _, e := service.googleService.Get(&requestConfig)
		if e != nil {
			return nil, e
		}

		events = append(events, eventsReponse.Items...)

		pageToken = eventsReponse.NextPageToken
		syncToken = eventsReponse.NextSyncToken
	}

	return &events, nil
}
