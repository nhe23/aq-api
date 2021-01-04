package model

type LocationResult struct {
	ID           string         `json:"_id" bson:"_id"`
	Location     string         `json:"location" bson:"location"`
	City         string         `json:"city" bson:"city"`
	Country      string         `json:"countryObj" bson:"country"`
	Measurements []*Measurement `json:"measurements" bson:"measurements"`
	Coordinates  *Coordinates   `json:"coordinates" bson:"coordinates"`
}
