package entity

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Merchant struct {
	ID 				primitive.ObjectID	`json:"id" bson:"id"`
	Logo			string				`json:"logo" bson:"logo"`
	Address			string				`json:"address" bson:"address"`
	Email			string				`json:"email" bson:"email" validate:"required"`
	Phone			string				`json:"phone" bson:"phone"`
	Pic_Name		string				`json:"pic_name" bson:"pic_name"`
	Pic_Phone		string				`json:"pic_phone" bson:"pic_phone"`
	Pic_Email		string				`json:"pic_email" bson:"pic_email"`
	Information		Information			`json:"information" bson:"information"`
	Service_Option	[]string			`json:"service_option" bson:"service_option"`
	Password		string				`json:"password" bson:"password" validate:"required"`
	Is_Registered	bool				`jsob:"is_registered" bson:"is_registered"`
	Created_At		*time.Time			`json:"created_at" bson:"created_at"`
	Updated_At		*time.Time			`json:"updated_at" bson:"updated_at"`
	Deleted_At		*time.Time			`json:"deleted_at" bson:"deleted_at"`
}

type Information struct {
	Category		string				`json:"category" bson:"category"`
	Payment_Option	[]string			`json:"payment_option" bson:"payment_option"`
	Opening_Hours	[]Opening_Hours		`json:"opening_hours" bson:"opening_hours"`
}

type Opening_Hours struct {
	Day		string		`json:"day" bson:"day"`
	Clock	[]Clock		`json:"clock" bson:"clock"`
}

type Clock struct {
	Open	string		`json:"open" bson:"open"`
	Close	string		`json:"close" bson:"close"`
}
