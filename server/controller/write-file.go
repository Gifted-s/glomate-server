package controller

import (
	"encoding/json"
	"glomate/server/helpers"
	"glomate/server/models"
	//"log"

	//"fmt"
	// "github.com/gorilla/websocket"
	// "io"
	"context"
	"net/http"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

// HandleCreateBlog this will help us create blog items
func HandleStoreFile(w http.ResponseWriter, r *http.Request) {
	collections := helpers.ConnectDB()
	var file models.FileStruct
	err := json.NewDecoder(r.Body).Decode(&file)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	if file.Type == "text/UTF-8" {
		file.Size = float64(len([]byte(file.Download_Link))) / 1000 // get link size and convert to kilobyte
	}
	// string to primitive.ObjectID

	result, err := collections.Files.InsertOne(context.TODO(), file)
	if err != nil {
		// If there is an error in hasing password
		err := models.ResponseStruct{Status: "failed", Body: map[string]string{"error": "Counld Not Insert  Request To DB please try again"}}
		json.NewEncoder(w).Encode(err)
		return
	}
	resp := models.ResponseStruct{Status: "success", Body: result, Token: ""}
	json.NewEncoder(w).Encode(resp)
	if file.Type == "text-UTF-8" {
		helpers.SendMail("New Link Alert", "sunkanmiadewumi1@gmail.com", "New Link Alert", "<h3>A Link was added</h3><br/><p>A new link has been added to glomate database, rememeber to share this across the devices that is needed. </p> <a style='padding:10px;background-color:royalblue;border:none;border-radius:10px;color:white;font-size:20px;' type='btn' href='https://glomate.netlify.app'>Check file out<a/>")
		return
	}
	helpers.SendMail("New File Alert", "sunkanmiadewumi1@gmail.com", "New File Alert", "<h3>A file was added</h3><br/><p>A new file has been added to glomate database, rememeber to share this across the devices that is needed. </p> <a style='padding:10px;background-color:royalblue;border:none;border-radius:10px;color:white;font-size:20px;' type='btn' href='https://glomate.netlify.app'>Check file out<a/>")

}
