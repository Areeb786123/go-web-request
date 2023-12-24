package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// model for courses
type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice string  `json:"coursePrice"`
	Author      *Author `json:"author"`

	/** using ` ` because it works as
	@Serialzed name as we do it  in  android **/

}

type Author struct {
	AuthorName string `json:"author"`
}

// Fake DB
var courses []Course

func main() {

}

// controllers
func homeController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to home controller</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
