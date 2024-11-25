package models

type Player struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Money   int      `json:"money"`
	Vessels []Vessel `json:"vessels"`
}

func (p *Player) AddVessel(v Vessel) {
	p.Vessels = append(p.Vessels, v)
}

func (p *Player) RemoveVessel(v Vessel) {
	for i, vessel := range p.Vessels {
		if vessel.ID == v.ID {
			p.Vessels = append(p.Vessels[:i], p.Vessels[i+1:]...)
			break
		}
	}
}

func (p *Player) GetVesselByID(id int) *Vessel {
	for _, vessel := range p.Vessels {
		if vessel.ID == id {
			return &vessel
		}
	}
	return nil
}
