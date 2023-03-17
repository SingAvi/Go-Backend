package controllers

import (
	"fmt"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type User struct{
	session *mgo.Session
}

func NewUserController(session *mgo.Session) *UserController{
	return &UserController{session}
}

func (uc UserController) GetUser( writer http.ResponseWriter, request *http.Request, params httpRouter.Params){

	id := params.ByName("id")

	if !bson.IsObjectIdHex(id){
		writer.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectHex(id)

	user := models.User{}

	uc.Session.DB("mongo-golang").C("users").FindId(oid)
}

