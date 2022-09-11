package entity

import (
	"time"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthRegister struct {
	ID 				primitive.ObjectID 	`json:"id" bson:"_id,omitempty"`
	Username		string				`json:"username" bson:"username,omitempty" validate:"required"`
	Email			string				`json:"email" bson:"email,omitempty" validate:"required,email"`
	Password		string				`json:"password" bson:"password,omitempty" validate:"required,min=8"`
	Created_At		time.Time			`json:"created_at" bson:"created_at,omitempty"`
	Updated_At		time.Time			`json:"updated_at" bson:"updated_at,omitempty"`
	Deleted_At		*time.Time			`json:"deleted_at" bson:"deleted_at,omitempty"`
}

type AuthLogin struct {
	Email 		string `bson:"email,omitempty" validate:"required"`
	Password 	string `bson:"password,omitempty" validate:"required"`
}

type Admin struct {
	ID 				primitive.ObjectID 	`json:"id" bson:"_id,omitempty"`
	Username		string				`json:"username" bson:"username,omitempty" validate:"required"`
	Email			string				`json:"email" bson:"email,omitempty" validate:"required,email"`
	Password		string				`json:"password" bson:"password,omitempty" validate:"required,min=8"`
	Created_At		time.Time			`json:"created_at" bson:"created_at,omitempty"`
	Updated_At		time.Time			`json:"updated_at" bson:"updated_at,omitempty"`
	Deleted_At		*time.Time			`json:"deleted_at" bson:"deleted_at,omitempty"`
}

type ResponseLogin struct {
	Admin 	Admin 	`json:"admin"`
	Token  	string 	`json:"token"`
}

type Claims struct {
	Username string
	Email    string
	jwt.StandardClaims
}