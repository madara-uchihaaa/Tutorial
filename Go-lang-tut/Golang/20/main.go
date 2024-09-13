package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// DB
var courses []Course

// Model
type Course struct {
	CourseID   string  `json:"course_id"`
	CourseName string  `json:"course_name"`
	Price      float64 `json:"price"`
	Author     *Author `json:"author"`
}

type Author struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

// Helper
func (c *Course) isEmpty() bool {
	return c.CourseName == "" && c.Price == 0 && c.Author == nil
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		<h1>Welcome to my API</h1>
	`))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(courses)
}
func getCourse(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for _, course := range courses {
		if course.CourseID == id {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Course not found")
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid request body")
		return
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid reques for course")
		return
	}

	rand.Seed(time.Now().UnixNano())
	course.CourseID = string(rand.Intn(100000000))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	for index, item := range courses {
		if item.CourseID == id {
			courses[index] = course
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Course not found")
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for index, item := range courses {
		if item.CourseID == id {
			courses = append(courses[:index], courses[index+1:]...)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Course deleted")
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Course not found")
}

func main() {
	r := mux.NewRouter()

	courses = append(courses, Course{
		CourseID:   "1",
		CourseName: "Golang",
		Price:      100,
		Author: &Author{
			FullName: "John Doe",
			Email:    "xyz@gmail.com",
		},
	})
	courses = append(courses, Course{
		CourseID:   "2",
		CourseName: "Python",
		Price:      200,
		Author: &Author{
			FullName: "John Doe",
			Email:    "risha@gmail.com"},
	})

	r.HandleFunc("/", serverHome).Methods(http.MethodGet)
	r.HandleFunc("/courses", getAllCourses).Methods(http.MethodGet)
	r.HandleFunc("/courses/{id}", getCourse).Methods(http.MethodGet)
	r.HandleFunc("/courses", createCourse).Methods(http.MethodPost)
	r.HandleFunc("/courses/{id}", updateCourse).Methods(http.MethodPut)
	r.HandleFunc("/courses/{id}", deleteCourse).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8080", r))
}
