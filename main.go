package main

import (
	"fmt"
	"net/http"

	"github.com/SingAvi/Go-Backend/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)


func main() {

	router := httprouter.New()
	uc := controllers.NewUserController(getSession())

	router.GET("/user/:id", uc.GetUser)
	router.POST("/user/add",uc.CreateUser)
	router.DELETE("/user/:id",uc.DeleteUser)
	http.ListenAndServe("localhost:9000", router)

}

func getSession() *mgo.Session{

	session, error := mgo.Dial("mongodb://localhost:27017")

	if error != nil{
		panic(error)
	}
	fmt.Printf("Server Running")
	return session
}