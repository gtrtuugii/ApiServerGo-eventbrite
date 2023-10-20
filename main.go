package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Event struct {
	Body struct {
		Name struct {
			Text string `json:"text"`
			HTML string `json:"html"`
		} `json:"name"`
		Description struct {
			Text string `json:"text"`
			HTML string `json:"html"`
		} `json:"description"`
		URL   string `json:"url"`
		Start struct {
			Timezone string `json:"timezone"`
			Local    string `json:"local"`
			UTC      string `json:"utc"`
		} `json:"start"`
		End struct {
			Timezone string `json:"timezone"`
			Local    string `json:"local"`
			UTC      string `json:"utc"`
		} `json:"end"`
		OrganizationID               string      `json:"organization_id"`
		Created                      time.Time   `json:"created"`
		Changed                      time.Time   `json:"changed"`
		Published                    time.Time   `json:"published"`
		Capacity                     interface{} `json:"capacity"`
		CapacityIsCustom             interface{} `json:"capacity_is_custom"`
		Status                       string      `json:"status"`
		Currency                     string      `json:"currency"`
		Listed                       bool        `json:"listed"`
		Shareable                    bool        `json:"shareable"`
		OnlineEvent                  bool        `json:"online_event"`
		TxTimeLimit                  int         `json:"tx_time_limit"`
		HideStartDate                bool        `json:"hide_start_date"`
		HideEndDate                  bool        `json:"hide_end_date"`
		Locale                       string      `json:"locale"`
		IsLocked                     bool        `json:"is_locked"`
		PrivacySetting               string      `json:"privacy_setting"`
		IsSeries                     bool        `json:"is_series"`
		IsSeriesParent               bool        `json:"is_series_parent"`
		InventoryType                string      `json:"inventory_type"`
		IsReservedSeating            bool        `json:"is_reserved_seating"`
		ShowPickASeat                bool        `json:"show_pick_a_seat"`
		ShowSeatmapThumbnail         bool        `json:"show_seatmap_thumbnail"`
		ShowColorsInSeatmapThumbnail bool        `json:"show_colors_in_seatmap_thumbnail"`
		Source                       string      `json:"source"`
		IsFree                       bool        `json:"is_free"`
		Version                      interface{} `json:"version"`
		Summary                      string      `json:"summary"`
		FacebookEventID              interface{} `json:"facebook_event_id"`
		LogoID                       string      `json:"logo_id"`
		OrganizerID                  string      `json:"organizer_id"`
		VenueID                      string      `json:"venue_id"`
		CategoryID                   string      `json:"category_id"`
		SubcategoryID                interface{} `json:"subcategory_id"`
		FormatID                     string      `json:"format_id"`
		ID                           string      `json:"id"`
		ResourceURI                  string      `json:"resource_uri"`
		IsExternallyTicketed         bool        `json:"is_externally_ticketed"`
		Logo                         struct {
			CropMask struct {
				TopLeft struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"top_left"`
				Width  int `json:"width"`
				Height int `json:"height"`
			} `json:"crop_mask"`
			Original struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"original"`
			ID           string `json:"id"`
			URL          string `json:"url"`
			AspectRatio  string `json:"aspect_ratio"`
			EdgeColor    string `json:"edge_color"`
			EdgeColorSet bool   `json:"edge_color_set"`
		} `json:"logo"`
	} `json:"body"`
	Status string `json:"status"`
}

func getEvents(c *gin.Context) {

	oauthToken := "XXXXXXX"
	events := fetchEvents(oauthToken)
	c.JSON(http.StatusOK, events)
}

func getEventByID(c *gin.Context) {
	eventID := c.Param("id")

	oauthToken := "XXXXXXX"
	event := fetchEvent(oauthToken, eventID)
	c.JSON(http.StatusOK, event)
}

func createEvent(c *gin.Context) {
	// JSON payload for creating an event
	requestBody := []byte(`
        {
            "event": {
                "name": {
                    "html": "<p>Some text</p>"
                },
                "description": {
                    "html": "<p>Some text</p>"
                },
                "start": {
                    "timezone": "UTC",
                    "utc": "2018-05-12T02:00:00Z"
                },
                "end": {
                    "timezone": "UTC",
                    "utc": "2018-05-12T02:00:00Z"
                },
                "currency": "USD",
                "online_event": false,
                "organizer_id": "",
                "listed": false,
                "shareable": false,
                "invite_only": false,
                "show_remaining": true,
                "password": "12345",
                "capacity": 100,
                "is_reserved_seating": true,
                "is_series": true,
                "show_pick_a_seat": true,
                "show_seatmap_thumbnail": true,
                "show_colors_in_seatmap_thumbnail": true,
                "locale": "de_AT"
            }
        }`)

	// Make a POST request to create the event
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://www.eventbriteapi.com/v3/organizations/12345/events/", bytes.NewBuffer(requestBody))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the event"})
		return
	}

	req.Header.Add("Authorization", "Bearer XXXX")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the event"})
		return
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)

	c.JSON(http.StatusOK, gin.H{
		"status": resp.Status,
		"body":   string(respBody),
	})
}

func updateEvent(c *gin.Context) {
	// Retrieve the event ID from the URI parameter
	eventID := c.Param("id")

	// JSON payload for updating an event
	requestBody := []byte(`
        {
            "event": {
                "name": {
                    "html": "<p>Updated text</p>"
                },
                "description": {
                    "html": "<p>Updated text</p>"
                },
                "start": {
                    "timezone": "UTC",
                    "utc": "2018-05-12T02:00:00Z"
                },
                "end": {
                    "timezone": "UTC",
                    "utc": "2018-05-12T02:00:00Z"
                },
                "currency": "USD",
                "online_event": false,
                "organizer_id": "",
                "listed": false,
                "shareable": false,
                "invite_only": false,
                "show_remaining": true,
                "password": "12345",
                "capacity": 100,
                "is_reserved_seating": true,
                "is_series": true,
                "show_pick_a_seat": true,
                "show_seatmap_thumbnail": true,
                "show_colors_in_seatmap_thumbnail": true
            }
        }`)

	// Make a POST request to update the event with the dynamic event ID
	client := &http.Client{}
	url := fmt.Sprintf("https://www.eventbriteapi.com/v3/events/%s/", eventID)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the event"})
		return
	}

	req.Header.Add("Authorization", "Bearer XXXXX")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the event"})
		return
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)

	c.JSON(http.StatusOK, gin.H{
		"status": resp.Status,
		"body":   string(respBody),
	})
}

func fetchEvents(oauthToken string) []gin.H {
	url := "https://www.eventbrite.com/api/v3/events/"

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+oauthToken)

	resp, err := client.Do(req)
	if err != nil {
		return []gin.H{{"error": "Failed to fetch events"}}
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)

	return []gin.H{
		{"status": resp.Status},
		{"body": string(respBody)},
	}
}

func fetchEvent(oauthToken, eventID string) gin.H {
	url := fmt.Sprintf("https://www.eventbrite.com/api/v3/events/%s/", eventID)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+oauthToken)

	resp, err := client.Do(req)
	if err != nil {
		return gin.H{"error": "Failed to fetch the event"}
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)

	return gin.H{
		"status": resp.Status,
		"body":   string(respBody),
	}
}

func main() {
	router := gin.Default()

	router.GET("/events", getEvents)
	router.GET("/events/:id", getEventByID)
	router.POST("/createEvent", createEvent)
	router.POST("/events/:id", updateEvent)
	// TODO:
	// Cancel event
	// Delete event
	router.Run("localhost:8080")
}
