package models

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
	R  *mux.Router
}

type Course struct {
	ID                 string `gorm:"primary_key" json:"id"`
	Name               string `json:"name"`
	Code               string `json:"code"`
	Monday             bool   `json:"monday"`
	Tuesday            bool   `json:"tuesday"`
	Wednesday          bool   `json:"wednesday"`
	Thursday           bool   `json:"thursday"`
	Friday             bool   `json:"friday"`
	MondayStartTime    string `json:"monday_start_time"`
	MondayEndTime      string `json:"monday_end_time"`
	TuesdayStartTime   string `json:"tuesday_start_time"`
	TuesdayEndTime     string `json:"tuesday_end_time"`
	WednesdayStartTime string `json:"wednesday_start_time"`
	WednesdayEndTime   string `json:"wednesday_end_time"`
	ThursdayStartTime  string `json:"thursday_start_time"`
	ThursdayEndTime    string `json:"thursday_end_time"`
	FridayStartTime    string `json:"friday_start_time"`
	FridayEndTime      string `json:"friday_end_time"`
	Instructor         string `json:"instructor"`
	Department         string `json:"department"`
}

type CourseRequest struct {
	ID                 string `gorm:"primary_key" json:"id"`
	Name               string `json:"name"`
	Code               string `json:"code"`
	Days               []bool `json:"days"`
	MondayStartTime    string `json:"monday_start_time"`
	MondayEndTime      string `json:"monday_end_time"`
	TuesdayStartTime   string `json:"tuesday_start_time"`
	TuesdayEndTime     string `json:"tuesday_end_time"`
	WednesdayStartTime string `json:"wednesday_start_time"`
	WednesdayEndTime   string `json:"wednesday_end_time"`
	ThursdayStartTime  string `json:"thursday_start_time"`
	ThursdayEndTime    string `json:"thursday_end_time"`
	FridayStartTime    string `json:"friday_start_time"`
	FridayEndTime      string `json:"friday_end_time"`
	Instructor         string `json:"instructor"`
	Department         string `json:"department"`
}

type CourseReply struct {
	ID         string   `gorm:"primary_key" json:"id"`
	Name       string   `json:"name"`
	Code       string   `json:"code"`
	Instructor string   `json:"instructor"`
	Department string   `json:"department"`
	Days       []string `json:"days"`
	StartTimes []string `json:"start_times"`
	EndTimes   []string `json:"end_times"`
}

type Project struct {
	ID         string `gorm:"primaryKey" json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Email      string `json:"email"`
	Link       string `json:"link"`
}

type User struct {
	ID        string `gorm:"primaryKey" json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `gorm:"unique" json:"email"`
	Password  string `json:"password"`
	Isadmin   bool   `json:"isadmin"`
	Created   string `json:"created"`
}

type LoginSignupReply struct {
	Userdata User   `json:"userdata"`
	Message  string `json:"message"`
	Allow    bool   `json:"allow"`
	Token    string `json:"token"`
}

type JWTToken struct {
	Token string `json:"token"`
}
