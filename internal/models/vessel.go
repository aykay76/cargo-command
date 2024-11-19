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
	ID          string  `json:"SHIP_ID"`
	Name        string  `json:"SHIPNAME"`
	Width       float64 `json:"WIDTH"`
	Latitude    float64 `json:"LAT"` // current latitude (degrees)
	Longitude   float64 `json:"LON"`
	SpeedKnots  float64 `json:"SPEED"`       // current speed in knots
	Course      int     `json:"COURSE"`      // current course in degrees
	Heading     int     `json:"HEADING"`     // current heading
	Elapsed     int     `json:"ELAPSED"`     // elapsed time
	Destination string  `json:"DESTINATION"` // destination
	Flag        string  `json:"FLAG"`        // flag code
	Length      int     `json:"LENGTH"`      // length of the vessel
	ShipType    int     `json:"SHIPTYPE"`    // type of the vessel
}
