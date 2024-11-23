package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FirstName string             `json:"firstName" bson:"firstName"`
	LastName  string             `json:"lastName" bson:"lastName"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"` // Store hashed passwords
	Role      string             `json:"role" bson:"role"`         // e.g., "admin", "user"
	Phone     string             `json:"phone" bson:"phone"`
	Address   string             `json:"address" bson:"address"` // Optional user address
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	Active    bool               `json:"active" bson:"active"` // Is the user account active
}
