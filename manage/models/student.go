package models

//表结构
type Studentgrade struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Grade float64 `json:"grade"`
}

//查询表中所有数据
func GetStudent() (stu []Studentgrade) {
	db.Table("studentGrade").Find(&stu)
	return
}

//插入数据
func InsertStudent(stu Studentgrade) bool {
	db.Table("studentGrade").Create(stu)
	if db.NewRecord(stu) {
		return true
	}
	return false
}

//根据学号设置成绩
func SetGrade(id int, grade float64) bool {
	db.Table("studentGrade").Model(&Studentgrade{}).Where("id=?", id).Update("grade", grade)
	return true
}

//成绩升序输出所有学生
func SortGrade() (stu []Studentgrade) {
	db.Table("studentGrade").Order("grade").Find(&stu)
	return
}

//删除学生
func DeleteById(id int) bool {
	db.Table("studentGrade").Where("id=?", id).Delete(&Studentgrade{})
	return true
}
