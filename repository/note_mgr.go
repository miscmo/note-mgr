package repository

import (
	"context"
	"errors"

	"github.com/miscmo/note-mgr/entity"
	"github.com/miscmo/note-mgr/resource"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var NoteMgr NoteMgrInterface

type NoteMgrInterface interface {
	InsertOne(note entity.Note) (string, error)
	Find() ([]entity.Note, error)
}

type NoteMongo struct {
}

func (this *NoteMongo) InsertOne(note entity.Note) (string, error) {
	coll := resource.GetNoteCollection()

	result, err := coll.InsertOne(context.Background(), note)
	if err != nil {
		return "", err
	}

	objID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("objID is not primitive.ObjectID")
	}

	return objID.String(), nil
}

func (this *NoteMongo) Find() ([]entity.Note, error) {

	coll := resource.GetNoteCollection()

	filter := bson.M{}

	cursor, err := coll.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var notes []entity.Note
	for cursor.Next(context.Background()) {
		var note entity.Note
		err := cursor.Decode(&note)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func GetNoteMgr() NoteMgrInterface {
	if NoteMgr == nil {
		panic("NoteMgr is nil")
	}

	return NoteMgr
}

func Init() {
	if NoteMgr == nil {
		NoteMgr = &NoteMongo{}
	}
	if ElementMgr == nil {
		ElementMgr = &ElementMongo{}
	}
}
