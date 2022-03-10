package scheduler

import (
	"github.com/OpenCal-FYDP/CalendarEventManagement/internal/storage"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestScheduler(t *testing.T) {
	//	config := &oauth2.Config{}
	//
	//	idClient := rpc.NewIdentityServiceJSONClient("http://ec2-54-82-78-138.compute-1.amazonaws.com:8080", &http.Client{})
	//	ctx, _ := context.WithTimeout(context.Background(), time.Second *5 )
	//	res, err := idClient.GetUser(ctx, &rpc.GetUserReq{Email: "jspsun@gmail.com", Username: "jspsun"})
	//	require.NoError(t, err)
	//	require.NotNil(t, res)
	//
	//	tokenAsBytes := res.GetOathToken()
	//
	//
	//	token := &oauth2.Token{}
	//	err = json.Unmarshal([]byte(tokenAsBytes), token)
	//	require.NoError(t, err)
	//
	//	srv, err := calendar.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
	//	require.NoError(t, err)
	//	require.NotNil(t, srv)
	//
	//	calEvent := &calendar.Event{
	//		Summary: "dank times",
	//		Location: "a dank address",
	//		Description: "wooooEvents",
	//		Start: &calendar.EventDateTime{
	//			DateTime: time.Now().Format(time.RFC3339),
	//			//TimeZone: "America/Los_Angeles",
	//		},
	//		End: &calendar.EventDateTime{
	//			DateTime: time.Now().Add(time.Hour).Format(time.RFC3339),
	//			//TimeZone: "America/Los_Angeles",
	//		},
	//		//Recurrence: []string{"RRULE:FREQ=DAILY;COUNT=2"},
	//		//Attendees: []*calendar.EventAttendee{
	//		//	&calendar.EventAttendee{Email:"lpage@example.com"},
	//		//	&calendar.EventAttendee{Email:"sbrin@example.com"},
	//		//},
	//	}
	//
	//	e, err := srv.Events.Insert("primary", calEvent).Do()
	//	require.NoError(t, err)
	//	require.NotNil(t, e)
	//
	//////////
	////	ti := time.Now().Format(time.RFC3339)
	////	events, err := srv.Events.List("primary").ShowDeleted(false).
	////		SingleEvents(true).TimeMin(ti).MaxResults(10).OrderBy("startTime").Do()
	////	if err != nil {
	////		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	////	}
	////	fmt.Println("Upcoming events:")
	////	if len(events.Items) == 0 {
	////		fmt.Println("No upcoming events found.")
	////	} else {
	////		for _, item := range events.Items {
	////			date := item.Start.DateTime
	////			if date == "" {
	////				date = item.Start.Date
	////			}
	////			fmt.Printf("%v (%v)\n", item.Summary, date)
	////		}
	////	}
	e := &storage.EventData{
		CalendarEventID: "",
		Start:           time.Now().Unix(),
		End:             time.Now().Add(time.Hour).Unix(),
		Attendees:       nil,
		Location:        "somewhere",
		Summary:         "dankSummary",
	}

	t.Run("CreateEvent", func(t *testing.T) {
		s := &Scheduler{}
		err := s.CreateEvent("", "", e)
		assert.NoError(t, err)
	})
}
