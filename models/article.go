package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/teplovyuriydeveloper/article/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	ID        string             `bson:"_id,omitempty" json:"id,omitempty"`
	Title     string             `bson:"title" json:"title"`
	Content   string             `bson:"content" json:"content"`
	Author    string             `bson:"author" json:"author"`
	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
}

func (a *Article) Create() error {
	a.ID = uuid.NewString()
	a.CreatedAt = primitive.NewDateTimeFromTime(time.Now())

	collection := db.Client.Database("article_db").Collection("articles")

	_, err := collection.InsertOne(context.Background(), a)
	if err != nil {
		return err
	}

	return nil
}
