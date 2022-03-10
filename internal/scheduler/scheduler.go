package scheduler

import (
	"context"
	"encoding/json"
	"github.com/OpenCal-FYDP/CalendarEventManagement/internal/storage"
	"github.com/OpenCal-FYDP/Identity/rpc"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"net/http"
	"time"
)

const identityServiceUrl = "http://ec2-54-82-78-138.compute-1.amazonaws.com:8080"

type Scheduler struct {
}

func (s *Scheduler) CreateEvent(eventOwnerEmail string, eventOwnerUsername string, data *storage.EventData) error {
	// data sanitize to default to jonathan
	if eventOwnerEmail == "" {
		eventOwnerEmail = "jspsun@gmail.com"
	}
	if eventOwnerUsername == "" {
		eventOwnerUsername = "jspsun"
	}

	config := &oauth2.Config{}

	idClient := rpc.NewIdentityServiceJSONClient(identityServiceUrl, &http.Client{})
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	res, err := idClient.GetUser(ctx, &rpc.GetUserReq{Email: eventOwnerEmail, Username: eventOwnerUsername})
	if err != nil {
		return err
	}

	tokenAsBytes := res.GetOathToken()

	token := &oauth2.Token{}
	err = json.Unmarshal([]byte(tokenAsBytes), token)
	if err != nil {
		return err
	}

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
