package web

import (
	"fmt"
	"net/http"
)

func StartServe() {
	http.HandleFunc("/search/course/", searchCourses)
	http.HandleFunc("/search/class/", searchClasses)
	fmt.Println("API:")
	fmt.Println("  - localhost:2048/search/course/?keyword=")
	fmt.Println("  - localhost:2048/search/class/?keyword=")
	http.ListenAndServe(":2048", nil)
}
