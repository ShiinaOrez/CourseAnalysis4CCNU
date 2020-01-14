package xlsxAnalyst

import (
	"errors"
	"fmt"
	"github.com/tealeg/xlsx"
	"strings"
)

type TeacherEntity struct {
	Code string
	Name string
}

type TeachingTime struct {
	Place string
	Range string // format: 1#7-8#2-17(A) 星期一, 七到八节, 二到十七周, A: 全部, S: 单周, D: 双周
}

type Course struct {
	Name       string
	CourseCode string
	Credit     string
}

type Class struct {
	Name        string
	Academy     string
	CourseCode  string
	ClassCode   string
	Cap         string
	TeachingWay string
	Teachers    []TeacherEntity
	Duty        string
	Time1       TeachingTime
	Time2       TeachingTime
	Time3       TeachingTime
}

var wb *xlsx.File

var CourseMap map[string]Course
var ClassList []Class

func getRange(src string) string {
	if src == "" {
		return ""
	}
	res := ""
	if strings.HasPrefix(src, "星期一") {
		res = "1#"
	} else if strings.HasPrefix(src, "星期二") {
		res = "2#"
	} else if strings.HasPrefix(src, "星期三") {
		res = "3#"
	} else if strings.HasPrefix(src, "星期四") {
		res = "4#"
	} else if strings.HasPrefix(src, "星期五") {
		res = "5#"
	} else if strings.HasPrefix(src, "星期六") {
		res = "6#"
	} else if strings.HasPrefix(src, "星期日") {
		res = "7#"
	}
	res += src[12:strings.Index(src, "节")]
	res += "#"
	res += src[strings.IndexByte(src, '{')+1 : strings.Index(src, "周")]
	if strings.Index(src, "单") != -1 {
		res += "(S)"
	} else if strings.Index(src, "双") != -1 {
		res += "(D)"
	} else {
		res += "(A)"
	}
	return res
}

func preloadFile() {
	var err error
	wb = nil
	if wb, err = xlsx.OpenFile("book.xlsx"); err != nil {
		fmt.Println("Pre-load file failed! Reason:", err.Error())
		wb = nil
	} else {
		fmt.Println("Pre-load file successful.")
	}
	CourseMap = make(map[string]Course)
}

func showSheetsName(tot int) {
	fmt.Printf("Book has %d sheets:\n", tot)
	for i := 0; i < tot; i++ {
		fmt.Println(i, "|", wb.Sheets[i].Name)
	}
}

func analysisSheet(sheetIndex int) error {
	// var gradeIndex, toIndex int
	var (
		nameIndex       = 0
		academyIndex    = 0
		courseCodeIndex = 0
		classCodeIndex  = 0
		creditIndex     = 0
		capIndex        = 0
		twIndex         = 0
		teacherIndex    = 0
		dutyIndex       = 0
		t1Index         = 0
		t2Index         = 0
		t3Index         = 0
	)
	sheet := wb.Sheets[sheetIndex]
	if sheet.MaxRow == 0 {
		fmt.Println("Sheet", sheet.Name, "is empty.")
	} else {
		headerRow := sheet.Rows[0]
		cellsTot := len(headerRow.Cells)
		for i := 0; i < cellsTot; i++ {
			switch headerRow.Cells[i].Value {
			/*case "年级":
				gradeIndex = i
			case "授课对象":
				toIndex = i*/
			case "课程名称":
				nameIndex = i
			case "开课学院":
				academyIndex = i
			case "课程编号":
				courseCodeIndex = i
			case "课堂号":
				classCodeIndex = i
			case "学分":
				creditIndex = i
			case "容量":
				capIndex = i
			case "教学方式":
				twIndex = i
			case "教师姓名":
				teacherIndex = i
			case "教师职称":
				dutyIndex = i
			case "上课时间1":
				t1Index = i
			case "上课时间2":
				t2Index = i
			case "上课时间3":
				t3Index = i
			}
		}
		for i := 1; i < int(sheet.MaxRow); i++ {
			rowNow := sheet.Rows[i]
			teacherStrs := strings.Split(rowNow.Cells[teacherIndex].Value, ",")
			fmt.Println(teacherStrs)
			teachers := make([]TeacherEntity, 0)
			for _, str := range teacherStrs {
				data := strings.Split(str, "/")
				if len(data) != 2 {
					return errors.New("Course: " + rowNow.Cells[nameIndex].Value + " loading teacher failed.")
				}
				teachers = append(teachers, TeacherEntity{
					Code: data[0],
					Name: data[1],
				})
			}
			var time1, time2, time3 TeachingTime
			if t1Index != 0 {
				time1.Place = rowNow.Cells[t1Index+1].Value
				time1.Range = getRange(rowNow.Cells[t1Index].Value)
			}
			if t2Index != 0 {
				time2.Place = rowNow.Cells[t2Index+1].Value
				time2.Range = getRange(rowNow.Cells[t2Index].Value)
			}
			if t3Index != 0 {
				time3.Place = rowNow.Cells[t3Index+1].Value
				time3.Range = getRange(rowNow.Cells[t3Index].Value)
			}
			newClass := Class{
				Name:        rowNow.Cells[nameIndex].Value,
				Academy:     rowNow.Cells[academyIndex].Value,
				CourseCode:  rowNow.Cells[courseCodeIndex].Value,
				ClassCode:   rowNow.Cells[classCodeIndex].Value,
				Cap:         rowNow.Cells[capIndex].Value,
				TeachingWay: rowNow.Cells[twIndex].Value,
				Teachers:    teachers,
				Duty:        rowNow.Cells[dutyIndex].Value,
				Time1:       time1,
				Time2:       time2,
				Time3:       time3,
			}
			ClassList = append(ClassList, newClass)
			if _, existed := CourseMap[newClass.CourseCode]; !existed {
				CourseMap[newClass.CourseCode] = Course{
					Name:       rowNow.Cells[nameIndex].Value,
					CourseCode: newClass.CourseCode,
					Credit:     rowNow.Cells[creditIndex].Value,
				}
				// fmt.Println(courseMap[newClass.CourseCode])
			}
		}
	}
	return nil
}

func analysisSheets(tot int) error {
	for i := 0; i < tot; i++ {
		err := analysisSheet(i)
		if err != nil {
			return err
		}
	}
	return nil
}

func Analysis() {
	preloadFile()
	if wb == nil {
		return
	}

	wbSheetsTot := len(wb.Sheets)
	if wbSheetsTot == 0 {
		fmt.Println("Book has no sheets to analysis. Program done.")
	} else {
		showSheetsName(wbSheetsTot)
		err := analysisSheets(wbSheetsTot)
		if err != nil {
			fmt.Println("Analysis failed, reason:", err.Error())
		}
	}
	return
}
