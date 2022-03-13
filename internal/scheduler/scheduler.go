package scheduler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/OpenCal-FYDP/CalendarEventManagement/internal/storage"
	"github.com/OpenCal-FYDP/Identity/rpc"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"net/http"
	"time"
)

const identityServiceUrl = "http://ec2-54-147-167-136.compute-1.amazonaws.com:8080"

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

	config := oauth2.Config{}
	client := config.Client(context.Background(), token)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	ctx = context.WithValue(ctx, oauth2.HTTPClient, client)

	srv, err := calendar.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
	if err != nil {
		return err
	}

	calEvent := &calendar.Event{
		Summary:     data.Summary,
		Location:    data.Location,
		Description: "meeting scheduled with openCal",
		Start: &calendar.EventDateTime{
			DateTime: time.Unix(data.Start, 0).Format(time.RFC3339),
			//TimeZone: "America/Los_Angeles",
		},
		End: &calendar.EventDateTime{
			DateTime: time.Unix(data.End, 0).Format(time.RFC3339),
			//TimeZone: "America/Los_Angeles",
		},
		Id: data.CalendarEventID,
		//Recurrence: []string{"RRULE:FREQ=DAILY;COUNT=2"},
		//Attendees: []*calendar.EventAttendee{
		//	&calendar.EventAttendee{Email:"lpage@example.com"},
		//	&calendar.EventAttendee{Email:"sbrin@example.com"},
		//},
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
