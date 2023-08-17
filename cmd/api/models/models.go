package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	BookName    string             `json:"movie_name,omitempty" bson:"movie_name,omitempty"`
	Author      string             `json:"author,omitempty" bson:"author,omitempty"`
	ReleaseDate int                `json:"release_date,omitempty" bson:"release_date,omitempty"`
	Genre       string             `json:"genre,omitempty" bson:"genre,omitempty"`
}
