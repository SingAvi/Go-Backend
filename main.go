package main


import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"
	"github.com/SingAvi/Go-Backend/controllers"
)


func main() {

	router := httprouter.New()
	uc := controllers.NewUserController(getSession())

	router.GET("/user/:id", uc.GetUser)
	router.POST("/user",uc.CreateUser)
	r.DELETE("/user/:id",uc.DeleteUser)
	http.ListenAndServe("localhost:9000", router)

}

func getSession() *mgo.Session{

	session, error := mgo.Dial("mongodb://localhost:27107")

	if error != nil{
		fmt.println("Error, session not started")
	}
	return session
}