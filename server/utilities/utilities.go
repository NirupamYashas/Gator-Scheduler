package utilities

import (
	"fmt"
	"server/models"
	"strings"
	"time"

	"github.com/google/uuid"
)

const SecretKey = "secret"

var App models.App

func Intersection(s1, s2 []string) (inter []string) {
	hash := make(map[string]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		// If elements present in the hashmap then append intersection list.
		if hash[e] {
			inter = append(inter, e)
		}
	}
	//Remove dups from slice.
	inter = removeDups(inter)
	return
}

//Remove dups from slice.
func removeDups(elements []string) (nodups []string) {
	encountered := make(map[string]bool)
	for _, element := range elements {
		if !encountered[element] {
			nodups = append(nodups, element)
			encountered[element] = true
		}
	}
	return
}

func IndexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func CheckClashingCourses(course models.CourseReply, courses []models.CourseReply) []models.CourseReply {
	var clashingCourses []models.CourseReply

	for _, c := range courses {
		var clasingDays = Intersection(course.Days, c.Days)

		if len(clasingDays) > 0 {

			for _, day := range clasingDays {

				courseDayIndex := IndexOf(day, course.Days)
				cDayIndex := IndexOf(day, c.Days)

				if ChangeTimeFormat(course.StartTimes[courseDayIndex]) <= ChangeTimeFormat(c.EndTimes[cDayIndex]) && ChangeTimeFormat(course.EndTimes[courseDayIndex]) >= ChangeTimeFormat(c.StartTimes[cDayIndex]) {
					clashingCourses = append(clashingCourses, c)
					break
				}
			}
		}
	}

	return clashingCourses
}

func CourseRequestToCourse(courseRequest models.CourseRequest) models.Course {
	var course models.Course
	course.ID = uuid.New().String()
	course.Name = courseRequest.Name
	course.Department = courseRequest.Department
	course.Instructor = courseRequest.Instructor
	course.Code = courseRequest.Code
	course.Monday = courseRequest.Days[0]
	course.Tuesday = courseRequest.Days[1]
	course.Wednesday = courseRequest.Days[2]
	course.Thursday = courseRequest.Days[3]
	course.Friday = courseRequest.Days[4]
	course.MondayStartTime = courseRequest.MondayStartTime
	course.MondayEndTime = courseRequest.MondayEndTime
	course.TuesdayStartTime = courseRequest.TuesdayStartTime
	course.TuesdayEndTime = courseRequest.TuesdayEndTime
	course.WednesdayStartTime = courseRequest.WednesdayStartTime
	course.WednesdayEndTime = courseRequest.WednesdayEndTime
	course.ThursdayStartTime = courseRequest.ThursdayStartTime
	course.ThursdayEndTime = courseRequest.ThursdayEndTime
	course.FridayStartTime = courseRequest.FridayStartTime
	course.FridayEndTime = courseRequest.FridayEndTime

	return course
}

func CoursesToCourseReplies(courses []models.Course) []models.CourseReply {
	var courseReplies []models.CourseReply
	for _, course := range courses {
		courseReplies = append(courseReplies, CourseToCourseReply(course))
	}
	return courseReplies
}

func CourseToCourseReply(course models.Course) models.CourseReply {
	var courseReply models.CourseReply
	courseReply.ID = course.ID
	courseReply.Name = course.Name
	courseReply.Department = course.Department
	courseReply.Instructor = course.Instructor
	courseReply.Code = course.Code

	if course.Monday {
		courseReply.Days = append(courseReply.Days, "Monday")
		courseReply.StartTimes = append(courseReply.StartTimes, course.MondayStartTime)
		courseReply.EndTimes = append(courseReply.EndTimes, course.MondayEndTime)
	}

	if course.Tuesday {
		courseReply.Days = append(courseReply.Days, "Tuesday")
		courseReply.StartTimes = append(courseReply.StartTimes, course.TuesdayStartTime)
		courseReply.EndTimes = append(courseReply.EndTimes, course.TuesdayEndTime)
	}

	if course.Wednesday {
		courseReply.Days = append(courseReply.Days, "Wednesday")
		courseReply.StartTimes = append(courseReply.StartTimes, course.WednesdayStartTime)
		courseReply.EndTimes = append(courseReply.EndTimes, course.WednesdayEndTime)
	}

	if course.Thursday {
		courseReply.Days = append(courseReply.Days, "Thursday")
		courseReply.StartTimes = append(courseReply.StartTimes, course.ThursdayStartTime)
		courseReply.EndTimes = append(courseReply.EndTimes, course.ThursdayEndTime)
	}

	if course.Friday {
		courseReply.Days = append(courseReply.Days, "Friday")
		courseReply.StartTimes = append(courseReply.StartTimes, course.FridayStartTime)
		courseReply.EndTimes = append(courseReply.EndTimes, course.FridayEndTime)
	}

	return courseReply
}

func ChangeTimeFormat(timestring string) string {
	layout1 := "03:04 PM"
	layout2 := "15:04"

	if len(strings.Split(timestring, ":")[0]) == 1 {
		t, err := time.Parse(layout1, "0"+timestring)

		if err != nil {
			fmt.Println(err)
		}

		return t.Format(layout2)
	} else {
		t, err := time.Parse(layout1, timestring)

		if err != nil {
			fmt.Println(err)
		}

		return t.Format(layout2)
	}

}
