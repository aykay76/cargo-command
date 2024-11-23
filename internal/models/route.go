package models

type Route struct {
	ID          int        `json:"id"`
	Description string     `json:"description,omitempty"`
	StartPort   string     `json:"start_port"`
	EndPort     string     `json:"end_port"`
	Distance    int        `json:"distance"` // Distance in nautical miles
	Points      []GeoPoint `json:"waypoints,omitempty"`
}

var Routes = []Route{
	{
		ID:          1,
		Description: "Route from Portsmouth, UK to Tianjin, China",
		StartPort:   "Portsmouth",
		EndPort:     "Tianjin",
		Distance:    9000,
		Points: []GeoPoint{
			{Latitude: 50.811400, Longitude: -1.092031},
			{Latitude: 38.982222, Longitude: 117.746194},
		},
	},
}
