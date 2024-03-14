package entity

type Satellite struct {
	Name     string   `json:"name"`
	Distance float64  `json:"distance"`
	Message  []string `json:"message"`
}

type SatelliteRequest struct {
	Satellites []Satellite `json:"satellites"`
}

type SatelliteData struct {
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}

type Satellitequery struct {
	Satellite string
}
