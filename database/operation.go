package database

func (c *Course) TableName() string {
	return "course"
}

func (c *Class) TableName() string {
	return "class"
}

func (c *Course) Insert() error {
	d := db.Create(c)
	return d.Error
}

func (c *Class) Insert() error {
	d := db.Create(c)
	return d.Error
}

func SearchCoursesByKeyword(kw string) ([]Course, error) {
	courses := &[]Course{}
	db.Table("course").Where("MATCH (`name`, `course_code`) AGAINST ('" + kw + "') ").Find(courses)
	return *courses, nil
}

func SearchClassesByKeyword(kw string) ([]Class, error) {
	classes := &[]Class{}
	db.Table("class").Where("MATCH (`name`, `course_code`, `teachers`) AGAINST ('" + kw + "') ").Find(classes)
	return *classes, nil
}
