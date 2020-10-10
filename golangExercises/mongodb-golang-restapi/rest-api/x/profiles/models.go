package profiles

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Profile ...
type Profile struct {
	User    primitive.ObjectID `json:"_id,omitempty" bson:"_id, omitempty"`
	Role    string             `json:"role,omitempty" bson:"role,omitempty"`
	Course  string             `json:"course,omitempty" bson:"course,omitempty"`
	Hobbies []string           `Json:"hobbies,omitempty" bson:"hobbies,omitempty"`
	Date    time.Time          `json:"date,omitempty" bson:"date,omitempty"`
}
