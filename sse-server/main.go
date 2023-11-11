package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type EventData struct {
	Timestamp           string   `json:"timestamp"`
	DecreaseConsumption string   `json:"decrease_consumption"`
	Forecast            []string `json:"forecast"`
}

func randomYesNo() string {
	if rand.Intn(2) == 0 {
		return "Y"
	}
	return "N"
}

func updateForecast(forecast []string, currentTime time.Time) []string {
	forecast = forecast[1:]
	newForecastTime := currentTime.Add(90 * time.Second) // Next 90 seconds
	forecast = append(forecast, fmt.Sprintf("%s %s", newForecastTime.Format(time.RFC1123), randomYesNo()))
	return forecast
}

func eventHandler(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	forecast := make([]string, 9)
	currentTime := time.Now().UTC()
	for i := range forecast {
		forecastTime := currentTime.Add(time.Duration(i+1) * 10 * time.Second)
		forecast[i] = fmt.Sprintf("%s %s", forecastTime.Format(time.RFC1123), randomYesNo())
	}

	for {
		select {
		case t := <-ticker.C:
			forecast = updateForecast(forecast, t.UTC())
			event := EventData{
				Timestamp:           t.UTC().Format(time.RFC1123),
				DecreaseConsumption: randomYesNo(),
				Forecast:            forecast,
			}
			jsonData, err := json.Marshal(event)
			if err != nil {
				http.Error(w, "Error generating JSON", http.StatusInternalServerError)
				return
			}

			fmt.Fprintf(w, "data: %s\n\n", jsonData)
			flusher.Flush()

		case <-r.Context().Done():
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/events", eventHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
