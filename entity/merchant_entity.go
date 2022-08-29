package entity

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MERCHANT CREATE
type MerchantCreate struct {
	ID 				primitive.ObjectID	`json:"id" bson:"_id,omitempty"`
	Address			string				`json:"address" bson:"address,omitempty" validate:"required"`
	Email			string				`json:"email" bson:"email,omitempty" validate:"required,email"`
	Phone			string				`json:"phone" bson:"phone,omitempty" validate:"required"`
	Pic_Name		string				`json:"pic_name" bson:"pic_name,omitempty" validate:"required"`
	Pic_Phone		string				`json:"pic_phone" bson:"pic_phone,omitempty" validate:"required"`
	Pic_Email		string				`json:"pic_email" bson:"pic_email,omitempty" validate:"required"`
	Information		Information			`json:"information" bson:"information,omitempty" validate:"required"`
	Service_Option	[]string			`json:"service_option" bson:"service_option,omitempty" validate:"required"`
	Password		string				`json:"password" bson:"password,omitempty" validate:"required,min=8"`
	Is_Registered	bool				`json:"is_registered" bson:"is_registered,omitempty"`
	Created_At		time.Time			`json:"created_at" bson:"created_at,omitempty"`
	Updated_At		time.Time			`json:"updated_at" bson:"updated_at,omitempty"`
	Deleted_At		*time.Time			`json:"deleted_at" bson:"deleted_at,omitempty"`
}

type Information struct {
	Category		string				`json:"category" bson:"category,omitempty" validate:"required"`
	Payment_Option	[]string			`json:"payment_option" bson:"payment_option,omitempty" validate:"required"`
	Opening_Hours	[]Opening_Hours		`json:"opening_hours" bson:"opening_hours,omitempty" validate:"required"`
}

type Opening_Hours struct {
	Day		string		`json:"day" bson:"day,omitempty" validate:"required"`
	Clock	[]Clock		`json:"clock" bson:"clock,omitempty" validate:"required"`
}

type Clock struct {
	Open	string		`json:"open" bson:"open,omitempty" validate:"required"`
	Close	string		`json:"close" bson:"close,omitempty" validate:"required"`
}


// MERCHANT UPDATE
type MerchantUpdate struct {
	Address			string				`json:"address" bson:"address,omitempty" validate:"required"`
	Email			string				`json:"email" bson:"email,omitempty" validate:"required,email"`
	Phone			string				`json:"phone" bson:"phone,omitempty" validate:"required"`
	Pic_Name		string				`json:"pic_name" bson:"pic_name,omitempty" validate:"required"`
	Pic_Phone		string				`json:"pic_phone" bson:"pic_phone,omitempty" validate:"required"`
	Pic_Email		string				`json:"pic_email" bson:"pic_email,omitempty" validate:"required"`
	Information		Information			`json:"information" bson:"information,omitempty" validate:"required"`
	Service_Option	[]string			`json:"service_option" bson:"service_option,omitempty" validate:"required"`
	Password		string				`json:"password" bson:"password,omitempty" validate:"required,min=8"`
	Is_Registered	bool				`json:"is_registered" bson:"is_registered,omitempty"`
	Updated_At		time.Time			`json:"updated_at" bson:"updated_at,omitempty"`
}
