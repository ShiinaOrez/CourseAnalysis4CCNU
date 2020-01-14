package web

import (
	"github.com/ShiinaOrez/CourseAnalysis4CCNU/database"
)

func coursesFilter(courses []database.Course, t string) []database.Course {
	if t == "" {
		return courses
	}
	newCourses := make([]database.Course, 0)
	for _, course := range courses {
		keep := true
		if t != "" {
			keep = keep && course.TypeFilter(t)
		}
		if keep {
			newCourses = append(newCourses, course)
		}
	}
	return newCourses
}

func classesFilter(classes []database.Class, t, a, w, p string) []database.Class {
	// t is type
	// a is academy
	// w is weekday
	// p is place
	if t == "" && a == "" && w == "" && p == "" {
		return classes
	}
	newClasses := make([]database.Class, 0)
	for _, class := range classes {
		keep := true
		if t != "" {
			keep = keep && class.TypeFilter(t)
		}
		if a != "" {
			keep = keep && class.AcademyFilter(a)
		}
		if w != "" {
			keep = keep && class.WeekdayFilter(w)
		}
		if p != "" {
			keep = keep && class.PlaceFilter(p)
		}
		if keep {
			newClasses = append(newClasses, class)
		}
	}
	return newClasses
}
