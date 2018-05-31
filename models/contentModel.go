package models

import "gopkg.in/mgo.v2/bson"

type ContentModel struct {
	ID            bson.ObjectId `bson:"_id" json:"id"`
	ApplicationId string        `bson:"app_id" json:"app_id"`
	Title         string        `bson:"title" json:"title"`
	Description   string        `bson:"description" json:"description"`
	CoverImage    string        `bson:"cover_image" json:"cover_image"`
	Url           string        `bson:"url" json:"url"`
	ContentType   string        `bson:"content_type" json:"content_type"`
}
