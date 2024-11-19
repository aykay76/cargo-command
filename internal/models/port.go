package models

type Port struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Capacity  int     `json:"capacity"`
	Location  string  `json:"location"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
