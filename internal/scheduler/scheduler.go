package scheduler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/OpenCal-FYDP/CalendarEventManagement/internal/secretfetcher"
	"github.com/OpenCal-FYDP/CalendarEventManagement/internal/storage"
	"github.com/OpenCal-FYDP/Identity/rpc"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"net/http"
	"time"
)

const identityServiceUrl = "http://ec2-54-197-128-149.compute-1.amazonaws.com:8080"

type Scheduler struct {
	idServiceClient rpc.IdentityService
}

func New() *Scheduler {
	return &Scheduler{
		idServiceClient: rpc.NewIdentityServiceJSONClient(identityServiceUrl, &http.Client{}),
	}
}

func (s *Scheduler) getToken(eventOwnerEmail string, eventOwnerUsername string) (*oauth2.Token, error) {
	// get last stored token
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	res, err := s.idServiceClient.GetUser(ctx, &rpc.GetUserReq{Email: eventOwnerEmail, Username: eventOwnerUsername})
	if err != nil {
		return nil, err
	}

	tokenAsBytes := res.GetOathToken()
	if tokenAsBytes == nil {
		return nil, errors.New("nil token recieved from identity service")
	}

	type tstruct struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type,omitempty"`
		RefreshToken string `json:"refresh_token,omitempty"`
		Expiry       int64  `json:"expiry,omitempty"`
	}

	t := new(tstruct)

	token := &oauth2.Token{}
	err = json.Unmarshal(tokenAsBytes, t)
	if err != nil {
		return nil, err
	}

	token.AccessToken = t.AccessToken
	token.RefreshToken = t.RefreshToken
	token.TokenType = t.TokenType
	token.Expiry = time.Unix(t.Expiry, 0)

	return token, nil
}

func getCalService(token *oauth2.Token, eventOwnerEmail string, eventOwnerUsername string) (*calendar.Service, error) {
	// get oauthClientCreds
	// fetch these each time so we can change them on the fly
	res, err := secretfetcher.GetOauthConfig()
	if err != nil {
		return nil, err
	}

	// thing for client redirect
	config := oauth2.Config{
		ClientID:     res.OAuthClientID,
		ClientSecret: res.ClientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost",
		Scopes:       nil,
	}
	token.Expiry = time.Now().Add(time.Second - 1)

	// refresh token
	newToken, err := config.TokenSource(context.Background(), token).Token()
	if err != nil {
		return nil, err
	}

	client := config.Client(context.Background(), newToken)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	ctx = context.WithValue(ctx, oauth2.HTTPClient, client)

	return calendar.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
}

//gets a list of time availabilities
func (s *Scheduler) GetTeamEvents(teamID string) ([]string, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	idReq := &rpc.GetTeamReq{
		TeamID: teamID,
	}

	idRes, err := s.idServiceClient.GetTeam(ctx, idReq)
	if err != nil {
		return nil, err
	}

	retEvents := []string{}
	for _, memberEmail := range idRes.GetTeamMembers() {
		userEvents, err := s.GetUserEvents(memberEmail, memberEmail)
		if err != nil {
			return nil, err
		}
		retEvents = append(retEvents, userEvents...)
	}
	return retEvents, err
}

//gets a list of time availabilities
func (s *Scheduler) GetUserEvents(eventOwnerEmail string, eventOwnerUsername string) ([]string, error) {
	// data sanitize to default to jonathan
	if eventOwnerEmail == "" {
		eventOwnerEmail = "jspsun@gmail.com"
	}
	if eventOwnerUsername == "" {
		eventOwnerUsername = "jspsun@gmail.com"
	}

	// get oath token from user service
	token, err := s.getToken(eventOwnerEmail, eventOwnerUsername)
	if err != nil {
		return nil, err
	}

	srv, err := getCalService(token, eventOwnerEmail, eventOwnerUsername)
	if err != nil {
		return nil, err
	}

	t := time.Now().Add(-time.Hour * 12 * 7).Format(time.RFC3339)
	events, err := srv.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(60).OrderBy("startTime").Do()
	if err != nil {
		return nil, err
	}

	ret := []string{}
	for _, item := range events.Items {
		//fmt.Println(item.Summary)

		start, err := time.Parse(time.RFC3339, item.Start.DateTime)
		if err != nil {
			return nil, err
		}
		if item.End.DateTime == "" {
			start, err = time.Parse(time.RFC3339, item.Start.DateTime)
			if err != nil {
				return nil, err
			}
		}

		end, err := time.Parse(time.RFC3339, item.End.DateTime)
		if err != nil {
			return nil, err
		}
		if item.End.DateTime == "" {
			end, err = time.Parse(time.RFC3339, item.End.DateTime)
			if err != nil {
				return nil, err
			}
		}

		ret = append(ret, fmt.Sprintf("%d-%d", start.Unix(), end.Unix()))
	}

	return ret, nil
}

func (s *Scheduler) CreateEvent(eventOwnerEmail string, eventOwnerUsername string, data *storage.EventData) (string, error) {

	if data == nil {
		return "", errors.New("nil data")
	}

	// data sanitize to default to jonathan
	if eventOwnerEmail == "" {
		eventOwnerEmail = "jspsun@gmail.com"
	}
	if eventOwnerUsername == "" {
		eventOwnerUsername = "jspsun@gmail.com"
	}

	// get oath token from user service
	token, err := s.getToken(eventOwnerEmail, eventOwnerUsername)
	if err != nil {
		return "", err
	}

	srv, err := getCalService(token, eventOwnerEmail, eventOwnerUsername)
	if err != nil {
		return "", err
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
		ConferenceData: &calendar.ConferenceData{
			CreateRequest: &calendar.CreateConferenceRequest{
				RequestId: uuid.New().String(),
			},
		},
	}

	res, err := srv.Events.Insert("primary", calEvent).ConferenceDataVersion(1).Do()
	if err != nil {
		return "", err
	}
	return res.HtmlLink, nil
}

func (s *Scheduler) DeleteEvent(eventOwnerEmail string, eventOwnerUsername string, calEventID string) error {
	// data sanitize to default to jonathan
	if eventOwnerEmail == "" {
		eventOwnerEmail = "jspsun@gmail.com"
	}
	if eventOwnerUsername == "" {
		eventOwnerUsername = "jspsun@gmail.com"
	}

	// get oath token from user service
	token, err := s.getToken(eventOwnerEmail, eventOwnerUsername)
	if err != nil {
		return err
	}

	srv, err := getCalService(token, eventOwnerEmail, eventOwnerUsername)
	if err != nil {
		return err
	}

	err = srv.Events.Delete("primary", calEventID).Do()
	if err != nil {
		return err
	}
	return nil
}
