package users

import (
	"encoding/json"
	"net/http"
	"server/cors"
	"server/models"
	"server/utilities"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/golang-jwt/jwt"
)

func Start() {
	utilities.App.DB.AutoMigrate(&models.User{})
	addRoutes()
}

func addRoutes() {
	utilities.App.R.HandleFunc("/api/users/signup", SignupUser).Methods("POST", "OPTIONS")
	utilities.App.R.HandleFunc("/api/users/login", LoginUser).Methods("POST", "OPTIONS")
	utilities.App.R.HandleFunc("/api/users", GetUsers).Methods("GET", "OPTIONS")
	utilities.App.R.HandleFunc("/api/users/{id}", DeleteUser).Methods("DELETE", "OPTIONS")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	cors.SetupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var users []models.User
	err := utilities.App.DB.Table("users").Find(&users).Error

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(users)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	cors.SetupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var user models.User
	vars := mux.Vars(r)
	id := vars["id"]
	user.ID = id
	err := utilities.App.DB.Table("users").Unscoped().Delete(user).Error

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	} else {
		json.NewEncoder(w).Encode("User deleted")
		return
	}
}

func SignupUser(w http.ResponseWriter, r *http.Request) {
	cors.SetupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var user models.User
	var reply models.LoginSignupReply
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		reply = models.LoginSignupReply{Message: "Error in decoding", Allow: false}
		json.NewEncoder(w).Encode(reply)
		return
	}

	err = utilities.App.DB.Table("users").Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error

	if err == gorm.ErrRecordNotFound {
		if user.Firstname == "" || user.Lastname == "" {
			json.NewEncoder(w).Encode(nil)
			return
		}

		user.ID = uuid.New().String()
		user.Isadmin = false
		user.Created = time.Now().String()[:10]
		err = utilities.App.DB.Table("users").Save(&user).Error

		if err != nil {
			json.NewEncoder(w).Encode(nil)
			return
		}

		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"issuer":  user.ID,
			"expires": time.Now().Add(time.Hour * 24).Unix(),
			"data":    user,
		})

		token, err := claims.SignedString([]byte(utilities.SecretKey))

		if err != nil {
			json.NewEncoder(w).Encode(nil)
		}

		json.NewEncoder(w).Encode(models.JWTToken{Token: token})
		return
	}

	json.NewEncoder(w).Encode(nil)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	cors.SetupCorsResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var user models.User
	var reply models.LoginSignupReply
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		reply = models.LoginSignupReply{Message: err.Error(), Allow: false}
		json.NewEncoder(w).Encode(reply)
		return
	}

	err = utilities.App.DB.Table("users").Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error

	if err == gorm.ErrRecordNotFound {
		json.NewEncoder(w).Encode(nil)
		return
	} else if err != nil {
		json.NewEncoder(w).Encode(nil)
		return
	}

	if user.ID != "" {
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"issuer":  user.ID,
			"expires": time.Now().Add(time.Hour * 24).Unix(),
			"data":    user,
		})

		token, err := claims.SignedString([]byte(utilities.SecretKey))

		if err != nil {
			json.NewEncoder(w).Encode(nil)
		}

		json.NewEncoder(w).Encode(models.JWTToken{Token: token})
	} else {
		json.NewEncoder(w).Encode(nil)
	}
}
