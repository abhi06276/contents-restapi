package dao

import (
	"log"
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
func (m *MoviesDAO) FindById(id string) (ContentModel, error) {
	var movie ContentModel
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

// Insert a movie into database
func (m *MoviesDAO) Insert(movie ContentModel) error {
	err := db.C(COLLECTION).Insert(&movie)
	return err
}

// Delete an existing movie
func (m *MoviesDAO) Delete(movie ContentModel) error {
	err := db.C(COLLECTION).Remove(&movie)
	return err
}

// Update an existing movie
func (m *MoviesDAO) Update(movie ContentModel) error {
	err := db.C(COLLECTION).UpdateId(movie.ID, &movie)
	return err
}
