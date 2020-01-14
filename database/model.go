package database

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

type Course struct {
	Id         uint64 `gorm:"column:id; primary_key"`
	Name       string `gorm:"column:name" json:"name"`
	CourseCode string `gorm:"course_code" json:"course_code"`
	Credit     string `gorm:"credit" json:"credit"`
}

type Class struct {
	Id          uint64 `gorm:"column:id; primary_key"`
	Name        string `gorm:"column:name" json:"name"`
	Academy     string `gorm:"column:academy" json:"academy"`
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

func (course *Course) TypeFilter(t string) bool {
	if t == "1" || t == "2" || t == "3" || t == "5" || t == "0" {
		return course.CourseCode[3] == t[0]
	}
	return false
}

func (class *Class) TypeFilter(t string) bool {
	if t == "1" || t == "2" || t == "3" || t == "5" || t == "0" {
		return class.CourseCode[3] == t[0]
	}
	return false
}

func (class *Class) AcademyFilter(a string) bool {
	return class.Academy == a
}

func (class *Class) WeekdayFilter(w string) bool {
	if class.Time1 != "" {
		if class.Time1[0] == w[0] {
			return true
		} else {
			if class.Time2 != "" {
				if class.Time2[0] == w[0] {
					return true
				} else {
					if class.Time3 != "" {
						return class.Time3[0] == w[0]
					}
				}
			}
		}
	}
	return false
}

func (class *Class) PlaceFilter(p string) bool {
	if p == "本校区" {
		return class.Place1[0] != 'N' &&
			(class.Place2 == "" || class.Place2[0] != 'N') &&
			(class.Place3 == "" || class.Place3[0] != 'N')
	}
	if p == "南湖校区" {
		return class.Place1[0] == 'N' &&
			(class.Place2 == "" || class.Place2[0] == 'N') &&
			(class.Place3 == "" || class.Place3[0] == 'N')
	}
	return true
}
