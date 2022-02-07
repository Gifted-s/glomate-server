package controller

import (
	"encoding/json"
	"fmt"
	"glomate/server/helpers"
	"glomate/server/models"
	"log"

	//"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	//"fmt"
	// "github.com/gorilla/websocket"
	// "io"
	"context"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// HandleCreateBlog this will help us create blog items
func HandleUploadCourse(w http.ResponseWriter, r *http.Request) {
	collections := helpers.ConnectDB()
	var course models.Course
	//var file models.FileStruct
	var resp models.ResponseStruct
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	if course.Lectures[0].Type == "text/UTF-8" {
		course.Lectures[0].Size = float64(len([]byte(course.Lectures[0].Download_Link))) / 1000 // get link size and convert to kilobyte
	}

	filter := bson.M{"name": course.Name}
	count, _err := collections.Courses.CountDocuments(context.TODO(), filter)
	if _err != nil {
		panic(_err)
	}
	// course does not exist yet
	if count == 0 {
		department := models.Departments{
			School:       course.School,
			Name:         course.Dept,
			Last_Updated: course.Lectures[0].Date_Added,
		}
		var session mongo.Session
		if session, err = helpers.GetDBClient().StartSession(); err != nil {
			fmt.Println("fialed to get client", err)
		}
		if err = session.StartTransaction(); err != nil {
			fmt.Println("fialed to get client", err)
		}
		if err = mongo.WithSession(context.Background(), session, func(sc mongo.SessionContext) error {
			filter := bson.M{"name": course.Dept}
			count, err := collections.Departments.CountDocuments(context.TODO(), filter)
			if count == 0 {
				_, department_insert_err := collections.Departments.InsertOne(context.Background(), department)
				fmt.Print("Inserted into Database")
				if department_insert_err != nil {
					log.Fatal("Could not insert new Department", err)
				}
			}
			fmt.Print(count)
			// string to primitive.ObjectID
			_, course_insert_err := collections.Courses.InsertOne(context.Background(), course)
			if course_insert_err != nil {
				log.Fatal("Could not insert Course", err)
			}

			if err = session.CommitTransaction(sc); err != nil {
				log.Fatal("Could not insert Course", err)
			}
			return nil
		}); err != nil {
			log.Fatal(err)
		}
		session.EndSession(context.Background())
		resp = models.ResponseStruct{Status: "success", Body: map[string]string{"message": "Course insert successfull"}, Token: ""}

	} else {
		change := bson.M{
			"$push": bson.M{
				"lectures": course.Lectures[0],
			},
			"$set": bson.M{
				"last_updated": course.Lectures[0].Date_Added,
			},
		}

		update_result, err := collections.Courses.UpdateOne(context.TODO(), filter, change)
		if err != nil {
			// If there is an error in hasing password
			err := models.ResponseStruct{Status: "failed", Body: map[string]string{"error": "Counld Not Insert  Request To DB please try again"}}
			json.NewEncoder(w).Encode(err)
			return
		}

		resp = models.ResponseStruct{Status: "success", Body: update_result, Token: ""}

	}

	json.NewEncoder(w).Encode(resp)
	if course.Lectures[0].Type == "text/UTF-8" {
		helpers.SendMail("New Link Alert", "sunkanmiadewumi1@gmail.com", "New Link Alert", "<h3>A Link was added</h3><br/><p>A new link has been added to glomate database, rememeber to share this across the devices that is needed. </p> <a style='padding:10px;background-color:royalblue;border:none;border-radius:10px;color:white;font-size:20px;' type='btn' href='https://glomate.netlify.app'>Check file out<a/>")
		return
	}
	helpers.SendMail("New File Alert", "sunkanmiadewumi1@gmail.com", "New File Alert", "<h3>A file was added</h3><br/><p>A new file has been added to glomate database, rememeber to share this across the devices that is needed. </p> <a style='padding:10px;background-color:royalblue;border:none;border-radius:10px;color:white;font-size:20px;' type='btn' href='https://glomate.netlify.app'>Check file out<a/>")

}
