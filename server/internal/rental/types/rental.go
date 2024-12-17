package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RentalType string

const (
	Shared      RentalType = "shared"
	Independent RentalType = "independent"
	Sale        RentalType = "sale"
)

type Standing string

const (
	Economy  Standing = "economy"
	Standard Standing = "standard"
	Luxury   Standing = "luxury"
)

type Status string

const (
	Agreed   Status = "agreed"
	Declined Status = "declined"
	Pending  Status = "pending"
)

type Address struct {
	StreetNumber string `json:"streetNumber" bson:"streetNumber" validate:"required"`
	Street       string `json:"street" bson:"street" validate:"required"`
	City         string `json:"city" bson:"city" validate:"required"`
	Country      string `json:"country" bson:"country" validate:"required"`
	FullAddress  string `json:"fullAddress" bson:"fullAddress""`
}

type Geometry struct {
	Lat string `json:"lat" bson:"lat" validate:"required,latitude,min=-30,max=37.5"`
	Lng string `json:"lng" bson:"lng" validate:"required,longitude,min=7,max=11.5"`
}

type Amenities struct {
	AirConditioning bool `json:"airConditioning" bson:"airConditioning" validate:"required"`
	Heating         bool `json:"heating" bson:"heating" validate:"required"`
	Refrigerator    bool `json:"refrigerator" bson:"refrigerator"`
	Parking         bool `json:"parking" bson:"parking"`
}

type Rules struct {
	PetsAllowed    bool `json:"petsAllowed" bson:"petsAllowed"`
	PartiesAllowed bool `json:"partiesAllowed" bson:"partiesAllowed"`
	SmokingAllowed bool `json:"smokingAllowed" bson:"smokingAllowed"`
}

type Rental struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name          string             `json:"name" bson:"name" validate:"required,min=3,max=100"`
	Address       Address            `json:"address" bson:"address"`
	Geometry      Geometry           `json:"geometry" bson:"geometry"`
	Images        []string           `json:"images" bson:"images" validate:"dive,url,max=10"` // URLs or file paths for uploaded images
	AgreeToTerms  bool               `json:"agreeToTerms" bson:"agreeToTerms" validate:"required"`
	Status        Status             `json:"status" bson:"status" validate:"required,oneof=agreed declined pending" default:"pending"`
	Description   string             `json:"description" bson:"description" validate:"required,max=500"`
	Price         int64              `json:"price" bson:"price" validate:"required,min=0"`
	Currency      string             `json:"currency" bson:"currency" validate:"required,oneof=TND USD EUR" default:"TND"`
	Bedrooms      int64              `json:"bedrooms" bson:"bedrooms" validate:"required,min=0"`
	Bathrooms     int64              `json:"bathrooms" bson:"bathrooms" validate:"required,min=0"`
	AreaSize      int64              `json:"areaSize" bson:"areaSize" validate:"required,min=0"`
	Available     bool               `json:"available" bson:"available" default:"true"`
	AvailableFrom time.Time          `json:"availableFrom" bson:"availableFrom" validate:"required"`
	Tags          []string           `json:"tags" bson:"tags" validate:"dive,min=1,max=50"`
	Type          RentalType         `json:"type" bson:"type" validate:"required,oneof=shared independent sale"`
	Standing      Standing           `json:"standing" bson:"standing" validate:"required,oneof=economy standard luxury" default:"standard"`
	Amenities     Amenities          `json:"amenities" bson:"amenities"`
	Rules         Rules              `json:"rules" bson:"rules"`
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt" validate:"required"`
	UpdatedAt     time.Time          `json:"updatedAt" bson:"updatedAt" validate:"required"`
	CreatedBy     primitive.ObjectID `json:"createdBy" bson:"createdBy" validate:"required"`         // Reference to User ID
	UpdatedBy     primitive.ObjectID `json:"updatedBy" bson:"updatedBy" validate:"required"`         // Reference to User ID
	DeletedAt     *time.Time         `json:"deletedAt" bson:"deletedAt"`                             // Soft delete field
	LastUpdatedBy primitive.ObjectID `json:"lastUpdatedBy" bson:"lastUpdatedBy" validate:"required"` // Audit logging
}
