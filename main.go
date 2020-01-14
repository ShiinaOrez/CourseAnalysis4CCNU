package main

import (
	"flag"
	"fmt"
	"github.com/ShiinaOrez/CourseAnalysis4CCNU/database"
	"github.com/ShiinaOrez/CourseAnalysis4CCNU/web"
	"github.com/ShiinaOrez/CourseAnalysis4CCNU/xlsxAnalyst"
	"os"
)

var (
	a bool // analysis
	d bool // database init
	s bool // web serve

	h bool // help
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")

	flag.BoolVar(&a, "a", false, "analysis xlsx")
	flag.BoolVar(&d, "d", false, "insert data to database")
	flag.BoolVar(&s, "s", false, "serve apis at localhost")
	flag.Usage = usage
}

func main() {
	flag.Parse()
	if h {
		flag.Usage()
	}
	if a {
		xlsxAnalyst.Analysis()
		fmt.Println("Class  Tot:", len(xlsxAnalyst.ClassList))
		fmt.Println("Course Tot:", len(xlsxAnalyst.CourseMap))
	}
	if a && d {
		database.InsertCoursesByMap()
		database.InsertClassesByList()
	} else if d {
		fmt.Println("Please choose `d` with `a`")
	}
	if s {
		web.StartServe()
	}
	return
}

func usage() {
	fmt.Fprintf(os.Stderr, `CCNU Course And Class Analyst version: 0.2.0
Usage: ./analyst [-a analysis] [-d database] [-h help] [-s serve]

Options:
`)
	flag.PrintDefaults()
}
