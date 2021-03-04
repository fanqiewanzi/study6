package models

//表结构
type Studentgrade struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Grade float64 `json:"grade"`
}

//查询表中所有数据
func GetStudent() (stu []Studentgrade) {
	db.Find(&stu)
	return
}

//插入数据
func InsertStudent(stu Studentgrade) bool {
	var stus Studentgrade
	db.Where("id=?", stu.Id).Find(&stus)
	if stus.Id != 0 {
		return false
	}
	db.Create(stu)
	return true
}

//根据学号设置成绩
func SetGrade(id int, grade float64) bool {
	db.Model(&Studentgrade{}).Where("id=?", id).Update("grade", grade)
	return true
}

//成绩升序输出所有学生
func SortGrade() (stu []Studentgrade) {
	db.Order("grade").Find(&stu)
	return
}

//删除学生
func DeleteById(id int) bool {
	db.Where("id=?", id).Delete(&Studentgrade{})
	return true
}
