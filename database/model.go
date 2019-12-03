package database

type Course struct {
	Id         uint64 `gorm:"column:id; primary_key"`
	Name       string `gorm:"column:name" json:"name"`
	CourseCode string `gorm:"course_code" json:"course_code"`
	Credit     string `gorm:"credit" json:"credit"`
}

type Class struct {
	Id          uint64 `gorm:"column:id; primary_key"`
	Name        string `gorm:"column:name" json:"name"`
	CourseCode  string `gorm:"column:course_code" json:"course_code"`
	ClassCode   string `gorm:"column:class_code" json:"class_code"`
	Cap         string `gorm:"column:cap" json:"cap"`
	TeachingWay string `gorm:"column:teaching_way" json:"teaching_way"`
	Teachers    string `gorm:"column:teachers" json:"teachers"`
	Duty        string `gorm:"column:duty" json:"duty"`
	Time1       string `gorm:"column:time1" json:"time1"`
	Place1      string `gorm:"column:place1" json:"place1"`
	Time2       string `gorm:"column:time2" json:"time2"`
	Place2      string `gorm:"column:place2" json:"place2"`
	Time3       string `gorm:"column:time3" json:"time3"`
	Place3      string `gorm:"column:place3" json:"place3"`
}
