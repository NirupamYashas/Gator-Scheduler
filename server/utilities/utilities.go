package utilities

import "server/models"

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

// func CheckClashingCourses(course models.Course, courses []models.Course) []models.Course {
// 	var clashingCourses []models.Course

// 	for _, c := range courses {
// 		var clasingDays = Intersection(course.Days, c.Days)

// 		if len(clasingDays) > 0 {

// 			for _, day := range clasingDays {

// 				courseDayIndex := IndexOf(day, course.Days)
// 				cDayIndex := IndexOf(day, c.Days)

// 				if course.From[courseDayIndex] <= c.To[cDayIndex] && course.To[courseDayIndex] >= c.From[cDayIndex] {
// 					clashingCourses = append(clashingCourses, c)
// 					break
// 				}
// 			}
// 		}
// 	}

// 	return clashingCourses
// }
