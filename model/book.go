package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title        string             `bson:"title,omitempty" json:"title,omitempty"`
	Author       string             `bson:"author,omitempty" json:"author,omitempty"`
	PublishedYear int                `bson:"publishedYear,omitempty" json:"publishedYear,omitempty"`
	Genre        string             `bson:"genre,omitempty" json:"genre,omitempty"`
}