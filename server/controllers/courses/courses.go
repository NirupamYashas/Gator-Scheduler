package courses

import (
	"encoding/json"
	"net/http"
	"server/cors"
	"server/models"
	"server/utilities"
)

func Start() {
	utilities.App.DB.AutoMigrate(&models.Course{})
	addRoutes()
}

func addRoutes() {
	utilities.App.R.HandleFunc("/api/courses", GetCourses).Methods("GET", "OPTIONS")
	// utilities.App.R.HandleFunc("/api/projects", AddProject).Methods("POST", "OPTIONS")
	// utilities.App.R.HandleFunc("/api/projects/{id}", UpdateProject).Methods("PUT", "OPTIONS")
	// utilities.App.R.HandleFunc("/api/projects/{id}", DeleteProject).Methods("DELETE", "OPTIONS")
	// utilities.App.R.HandleFunc("/api/projects/departments/{department}", GetProjectsByDepartment).Methods("GET", "OPTIONS")
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
