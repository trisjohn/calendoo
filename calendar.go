package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	calendar "google.golang.org/api/calendar/v3"
)

// tokenFromFile retrieves a Token from a given file path.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// getTokenFromWeb requests a Token from the web, then returns the retrieved Token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// saveToken saves a Token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func getService() *calendar.Service {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarEventsScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}
	return srv
}

func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

type EventConfig struct {
	Summary     string
	Location    string
	Description string
	Start       EventDateTime
	End         EventDateTime
	Recurrence  []string
	Attendees   []EventAttendee
}

type EventDateTime struct {
	DateTime string
	TimeZone string
}

type EventAttendee struct {
	Email string
}

func createEvent(config *EventConfig, id string) {
	// Map EventConfig to google calendar event struct
	event := &calendar.Event{
		Summary:     config.Summary,
		Location:    config.Location,
		Description: config.Description,
		Start: &calendar.EventDateTime{
			DateTime: config.Start.DateTime,
			TimeZone: config.Start.TimeZone,
		},
		End: &calendar.EventDateTime{
			DateTime: config.End.DateTime,
			TimeZone: config.End.TimeZone,
		},
		Recurrence: config.Recurrence,
		Attendees:  []*calendar.EventAttendee{},
	}

	// Map EventAttendee to google calendar event attendee struct
	for _, attendee := range config.Attendees {
		event.Attendees = append(event.Attendees, &calendar.EventAttendee{
			Email: attendee.Email,
		})
	}

	srv := getService()
	// Insert event to primary calendar
	calendarId := id
	event, err := srv.Events.Insert(calendarId, event).Do()
	if err != nil {
		log.Fatalf("Unable to create event. %v\n", err)
	}

	log.Printf("Event created: %s\n", event.HtmlLink)
}
