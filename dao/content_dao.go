package dao

import (
	"log"

	. "github.com/abhi06276/contents-restapi/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ContentsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "content"
)

// Establish a connection to database
func (m *ContentsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of movies
func (m *ContentsDAO) FindAll() ([]ContentModel, error) {
	var movies []ContentModel
	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}

// Find a movie by its id
func (m *ContentsDAO) FindByAppId(id string) ([]ContentModel, error) {
	var contents []ContentModel
	err := db.C(COLLECTION).Find(bson.M{"app_id": id}).All(&contents)
	return contents, err
}

// Insert a movie into database
func (m *ContentsDAO) Insert(movie ContentModel) error {
	err := db.C(COLLECTION).Insert(&movie)
	return err
}

// Delete an existing movie
func (m *ContentsDAO) Delete(movie ContentModel) error {
	err := db.C(COLLECTION).Remove(&movie)
	return err
}

// Update an existing movie
func (m *ContentsDAO) Update(movie ContentModel) error {
	err := db.C(COLLECTION).UpdateId(movie.ID, &movie)
	return err
}
