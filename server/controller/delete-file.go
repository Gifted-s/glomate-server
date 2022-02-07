package controller

import (
	"encoding/json"
	"fmt"
	//	"fmt"
	//"log"

	"glomate/server/helpers"
	"glomate/server/models"

	"go.mongodb.org/mongo-driver/bson"

	//"fmt"
	// "github.com/gorilla/websocket"

	// "io"
	"context"
	"net/http"

	"github.com/gorilla/mux"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

func HandleDeleteFile(w http.ResponseWriter, r *http.Request) {
	collections := helpers.ConnectDB()
	var params = mux.Vars(r)
	// string to primitive.ObjectID
	course := params["course"]
	id:= params["id"]
	filter := bson.M{"name": course}
	change := bson.M{
		"$pull": bson.M{
			"lectures": bson.M{"name":id},
		},
	}
	update_result, err := collections.Courses.UpdateOne(context.TODO(), filter, change)
		if err != nil {
			// If there is an error in hasing password
			err := models.ResponseStruct{Status: "failed", Body: map[string]error{"error": err}}
			json.NewEncoder(w).Encode(err)
			return
		}
  fmt.Print(update_result)
	resp := models.ResponseStruct{Status: "success", Body: update_result, Token: ""}
	json.NewEncoder(w).Encode(resp)
	helpers.SendMail("File deletion alert", "sunkanmiadewumi1@gmail.com", id + " from "+ course ,  "<h3>This file was deleted</h3><br/><p>A new file has been deleted from glomate database. </p> <a style='padding:10px;background-color:royalblue;border:none;border-radius:10px;color:white;font-size:20px;' type='btn' href='https://glomate.netlify.app'>Check file out<a/>")

}
