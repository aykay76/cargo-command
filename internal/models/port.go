package models

type Port struct {
	ID        int     `json:"id"`
	Type      string  `json:"type"`
	Name      string  `json:"name"`
	Berths    int     `json:"berths"`
	Capacity  int     `json:"capacity"`
	Location  string  `json:"location"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

var Ports = []Port{
	{
		ID:        1,
		Name:      "Portsmouth",
		Capacity:  100,
		Berths:    1,
		Location:  "United Kingdom",
		Latitude:  50.811400,
		Longitude: -1.092031,
	},
	{
		ID:        2,
		Name:      "Tianjin",
		Capacity:  200,
		Berths:    4,
		Location:  "China",
		Latitude:  38.982222,
		Longitude: 117.746194,
	},
}
