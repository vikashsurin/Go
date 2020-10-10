package usr

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User ...
type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Password []byte             `json:"password" bson:"password"`
	Avatar   string             `json:"avatar" bson:"avatar"`
	Date     time.Time          `json:"date" bson:"date"`
}
