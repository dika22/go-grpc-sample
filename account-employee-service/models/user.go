package models

import "time"

type User struct {
	ID 			string 	`json:"_id" bson:"_id,omitempty"`
	Nama 		string 	`json:"kodeCabang" bson:"kodeCabang"`
	Alamat 		string 	`json:"namaCabang" bson:"namaCabang"`
	KodePos 	int64 	`json:"kodePos" bson:"kodePos"`
	Provinsi 	string 	`json:"provinsi" bson:"provinsi"`
	Kantor 		string 	`json:"kantor" bson:"kantor"`
	CreatedAt 	time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt 	time.Time `json:"updatedAt" bson:"updatedAt"`
}

type UpdateUser struct {
	Nama 		string 	`json:"kodeCabang" bson:"kodeCabang"`
	Alamat 		string 	`json:"namaCabang" bson:"namaCabang"`
	KodePos 	int64 	`json:"kodePos" bson:"kodePos"`
	Provinsi 	string 	`json:"provinsi" bson:"provinsi"`
	Kantor 		string 	`json:"kantor" bson:"kantor"`
	UpdatedAt 	time.Time `json:"updatedAt" bson:"updatedAt"`
}