package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Title string             `bson:"title" json:"title"`
	Desc  string             `bson:"description" json:"description"`
	Done  bool               `bson:"done" json:"done"`
}

func (col Todo) GetCollName() string {
	return "todos"
}
