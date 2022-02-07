package main

import (
	"fmt"
	"glomate/server/controller"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/add-course", controller.HandleUploadCourse).Methods("POST")
	r.HandleFunc("/get-files", controller.HandleGetFiles).Methods("GET")
	r.HandleFunc("/delete-file/{course}/{id}", controller.HandleDeleteFile).Methods("GET")
	r.HandleFunc("/get-departments", controller.HandleGetDepartments).Methods("GET")
	r.HandleFunc("/get-courses/{dept}/{lev}", controller.HandleGetCourses).Methods("GET")
	r.HandleFunc("/get-course-material/{dept}/{lev}/{course}", controller.HandleGetCourseMaterials).Methods("GET")
	http.Handle("/", r)
	// log.Println("Starting Server on Port 8080")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Print("LISTENING ON PORT: localhost:", port)
	err := http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r))
	if err != nil {
		log.Fatal(err)
	}
}
