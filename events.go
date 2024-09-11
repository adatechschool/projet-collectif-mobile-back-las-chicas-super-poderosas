package main

import (
	"encoding/json"
	"net/http"
)

// Définition des structures Go pour les données JSON

// Thumbnail représente les différentes tailles de vignettes d'une photo
type Thumbnail struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// Photo contient les informations d'une photo et ses vignettes
type Photo struct {
	ID         string              `json:"id"`
	Width      int                 `json:"width"`
	Height     int                 `json:"height"`
	URL        string              `json:"url"`
	Filename   string              `json:"filename"`
	Size       int                 `json:"size"`
	Type       string              `json:"type"`
	Thumbnails map[string]Thumbnail `json:"thumbnails"`
}

// Fields contient les champs de chaque enregistrement
type Fields struct {
	MagicSeaweedLink        string   `json:"Magic Seaweed Link"`
	Photos                  []Photo  `json:"Photos"`
	Geocode                 string   `json:"Geocode"`
	Influencers             []string `json:"Influencers"`
	SurfBreak               []string `json:"Surf Break"`
	PeakSurfSeasonBegins    string   `json:"Peak Surf Season Begins"`
	DestinationStateCountry string   `json:"Destination State/Country"`
	PeakSurfSeasonEnds      string   `json:"Peak Surf Season Ends"`
	DifficultyLevel         int      `json:"Difficulty Level"`
	Destination             string   `json:"Destination"`
	Address                 string   `json:"Address"`
}

// Event représente chaque enregistrement avec son ID, l'heure de création et les champs
type Event struct {
	ID          string `json:"id"`
	CreatedTime string `json:"createdTime"`
	Fields      Fields `json:"fields"`
}

type AllEvents struct {
	Records []Event `json:"records"`
}

var events = AllEvents{
	Records: []Event{
		{
			ID:          "recTY7UmjDJl0VCvI",
			CreatedTime: "2018-05-31T00:16:16.000Z",
			Fields: Fields{
				MagicSeaweedLink: "https://magicseaweed.com/Pipeline-Backdoor-Surf-Report/616/",
				Photos: []Photo{
					{
						ID:       "attf6yu03NAtCuv5L",
						Width:    2233,
						Height:   1536,
						URL:      "https://v5.airtableusercontent.com/v3/u/32/32/1726056000000/wX1jRSJPeSuql1StB4zTLQ/_QEAcliZ1VmxvL4_kGbu4hfHEJhJu3fZBUi7OQ2FHTmURbLfJQmWqP33kuZU9J6CoflJZ7QY-1KVDPZ-MrYbQoNvmgjCwwUw_bFl_iJnsv8crJy0nXFIPa0gSDE57paU/u78tYA5D0MsHiYkreazUHtNnM6iH4jY6CtCIsEDNYPg",
						Filename: "thomas-ashlock-64485-unsplash.jpg",
						Size:     688397,
						Type:     "image/jpeg",
						Thumbnails: map[string]Thumbnail{
							"small": {
								URL:    "https://v5.airtableusercontent.com/v3/u/32/32/1726056000000/6Mvid7iBhitC3mnPR55T9A/Nl4M9mOBeEDP4wnnyrgTag7PtaIOfX14yMVSlhxA8DipkttMZs4qVWrBkQIlUOShovDypE6fnZ85xi856HOP_g8bBjhMWQrG0EyhoKFDZDjc5p_EqsC-cdTcYBqFCQVV-cXfGSoo0LulN1UB6Rrcmg/7McZvlKXCa_l2tOh9igLIk9-KXKrtocgbk8JyZPHPCM",
								Width:  52,
								Height: 36,
							},
							"large": {
								URL:    "https://v5.airtableusercontent.com/v3/u/32/32/1726056000000/PWd4HHrlDXs02WZJnUJrvg/3jB0G7KwFtdQPdDTlrmToouqWP8LXQXWy9TM4OpiKuU2fggGgLMq9dQi38wX4jk82GaZoqDKI68HxJ3X11oPWrqinlYzyk1TOsBKpUqVCfbmTmXrke5in6ut3CTcPGsqoIb98w9cbumg1lBFCKfexw/EQowCeHCBaHU4cu6VqTtcniQc7kwmm4_DdcbV8au_Po",
								Width:  744,
								Height: 512,
							},
							"full": {
								URL:    "https://v5.airtableusercontent.com/v3/u/32/32/1726056000000/RiaM5hlHlVxkAqid1qJ-WQ/BstFxa0p9hrC9Jt78XOSQOKPRq5Boj15CjMKnAT9nq4vuNK36XB5qCsulI0VTVscYn_3hQmKhNh_9ViTZx7goW9ZpXpwhZFt9_k53IWR1KJdGA3qwukWSdDMbAddfS2H/usms1qTLNWpRj4leEQjxbEQkS_-APLyVnfp1BrMHobo",
								Width:  2233,
								Height: 1536,
							},
						},
					},
				},
				Geocode:                 "eyJpIjoiUGlwZWxpbmUsIE9haHUsIEhhd2FpaSIsIm8iOnsic3RhdHVzIjoiT0siLCJmb3JtYXR0ZWRBZGRyZXNzIjoiRWh1a2FpIEJlYWNoIFBhcmssIEhhbGVpd2EsIEhJIDk2NzEyLCBVbml0ZWQgU3RhdGVzIiwibGF0IjoyMS42NjUwNTYyLCJsbmciOi0xNTguMDUxMjA0Njk5OTk5OTd9LCJlIjoxNTM1MzA3MDE5OTE1fQ==",
				Influencers:             []string{"recrP1aupHoWQxMe0", "recPdVWkPoHCQawnl"},
				SurfBreak:               []string{"Reef Break"},
				PeakSurfSeasonBegins:    "2024-07-22",
				DestinationStateCountry: "Oahu, Hawaii",
				PeakSurfSeasonEnds:      "2024-08-31",
				DifficultyLevel:         4,
				Destination:             "Pipeline",
				Address:                 "Pipeline, Oahu, Hawaii",
			},
		},
	},
}


//EVENTS
func createEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}