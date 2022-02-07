package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FileStruct struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name          string             `json:"name" bson:"name,omitempty"`
	Type          string             `json:"type" bson:"type,omitempty"`
	Date_Added    string             `json:"date_added" bson:"date_added,omitempty"`
	Size          float64            `json:"size" bson:"size,omitempty"`
	Download_Link string             `json:"download_link" bson:"download_link,omitempty"`
}
type Departments struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	School       string             `json:"school" bson:"school,omitempty"`
	Name         string             `json:"name" bson:"name,omitempty"`
	Last_Updated string             `json:"last_updated" bson:"last_updated,omitempty"`
}

type Course struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Dept         string             `json:"dept" bson:"dept,omitempty"`
	Level        string             `json:"level" bson:"level,omitempty"`
	School       string             `json:"school" bson:"school,omitempty"`
	Name         string             `json:"name" bson:"name,omitempty"`
	Lectures     []FileStruct       `json:"lectures" bson:"lectures,omitempty"`
	Last_Updated string             `json:"last_updated" bson:"last_updated,omitempty"`
}

type ResponseStruct struct {
	Status string      `json:"status"`
	Token  string      `json:"token"`
	Body   interface{} `json:"body"`
}
