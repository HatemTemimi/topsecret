package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	Description  string             `json:"description" bson:"description"`
	Price        int64              `json:"price" bson:"price"`
	Bedrooms     int64              `json:"bedrooms" bson:"bedrooms"`
	Bathrooms    int64              `json:"bathrooms" bson:"bathrooms"`
	AreaSize     int64              `json:"areaSize" bson:"areaSize"`
	Available    bool               `json:"available" bson:"available"`
	Tags         []string           `json:"tags" bson:"tags"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt" bson:"updatedAt"`
	CreatedBy    primitive.ObjectID `json:"createdBy" bson:"createdBy"` // Reference to User ID
	UpdatedBy    primitive.ObjectID `json:"updatedBy" bson:"updatedBy"` // Reference to User ID
}
