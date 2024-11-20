package main

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDb struct {
	Uri string
}

func (p *MongoDb) create(a *Acl) *Acl {
	coll := p.aclCollection()
	result, err := coll.InsertOne(
		context.TODO(),
		a)
	if err != nil {
		panic(err)
	}
	a.Id = result.InsertedID.(primitive.ObjectID)
	return a
}

func (p *MongoDb) read(i *primitive.ObjectID) *Acl {
	coll := p.aclCollection()
	filter := bson.D{{Key: "_id", Value: i}}
	var result Acl
	err := coll.FindOne(
		context.TODO(),
		filter).Decode(&result)
	if err != nil {
		panic(err)
	}
	return &result
}

func (p *MongoDb) readByUserId(i *primitive.ObjectID) *[]Acl {
	coll := p.aclCollection()
	filter := bson.D{{Key: "userId", Value: i}}
	var results []Acl
	cur, err := coll.Find(
		context.TODO(),
		filter)
	if err != nil {
		panic(err)
	}
	cur.All(
		context.TODO(),
		&results)
	return &results
}

func (p *MongoDb) readByGuestbookId(i *primitive.ObjectID) *[]Acl {
	coll := p.aclCollection()
	filter := bson.D{{Key: "guestbookId", Value: i}}
	var results []Acl
	cur, err := coll.Find(
		context.TODO(),
		filter)
	if err != nil {
		panic(err)
	}
	cur.All(
		context.TODO(),
		&results)
	return &results
}

func (p *MongoDb) update(a *Acl) *Acl {
	coll := p.aclCollection()
	filter := bson.D{{Key: "_id", Value: a.Id}}
	var result Acl
	err := coll.FindOneAndReplace(
		context.TODO(),
		filter,
		a).Decode(&result)
	if err != nil {
		panic(err)
	}
	return &result
}

func (p *MongoDb) delete(i *primitive.ObjectID) *Acl {
	coll := p.aclCollection()
	filter := bson.D{{Key: "_id", Value: i}}
	var deletedAcl Acl
	err := coll.FindOneAndDelete(
		context.TODO(),
		filter).Decode(&deletedAcl)
	if err != nil {
		panic(err)
	}
	return &deletedAcl
}

func (p *MongoDb) aclCollection() *mongo.Collection {
	p.Uri = os.Getenv("GB_CONSTRING")
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(p.Uri))
	if err != nil {
		panic(err)
	}
	return client.Database("acl-service").Collection("acls")
}
