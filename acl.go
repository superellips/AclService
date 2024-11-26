package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Acl struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	GuestbookId primitive.ObjectID `json:"guestbookId" bson:"guestbookId"`
	UserId      primitive.ObjectID `json:"userId" bson:"userId"`
}

func GetAclById(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("aclId"))
	if err != nil {
		return
	}
	var db MongoDb
	result := db.read(&id)
	c.IndentedJSON(http.StatusOK, result)
}

func GetAclsByUserId(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("userId"))
	if err != nil {
		return
	}
	var db MongoDb
	results := db.readByUserId(&id)
	c.IndentedJSON(http.StatusOK, &results)
}

func GetAclsByGuestbookId(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("gbId"))
	if err != nil {
		return
	}
	var db MongoDb
	results := db.readByGuestbookId(&id)
	c.IndentedJSON(http.StatusOK, &results)
}

func PostAcl(c *gin.Context) {
	var newAcl Acl
	if err := c.BindJSON(&newAcl); err != nil {
		return
	}
	var db MongoDb
	db.create(&newAcl)
	c.IndentedJSON(http.StatusCreated, newAcl)
}

func PutAcl(c *gin.Context) {
	var updatedAcl Acl
	if err := c.BindJSON(&updatedAcl); err != nil {
		return
	}
	var db MongoDb
	db.update(&updatedAcl)
	c.IndentedJSON(http.StatusAccepted, updatedAcl)
}

func DeleteAcl(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("aclId"))
	if err != nil {
		return
	}
	var db MongoDb
	deletedAcl := db.delete(&id)
	c.IndentedJSON(http.StatusAccepted, deletedAcl)
}
