package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct{
	Id          bson.ObjectId	`json:"id"` 			
	Name		string			`json:"name"` 			
	Designation string			`json:"designation"` 
	Experience	int				`json:"experience"` 	
}