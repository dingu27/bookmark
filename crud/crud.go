package crud

import (
	"github.com/dingu27/bookmark/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Mongo ...
type Mongo struct {
	db         *mgo.Database
	collection string
}

//NewMongo ...
func NewMongo(db *mgo.Database, collection string) *Mongo {
	return &Mongo{
		db:         db,
		collection: collection,
	}

}

//Save ...
func (r *Mongo) Save(readmark *model.Readmark) error {
	var bookmark model.Bookmark
	bookmark.Email = readmark.Email
	bookmark.URL = readmark.URL
	err := r.db.C(r.collection).Insert(bookmark)
	return err
}

//Delete ...
func (r *Mongo) Delete(bookmark *model.Bookmark) error {
	err := r.db.C(r.collection).Remove(bson.M{"email": bookmark.Email, "url": bookmark.URL})
	return err
}

//FindByEmail ...
func (r *Mongo) FindByEmail(email string) (*model.GetBookmarks, error) {

	var bookmarks model.Bookmarks
	err := r.db.C(r.collection).Find(bson.M{"email": email}).All(&bookmarks)
	if err != nil {
		return nil, err
	}
	var getmarks model.GetBookmarks
	for _, item := range bookmarks {
		getmarks.URLs = append(getmarks.URLs, item.URL)
	}
	return &getmarks, nil
}

//FindByEmailURL ...
func (r *Mongo) FindByEmailURL(readmark *model.Readmark) (*model.Bookmark, error) {

	var bookmark model.Bookmark
	err := r.db.C(r.collection).Find(bson.M{"email": readmark.Email, "url": readmark.URL}).One(&bookmark)

	if err != nil {
		return nil, err
	}
	return &bookmark, nil
}
