package courses

import (
	"encoding/json"
	"net/http"
	"server/cors"
	"server/models"
	"server/utilities"

	"github.com/google/uuid"
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
	// utilities.App.R.HandleFunc("/api/projects/search/{search_phrase}", GetProjectsBySearch).Methods("GET", "OPTIONS")
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

	err = json.NewEncoder(w).Encode(courses)

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

	var course models.Course
	err := json.NewDecoder(r.Body).Decode(&course)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}

	course.ID = uuid.New().String()
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

	err = json.NewEncoder(w).Encode(courses)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
}
