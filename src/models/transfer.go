package model

import "time"

type Transfer struct {
	Id      string  `json:"id,omitempty" bson:"_id,omitempty"`
	Value   float32 `json:"value,omitempty" bson:"value,omitempty"`
	Contact struct {
		Fullname      string `json:"fullname,omitempty" bson:"fullname,omitempty"`
		AccountNumber int    `json:"accountNumber,omitempty" bson:"accountNumber,omitempty"`
	} `json:"contact,omitempty" bson:"contact,omitempty"`
	DateTime time.Time `json:"dateTime,omitempty" bson:"dateTime,omitempty"`
}
