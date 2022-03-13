package main

import (
	"github.com/OpenCal-FYDP/CalendarEventManagement/internal/service"
	"github.com/OpenCal-FYDP/CalendarEventManagement/rpc"
	"log"
	"net/http"
)

func main() {
	svc := service.New()
	server := rpc.NewCalendarEventManagementServiceServer(svc)
	log.Fatal(http.ListenAndServe(":8080", server))
}

// uncomment and run to manually get new oauth token
//
//// Retrieve a token, saves the token, then returns the generated client.
//func getClient(config *oauth2.Config) *http.Client {
//	// The file token.json stores the user's access and refresh tokens, and is
//	// created automatically when the authorization flow completes for the first
//	// time.
//	tokFile := "token.json"
//	tok, err := tokenFromFile(tokFile)
//	if err != nil {
//		tok = getTokenFromWeb(config)
//		saveToken(tokFile, tok)
//	}
//	return config.Client(context.Background(), tok)
//}
//
//// Request a token from the web, then returns the retrieved token.
//func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
//	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
//	fmt.Printf("Go to the following link in your browser then type the "+
//		"authorization code: \n%v\n", authURL)
//
//	var authCode string
//	if _, err := fmt.Scan(&authCode); err != nil {
//		log.Fatalf("Unable to read authorization code: %v", err)
//	}
//
//	tok, err := config.Exchange(context.TODO(), authCode)
//	if err != nil {
//		log.Fatalf("Unable to retrieve token from web: %v", err)
//	}
//	return tok
//}
//
//// Retrieves a token from a local file.
//func tokenFromFile(file string) (*oauth2.Token, error) {
//	f, err := os.Open(file)
//	if err != nil {
//		return nil, err
//	}
//	defer f.Close()
//	tok := &oauth2.Token{}
//	err = json.NewDecoder(f).Decode(tok)
//	return tok, err
//}
//
//// Saves a token to a file path.
//func saveToken(path string, token *oauth2.Token) {
//	fmt.Printf("Saving credential file to: %s\n", path)
//	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
//	if err != nil {
//		log.Fatalf("Unable to cache oauth token: %v", err)
//	}
//	defer f.Close()
//	json.NewEncoder(f).Encode(token)
//}
//
//func main() {
//	ctx := context.Background()
//	b, err := ioutil.ReadFile("credentials.json")
//	if err != nil {
//		log.Fatalf("Unable to read client secret file: %v", err)
//	}
//
//	// If modifying these scopes, delete your previously saved token.json.
//	config, err := google.ConfigFromJSON(b, calendar.CalendarEventsScope)
//	if err != nil {
//		log.Fatalf("Unable to parse client secret file to config: %v", err)
//	}
//	client := getClient(config)
//
//	_, err = calendar.NewService(ctx, option.WithHTTPClient(client))
//	if err != nil {
//		log.Fatalf("Unable to retrieve Calendar client: %v", err)
//	}
//	//
//	//ti := time.Now().Format(time.RFC3339)
//	//events, err := srv.Events.List("primary").ShowDeleted(false).
//	//	SingleEvents(true).TimeMin(ti).MaxResults(10).OrderBy("startTime").Do()
//	//if err != nil {
//	//	log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
//	//}
//	//fmt.Println("Upcoming events:")
//	//if len(events.Items) == 0 {
//	//	fmt.Println("No upcoming events found.")
//	//} else {
//	//	for _, item := range events.Items {
//	//		date := item.Start.DateTime
//	//		if date == "" {
//	//			date = item.Start.Date
//	//		}
//	//		fmt.Printf("%v (%v)\n", item.Summary, date)
//	//	}
//	//}
//
//	//calEvent := &calendar.Event{
//	//	Summary: "dank times",
//	//	Location: "a dank address",
//	//	Description: "wooooEvents",
//	//	Start: &calendar.EventDateTime{
//	//		DateTime: time.Now().Format(time.RFC3339),
//	//		//TimeZone: "America/Los_Angeles",
//	//	},
//	//	End: &calendar.EventDateTime{
//	//		DateTime: time.Now().Add(time.Hour).Format(time.RFC3339),
//	//		//TimeZone: "America/Los_Angeles",
//	//	},
//	//	//Recurrence: []string{"RRULE:FREQ=DAILY;COUNT=2"},
//	//	//Attendees: []*calendar.EventAttendee{
//	//	//	&calendar.EventAttendee{Email:"lpage@example.com"},
//	//	//	&calendar.EventAttendee{Email:"sbrin@example.com"},
//	//	//},
//	//}
//	//_, err = srv.Events.Insert("primary", calEvent).Do()
//	//if err != nil {
//	//	log.Fatalf("Unable to parse client secret file to config: %v", err)
//	//}
//}
