package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SingAvi/Go-Backend/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct{
	session *mgo.Session
}

func NewUserController(session *mgo.Session) *UserController{
	return &UserController{session}
}



func (uc UserController) GetAllUser (writer http.ResponseWriter, request *http.Request, params httprouter.Params){

	id := params.ByName("id")

	if !bson.IsObjectIdHex(id){
		writer.WriteHeader(http.StatusNotFound)
	}


	user := models.User{}

	if err := uc.session.DB("mongo-golang").C("users"); err!=nil{
		writer.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(user)

	if err!= nil{
		fmt.Println(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	fmt.Fprint(writer,"%s\n",uj)

}

func (uc UserController) GetUser (writer http.ResponseWriter, request *http.Request, params httprouter.Params){

	id := params.ByName("id")

	if !bson.IsObjectIdHex(id){
		writer.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	user := models.User{}

	if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&user); err!=nil{
		writer.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(user)

	if err!= nil{
		fmt.Println(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	fmt.Fprint(writer,"%s\n",uj)

}


func (uc UserController) CreateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params){

	user := models.User{}

	json.NewDecoder(request.Body).Decode(&user)

	user.Id = bson.NewObjectId()

	uc.session.DB("mongo-golang").C("users").Insert(user)

	uj,err := json.Marshal(user)

	if err != nil{
		fmt.Println(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	fmt.Println("User aDDED",uj)

}

func (uc UserController) DeleteUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params){

	id := params.ByName("id")

	if !bson.IsObjectIdHex(id){
		writer.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("mongo-golang").C("users").RemoveId(oid); err!=nil{
		writer.WriteHeader(404)
	}

	writer.WriteHeader(http.StatusOK)
	fmt.Fprint(writer,"Deleted User",oid, "\n")

}

