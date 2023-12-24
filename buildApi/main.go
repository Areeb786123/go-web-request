package main

// model for courses
type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice string  `json:"coursePrice"`
	Author      *Author `json:"author"` // using ` ` because it works as @Serialzed name as we do it  in  android

}

type Author struct {
	AuthorName string `json:"author"`
}

// Fake DB
var course []Course

func main() {

}
