package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	BookName    string             `json:"book_name,omitempty" bson:"book_name,omitempty"`
	Author      string             `json:"author,omitempty" bson:"author,omitempty"`
	ReleaseDate int                `json:"release_date,omitempty" bson:"release_date,omitempty"`
	Genre       string             `json:"genre,omitempty" bson:"genre,omitempty"`
	StockCount  int                `json:"stock_count,omitempty" bson:"stock_count,omitempty"`
	Review      []Review           `json:"reviews,omitempty" bson:"reviews,omitempty"`
}

type Review struct {
	Rating int    `json:"rating,omitempty" bson:"rating,omitempty"`
	Review string `json:"review,omitempty" bson:"review,omitempty"`
}
