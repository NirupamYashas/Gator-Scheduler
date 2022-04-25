package courses

import (
	"encoding/json"
	"net/http"
	"server/cors"
	"server/models"
	"server/utilities"
	"strings"

	"github.com/gorilla/mux"
)

func Start() {
	utilities.App.DB.AutoMigrate(&models.Course{})
	addRoutes()
}

func addRoutes() {
	utilities.App.R.HandleFunc("/api/courses", GetCourses).Methods("GET", "OPTIONS")
	utilities.App.R.HandleFunc("/api/courses", AddCourse).Methods("POST", "OPTIONS")
	// utilities.App.R.HandleFunc("/api/projects/{id}", UpdateProject).Methods("PUT", "OPTIONS")
	// utilities.App.R.HandleFunc("/api/projects/{id}", DeleteProject).Methods("DELETE", "OPTIONS")
	utilities.App.R.HandleFunc("/api/courses/departments/{department}", GetCoursesByDepartment).Methods("GET", "OPTIONS")
	utilities.App.R.HandleFunc("/api/courses/search/{search_phrase}", GetCoursesBySearch).Methods("GET", "OPTIONS")
}

func GetCourses(w http.ResponseWriter, r *http.Request) {
	cors.SetupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var courses []models.Course
	err := utilities.App.DB.Table("courses").Find(&courses).Error

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}

	var courseReplies []models.CourseReply = utilities.CoursesToCourseReplies(courses)

	err = json.NewEncoder(w).Encode(courseReplies)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
}

func AddCourse(w http.ResponseWriter, r *http.Request) {
	cors.SetupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var courseRequest models.CourseRequest
	err := json.NewDecoder(r.Body).Decode(&courseRequest)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}

	var course models.Course = utilities.CourseRequestToCourse(courseRequest)

	err = utilities.App.DB.Table("courses").Save(&course).Error

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func GetCoursesByDepartment(w http.ResponseWriter, r *http.Request) {
	cors.SetupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var courses []models.Course
	err := utilities.App.DB.Table("courses").Find(&courses, "department = ?", mux.Vars(r)["department"]).Error

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}

	err = json.NewEncoder(w).Encode(utilities.CoursesToCourseReplies(courses))

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
}

func GetCoursesBySearch(w http.ResponseWriter, r *http.Request) {
	cors.SetupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	res := strings.Split(mux.Vars(r)["search_phrase"], " ")
	tx := utilities.App.DB.Table("courses")

	for _, element := range res {
		search_term := "%" + element + "%"
		tx = tx.Where("name LIKE ? OR instructor LIKE ?", search_term, search_term)
	}

	var courses []models.Course
	err := tx.Find(&courses).Error

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}

	err = json.NewEncoder(w).Encode(utilities.CoursesToCourseReplies(courses))

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
}

func GetClashingCourses(w http.ResponseWriter, r *http.Request) {
	cors.SetupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var course models.Course
	err := json.NewDecoder(r.Body).Decode(&course)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}

	var courses []models.Course
	err = utilities.App.DB.Table("courses").Find(&courses).Error

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}

	var clashingCourses []models.CourseReply = utilities.CheckClashingCourses(utilities.CourseToCourseReply(course), utilities.CoursesToCourseReplies(courses))

	err = json.NewEncoder(w).Encode(clashingCourses)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
}
