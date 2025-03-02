package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type HackathonPost struct {
	Id            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name          string             `json:"name" bson:"name"`
	Date          string             `json:"date" bson:"date"`
	Time          string             `json:"time" bson:"time"`
	Location      string             `json:"location" bson:"location"`
	Description   string             `json:"description" bson:"description"`
	Theme         string             `json:"theme" bson:"theme"`
	TeamSizeLimit int                `json:"team_size_limit" bson:"team_size_limit"`
	MediaUrls     []string           `json:"media_url" bson:"media_url"`
	CreatedAt     primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt     primitive.DateTime `json:"updated_at" bson:"updated_at"`
}
