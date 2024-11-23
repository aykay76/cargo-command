package models

type VesselFile struct {
	Type string     `json:"type"`
	Data VesselData `json:"data"`
}

type VesselData struct {
	Vessels   []Vessel `json:"rows"`
	AreaShips int      `json:"areaShips"`
}

type Vessel struct {
	ID          string  `json:"SHIP_ID"`     // unique identifier
	Type        string  `json:"TYPE"`        // type of the vessel
	Name        string  `json:"SHIPNAME"`    // name of the vessel
	Width       float64 `json:"WIDTH"`       // width of the vessel
	Latitude    float64 `json:"LAT"`         // current latitude (degrees)
	Longitude   float64 `json:"LON"`         // current longitude (degrees)
	SpeedKnots  float64 `json:"SPEED"`       // current speed in knots
	Course      int     `json:"COURSE"`      // current course in degrees
	Heading     int     `json:"HEADING"`     // current heading
	Elapsed     int     `json:"ELAPSED"`     // elapsed time
	Destination string  `json:"DESTINATION"` // destination
	Flag        string  `json:"FLAG"`        // flag code
	Length      int     `json:"LENGTH"`      // length of the vessel
	ShipType    int     `json:"SHIPTYPE"`    // type of the vessel
}

var Vessels = []Vessel{
	{
		ID:          "IMO100000001",
		Name:        "Vessel 1",
		Width:       10.0,
		Latitude:    50.811400,
		Longitude:   -1.092031,
		SpeedKnots:  10.0,
		Course:      90,
		Heading:     90,
		Elapsed:     0,
		Destination: "Portsmouth",
		Flag:        "GB",
		Length:      100,
		ShipType:    1,
	},
}
