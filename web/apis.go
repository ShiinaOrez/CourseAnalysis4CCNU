package web

import (
	"encoding/json"
	"github.com/ShiinaOrez/CourseAnalysis4CCNU/database"
	"net/http"
)

type message struct {
	Content string `json:"content"`
}

func searchCourses(w http.ResponseWriter, r *http.Request) {
	kw := r.URL.Query().Get("keyword")
	if kw == "" {
		res, _ := json.Marshal(message{
			Content: "Keyword NOT Be Empty String",
		})
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	} else {
		courses, _ := database.SearchCoursesByKeyword(kw)
		res, _ := json.Marshal(courses)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
	return
}

func searchClasses(w http.ResponseWriter, r *http.Request) {
	kw := r.URL.Query().Get("keyword")
	if kw == "" {
		res, _ := json.Marshal(message{
			Content: "Keyword NOT Be Empty String",
		})
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	} else {
		classes, _ := database.SearchClassesByKeyword(kw)
		res, _ := json.Marshal(classes)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
	return
}
