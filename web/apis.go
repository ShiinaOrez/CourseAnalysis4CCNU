package web

import (
	"encoding/json"
	"github.com/ShiinaOrez/CourseAnalysis4CCNU/database"
	"net/http"
)

/*
筛选条件:
	kw: keyword, 关键字, 使用全文索引进行匹配
	type: 课程类型, 根据课程号的第4位来进行筛选
		1-专业必修课
		2-专业选修课
		3-通识选修课
		5-通识核心课
		0-公共课
	academy: 开课学院, 要求学院名称完全匹配
	weekday: 上课时间, 以 1~7 代表周一到周日
	place: 上课地点片区, 分为南湖和本校区
*/

type message struct {
	Content string `json:"content"`
}

func searchCourses(w http.ResponseWriter, r *http.Request) {
	kw := r.URL.Query().Get("keyword")
	t := r.URL.Query().Get("type")
	if kw == "" {
		res, _ := json.Marshal(message{
			Content: "Keyword NOT Be Empty String",
		})
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	} else {
		courses, _ := database.SearchCoursesByKeyword(kw)
		courses = coursesFilter(courses, t)
		res, _ := json.Marshal(courses)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
	return
}

func searchClasses(w http.ResponseWriter, r *http.Request) {
	kw := r.URL.Query().Get("keyword")
	t := r.URL.Query().Get("type")
	a := r.URL.Query().Get("academy")
	_w := r.URL.Query().Get("weekday")
	p := r.URL.Query().Get("place")
	if kw == "" {
		res, _ := json.Marshal(message{
			Content: "Keyword NOT Be Empty String",
		})
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	} else {
		classes, _ := database.SearchClassesByKeyword(kw)
		classes = classesFilter(classes, t, a, _w, p)
		res, _ := json.Marshal(classes)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
	return
}
