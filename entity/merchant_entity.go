package entity

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MerchantCreate struct {
	ID 				primitive.ObjectID	`json:"id" bson:"_id,omitempty"`
	Address			string				`json:"address" bson:"address,omitempty"`
	Email			string				`json:"email" bson:"email,omitempty" validate:"required"`
	Phone			string				`json:"phone" bson:"phone,omitempty"`
	Pic_Name		string				`json:"pic_name" bson:"pic_name,omitempty"`
	Pic_Phone		string				`json:"pic_phone" bson:"pic_phone,omitempty"`
	Pic_Email		string				`json:"pic_email" bson:"pic_email,omitempty"`
	Information		Information			`json:"information" bson:"information,omitempty"`
	Service_Option	[]string			`json:"service_option" bson:"service_option,omitempty"`
	Password		string				`json:"password" bson:"password,omitempty" validate:"required"`
	Is_Registered	bool				`json:"is_registered" bson:"is_registered,omitempty"`
	Created_At		time.Time			`json:"created_at" bson:"created_at,omitempty"`
	Updated_At		time.Time			`json:"updated_at" bson:"updated_at,omitempty"`
	Deleted_At		*time.Time			`json:"deleted_at" bson:"deleted_at,omitempty"`
}

type Information struct {
	Category		string				`json:"category" bson:"category,omitempty"`
	Payment_Option	[]string			`json:"payment_option" bson:"payment_option,omitempty"`
	Opening_Hours	[]Opening_Hours		`json:"opening_hours" bson:"opening_hours,omitempty"`
}

type Opening_Hours struct {
	Day		string		`json:"day" bson:"day,omitempty"`
	Clock	[]Clock		`json:"clock" bson:"clock,omitempty"`
}

type Clock struct {
	Open	string		`json:"open" bson:"open,omitempty"`
	Close	string		`json:"close" bson:"close,omitempty"`
}

type MerchantUpdate struct {
	Address			string				`json:"address" bson:"address,omitempty"`
	Email			string				`json:"email" bson:"email" validate:"required,omitempty"`
	Phone			string				`json:"phone" bson:"phone,omitempty"`
	Pic_Name		string				`json:"pic_name" bson:"pic_name,omitempty"`
	Pic_Phone		string				`json:"pic_phone" bson:"pic_phone,omitempty"`
	Pic_Email		string				`json:"pic_email" bson:"pic_email,omitempty"`
	Information		Information			`json:"information" bson:"information,omitempty"`
	Service_Option	[]string			`json:"service_option" bson:"service_option,omitempty"`
	Password		string				`json:"password" bson:"password" validate:"required,omitempty"`
	Is_Registered	bool				`json:"is_registered" bson:"is_registered,omitempty"`
	Updated_At		time.Time			`json:"updated_at" bson:"updated_at,omitempty"`
}
