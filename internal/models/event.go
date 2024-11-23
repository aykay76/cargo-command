package models

import "time"

type Event struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// the time the event starts and ends
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	// some events will have a sphere of influence
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Radius    float64 `json:"radius"`
	// the event will come from a source
	Source string `json:"source"`
	// the event will affect objects of a certain type
	TargetType string `json:"target_type"`
	// the event will have a certain effect
	Effect string `json:"effect"`
	// the event will have a certain magnitude
	Magnitude float64 `json:"magnitude"`
}
