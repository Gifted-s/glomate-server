package helpers

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"

)

type Collections struct {
	Files *mongo.Collection
	Courses *mongo.Collection
	Departments *mongo.Collection
}
var client = GetDBClient()
func ConnectDB() Collections {
	fmt.Print("MongDB oonnected")
	files := client.Database("glomate_db").Collection("files")
    courses := client.Database("glomate_db").Collection("courses")
	departments := client.Database("glomate_db").Collection("departments")
	
	var collections Collections
    
	collections.Files = files
	collections.Courses= courses
	collections.Departments =  departments
	return collections
}
