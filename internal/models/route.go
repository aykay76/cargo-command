package models

type Route struct {
	ID          int        `json:"id"`
	StartPort   string     `json:"start_port"`
	EndPort     string     `json:"end_port"`
	Distance    int        `json:"distance"` // Distance in nautical miles
	Description string     `json:"description,omitempty"`
	Points      []GeoPoint `json:"waypoints,omitempty"`
}
