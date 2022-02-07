package controller

import (
	"encoding/json"
	"fmt"
	//"fmt"

	"log"

	"glomate/server/helpers"
	"glomate/server/models"

	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"

	//"fmt"
	// "github.com/gorilla/websocket"

	// "io"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleGetCourses(w http.ResponseWriter, r *http.Request) {
	collections := helpers.ConnectDB()
	var courses []models.Course
	var params = mux.Vars(r)
	fmt.Print(params)
	// string to primitive.ObjectID
	dept := params["dept"]
	level := params["lev"]
	filter := bson.M{"dept": dept, "level": level}
	cur, err := collections.Courses.Find(context.TODO(), filter)
	if err != nil {
		resp := models.ResponseStruct{Status: "failed", Body: map[string]string{"msg": "error occured while fetchinng courses"}}
		json.NewEncoder(w).Encode(resp)
		return
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var course models.Course
		// & character returns the memory address of the following variable.
		err := cur.Decode(&course) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}
		// add item our array
		courses = append(courses, course)
	}
	resp := models.ResponseStruct{Status: "success", Body: courses, Token: ""}
	json.NewEncoder(w).Encode(resp)

	// save loan request
	//helpers.SendMail("Helpful Platform","sunkanmiadewumi1@gmail.com","New Loan Request","<h1>New loan request</h1><br/><p>We are always willing to help</p> <a type='btn' href='https://hello.com'>View Rewuest<a/>" )
}
