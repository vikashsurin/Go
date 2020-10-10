package posts

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Post ...
type Post struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User  primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
	Title string             `json:"title,omitempty" bson:"title,omitempty"`
	Desc  string             `json:"desc,omitempty" bson:"desc,omitempty"`
	Date  time.Time          `json:"time,omitempty" bson:"time,omitempty"`
}
