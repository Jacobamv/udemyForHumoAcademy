package services

import (
	"udemy/pkg"
	"udemy/pkg/models"
)

func GetCoursesByUserIndex(index int) []models.Course {
	courses := make([]models.Course, 0)

	user := pkg.Users[index]

	for _, v := range pkg.Subscriptions {
		if user == v.User {
			courses = append(courses, v.Course)
		}
	}

	return courses
}

// Пройдемся по подпискам (subscriptions).
// если subscription.user является нашим то
// мы должны получить его курс и добавить в конечный массив
