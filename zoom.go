package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
    "github.com/joho/godotenv"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

type Meeting struct {
	Topic     string `json:"topic"`
	Type      int    `json:"type"`
	StartTime string `json:"start_time"`
	Duration  int    `json:"duration"`
}

func CreateMeeting(meeting *Meeting) {
	err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }
	
    clientID := os.Getenv("ZOOM_CLIENT")
    clientSecret := os.Getenv("ZOOM_SECRET")
    userID := os.Getenv("ZOOM_ACCOUNT")

	// Get access token
	data := fmt.Sprintf(`{"client_id": "%s", "client_secret": "%s", "grant_type": "client_credentials"}`, clientID, clientSecret)
	resp, err := http.Post("https://zoom.us/oauth/token", "application/json", bytes.NewBuffer([]byte(data)))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var tokenResponse TokenResponse
	json.Unmarshal(body, &tokenResponse)
	meetingData, _ := json.Marshal(meeting)

	req, _ := http.NewRequest("POST", fmt.Sprintf("https://api.zoom.us/v2/users/%s/meetings", userID), bytes.NewBuffer(meetingData))
	req.Header.Set("Authorization", "Bearer "+tokenResponse.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ = ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
