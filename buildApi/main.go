package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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
	r := mux.NewRouter()

	courses = append(courses, Course{CourseId: "1", CourseName: "Android", CoursePrice: "233", Author: &Author{AuthorName: "Areeb"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "UI/UX", CoursePrice: "103", Author: &Author{AuthorName: "Jerry"}})
	//seeding
	r.HandleFunc("/", homeController).Methods("GET")
	r.HandleFunc("/allCourses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/addCourse", addCourse).Methods("POST")
	r.HandleFunc("/updateCourse/{id}", updateCourse).Methods("UPDATE")
	r.HandleFunc("/deleteCourse/{id}", deleteCousreById).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers
func homeController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to home controller</h1>"))
}

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func addCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//what if entire body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please add some data ")
	}

	//what about data which is send like {} empty json
	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("empty value are not allowed")
		return
	}

	//generate random id and convert it in a string
	rand.Seed(time.Now().UnixNano())
	// creating random number and converting it in a string
	/**
	We can create courseID as int but for learning purpose of a string we are using it
	**/
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}

	}
	json.NewEncoder(w).Encode("<h1>Course not found </h1>")
	return

}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//grab id  from the request

	/**
	loop through course
	find id
	update id
	**/

	for index, course := range courses {
		params := mux.Vars(r)
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
		}
	}

	// send response when id is not found

}

func deleteCousreById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
	//send json response of successfull message
}
