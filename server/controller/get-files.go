package controller

import (
	"encoding/json"
	"log"

	"glomate/server/helpers"
	"glomate/server/models"

	"go.mongodb.org/mongo-driver/bson"

	//"fmt"
	// "github.com/gorilla/websocket"

	// "io"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleGetFiles(w http.ResponseWriter, r *http.Request) {
	collections := helpers.ConnectDB()
	var files []models.FileStruct
	var params = mux.Vars(r)
	log.Println(params)
	// string to primitive.ObjectID

	filter := bson.M{}

	cur, err := collections.Files.Find(context.TODO(), filter)

	if err != nil {
		resp := models.ResponseStruct{Status: "failed", Body: map[string]string{"msg": "error occured while using the software"}}
		json.NewEncoder(w).Encode(resp)
		return
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var file models.FileStruct
		// & character returns the memory address of the following variable.
		err := cur.Decode(&file) // decode similar to deserialize process.
		log.Println(file)
		if err != nil {
			log.Fatal(err)
		}

		// add item our array

		files = append(files, file)
	}

	resp := models.ResponseStruct{Status: "success", Body: files, Token: ""}
	json.NewEncoder(w).Encode(resp)

	// save loan request
	//helpers.SendMail("Helpful Platform","sunkanmiadewumi1@gmail.com","New Loan Request","<h1>New loan request</h1><br/><p>We are always willing to help</p> <a type='btn' href='https://hello.com'>View Rewuest<a/>" )
}
