package model

type Transfer struct {
	Id      string  `json:"id,omitempty" bson:"_id,omitempty"`
	Value   float32 `json:"value,omitempty" bson:"value,omitempty"`
	Contact struct {
		Name          string `json:"name,omitempty" bson:"name,omitempty"`
		AccountNumber int    `json:"accountNumber,omitempty" bson:"accountNumber,omitempty"`
	} `json:"contact,omitempty" bson:"contact,omitempty"`
	DateTime string `json:"dateTime,omitempty" bson:"dateTime,omitempty"`
}
