package scheduler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/OpenCal-FYDP/CalendarEventManagement/internal/secretfetcher"
	"github.com/OpenCal-FYDP/CalendarEventManagement/internal/storage"
	"github.com/OpenCal-FYDP/Identity/rpc"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"net/http"
	"time"
)

const identityServiceUrl = "ec2-54-197-128-149.compute-1.amazonaws.com:8080"

type Scheduler struct {
}

func New() *Scheduler {
	return &Scheduler{}
}

func getToken(eventOwnerEmail string, eventOwnerUsername string) (*oauth2.Token, error) {
	// get last stored token
	idClient := rpc.NewIdentityServiceJSONClient(identityServiceUrl, &http.Client{})
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	res, err := idClient.GetUser(ctx, &rpc.GetUserReq{Email: eventOwnerEmail, Username: eventOwnerUsername})
	if err != nil {
		return nil, err
	}

	tokenAsBytes := res.GetOathToken()
	if tokenAsBytes == nil {
		return nil, errors.New("nil token recieved from identity service")
	}

	token := &oauth2.Token{}
	err = json.Unmarshal(tokenAsBytes, token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *Scheduler) CreateEvent(eventOwnerEmail string, eventOwnerUsername string, data *storage.EventData) error {

	if data == nil {
		return errors.New("nil data")
	}

	// data sanitize to default to jonathan
	if eventOwnerEmail == "" {
		eventOwnerEmail = "jspsun@gmail.com"
	}
	if eventOwnerUsername == "" {
		eventOwnerUsername = "jspsun@gmail.com"
	}

	// get oath token from user service
	token, err := getToken(eventOwnerEmail, eventOwnerUsername)
	if err != nil {
		return err
	}

	// get oauthClientCreds
	// fetch these each time so we can change them on the fly
	res, err := secretfetcher.GetOauthConfig()
	if err != nil {
		return err
	}

	// thing for client redirect
	config := oauth2.Config{
		ClientID:     res.OAuthClientID,
		ClientSecret: res.ClientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost",
		Scopes:       nil,
	}

	// refresh token
	newToken, err := config.TokenSource(context.Background(), token).Token()

	client := config.Client(context.Background(), newToken)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	ctx = context.WithValue(ctx, oauth2.HTTPClient, client)

	srv, err := calendar.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
	if err != nil {
		return err
	}

	// construct attendees
	atts := []*calendar.EventAttendee{}

	for _, attendee := range data.Attendees {
		atts = append(atts, &calendar.EventAttendee{Email: attendee})
	}

	// construct cal event
	calEvent := &calendar.Event{
		Summary:     data.Summary,
		Location:    data.Location,
		Description: "meeting scheduled with openCal",
		Start: &calendar.EventDateTime{
			DateTime: time.Unix(data.Start, 0).UTC().Format(time.RFC3339),
		},
		End: &calendar.EventDateTime{
			DateTime: time.Unix(data.End, 0).UTC().Format(time.RFC3339),
		},
		Id: data.CalendarEventID,
		//Recurrence: []string{"RRULE:FREQ=DAILY;COUNT=2"},
		Attendees: atts,
	}

	_, err = srv.Events.Insert("primary", calEvent).Do()
	if err != nil {
		return err
	}
	return nil
}

func (s *Scheduler) DeleteEvent(eventOwnerEmail string, eventOwnerUsername string, data *storage.EventData) error {
	return nil
}
