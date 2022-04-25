package begin

import (
	"encoding/json"
	"net/http"
	"server/cors"
	"server/utilities"
)

func Start() {
	addRoutes()
	// Sample()
}

func addRoutes() {
	utilities.App.R.HandleFunc("/api/test", Test).Methods("GET", "OPTIONS")
}

func Test(w http.ResponseWriter, r *http.Request) {
	cors.SetupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{"message": "Hello, World!"})
}

// func Sample() {
// 	var courses []models.Course
// 	utilities.App.DB.Table("courses").Find(&courses)
// 	fmt.Println(courses[0].TuesdayStartTime)
// 	fmt.Println(courses[0].TuesdayEndTime)

// 	t, err := time.Parse("0"+courses[0].TuesdayStartTime, "12:00")
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(t.Sub(t).Seconds())
// }
