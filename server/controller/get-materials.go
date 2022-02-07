package controller

import (
	"encoding/json"
	"fmt"
	//"fmt"
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
func HandleGetCourseMaterials(w http.ResponseWriter, r *http.Request) {
	collections := helpers.ConnectDB()
	var course models.Course
	var params = mux.Vars(r)
	fmt.Print(params)
	// string to primitive.ObjectID
	dept := params["dept"]
	level := params["lev"]
	course_name := params["course"]
	filter := bson.M{"dept": dept, "level": level,"name":course_name }
	err := collections.Courses.FindOne(context.TODO(), filter).Decode(&course)
	if err != nil {
		resp := models.ResponseStruct{Status: "failed", Body: map[string]string{"msg": "error occured while fetchinng courses"}}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := models.ResponseStruct{Status: "success", Body: course.Lectures, Token: ""}
	json.NewEncoder(w).Encode(resp)

	// save loan request
	//helpers.SendMail("Helpful Platform","sunkanmiadewumi1@gmail.com","New Loan Request","<h1>New loan request</h1><br/><p>We are always willing to help</p> <a type='btn' href='https://hello.com'>View Rewuest<a/>" )
}
