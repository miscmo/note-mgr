package repository

import (
	"context"

	"github.com/miscmo/note-mgr/entity"
	"github.com/miscmo/note-mgr/resource"
	"go.mongodb.org/mongo-driver/bson"
)

var ElementMgr ElementMgrInterface

type ElementMgrInterface interface {
	GetSubElems(string) ([]entity.Element, error)
	GetElem(string) (entity.Element, error)
	AddElem(entity.Element) (string, error)

	// move elem to other root elem
	Move(string, string) error
	// Attach(string, string) error
}

type ElementMongo struct{}

func (this *ElementMongo) FindOne() {}

func (this *ElementMongo) GetElem(id string) (entity.Element, error) {

	var (
		result entity.Element
		err    error
	)

	filter := bson.M{"_id": id}

	err = resource.GetElementCollection().FindOne(context.Background(), filter).Decode(&result)

	return result, err
}

func (this *ElementMongo) GetSubElems(id string) ([]entity.Element, error) {
	var (
		result []entity.Element
		root   entity.Element
		err    error
	)

	root, err = this.GetElem(id)
	if err != nil {
		return result, err
	}

	var subElems []entity.Element

	for _, e := range root.SubElems {
		elem, err := this.GetElem(e)
		if err != nil {
			panic(err)
		}

		subElems = append(subElems, elem)
	}

	return result, nil
}

func (this *ElementMongo) AddElem(elem entity.Element) (string, error) {
	return "", nil
}

func (this *ElementMongo) Move(id string, parentID string) error {
	return nil
}

func GetElementMgr() ElementMgrInterface {
	return ElementMgr
}
