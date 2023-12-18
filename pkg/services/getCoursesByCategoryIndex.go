package services

import (
	"udemy/pkg"
	"udemy/pkg/models"
)

func GetCoursesByCategoryIndex(index int) []models.Course {
	courses := make([]models.Course, 0)
	category := pkg.Categories[index]

	for _, v := range pkg.Courses {
		if v.Category == category {
			courses = append(courses, v)
		}
	}

	return courses
}
