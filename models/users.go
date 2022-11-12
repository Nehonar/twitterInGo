package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
invalidLengthUserEmail invalid number of characters to pass user email
invalidLengthUserPass invalid number of characters to pass user password
*/
var InvalidLengthUserEmail = 0
var InvalidLengthUserPass = 6

/*
User user model to MongoDB
*/
type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name,omitempty"`
	Lastname    string             `bson:"lastname" json:"lastname,omitempty"`
	DateOfBirth time.Time          `bson:"dateofbirth" json:"dateofbirth,omitempty"`
	Email       string             `bson:"email" json:"email,omitempty"`
	Password    string             `bson:"password" json:"password,omitempty"`
	Avatar      string             `bson:"avatar" json:"avatar,omitempty"`
	Banner      string             `bson:"banner" json:"banner,omitempty"`
	Biography   string             `bson:"biography" json:"biography,omitempty"`
	Location    string             `bson:"location" json:"location,omitempty"`
	WebSite     string             `bson:"webSite" json:"webSite,omitempty"`
}
