package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Rental struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name"`
	StreetNumber string             `json:"streetNumber" bson:"streetNumber"`
	Street       string             `json:"street" bson:"street"`
	City         string             `json:"city" bson:"city"`
	Country      string             `json:"country" bson:"country"`
	FullAddress  string             `json:"fullAddress" bson:"fullAddress"`
	Lat          string             `json:"lat" bson:"lat"`
	Lng          string             `json:"lng" bson:"lng"`
	Images       []string           `json:"images" bson:"images"` // URLs or file paths for uploaded images
	Agree        bool               `json:"agree" bson:"agree"`
	Status       bool               `json:"status" bson:"status"`
}
