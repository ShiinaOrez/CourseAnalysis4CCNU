package database

import (
	"fmt"
	"github.com/ShiinaOrez/CourseAnalysis4CCNU/xlsxAnalyst"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
)

var db *gorm.DB

func init() {
	var err error
	connStr := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("COURSE_DB_USERNAME"),
		os.Getenv("COURSE_DB_PASSWORD"),
		os.Getenv("COURSE_DB_HOST"),
		os.Getenv("COURSE_DB_NAME"))
	db, err = gorm.Open("mysql", connStr)
	if err != nil {
		fmt.Println("Connect to database failed.")
		db = nil
	}
}

func InsertCoursesByMap() {
	for _, course := range xlsxAnalyst.CourseMap {
		newCourse := Course{
			Name:       course.Name,
			CourseCode: course.CourseCode,
			Credit:     course.Credit,
		}
		err := newCourse.Insert()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Insert", newCourse.Name, "successful.")
	}
}

func InsertClassesByList() {
	for _, class := range xlsxAnalyst.ClassList {
		teacherStr := ""
		for _, teacher := range class.Teachers {
			teacherStr = teacherStr + teacher.Code + "/"
			teacherStr = teacherStr + teacher.Name + "/"
		}
		newClass := Class{
			CourseCode:  class.CourseCode,
			ClassCode:   class.ClassCode,
			Name:        class.Name,
			Cap:         class.Cap,
			TeachingWay: class.TeachingWay,
			Teachers:    teacherStr,
			Duty:        class.Duty,
			Time1:       class.Time1.Range,
			Place1:      class.Time1.Place,
			Time2:       class.Time2.Range,
			Place2:      class.Time2.Place,
			Time3:       class.Time3.Range,
			Place3:      class.Time3.Place,
		}
		err := newClass.Insert()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Insert", newClass.Name, "successful.")
	}
}
