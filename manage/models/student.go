package models

//表结构
type Studentgrade struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Grade float32 `json:"grade"`
}

//查询表中所有数据
func GetStudent() (stu []Studentgrade) {
	db.Find(&stu)
	return
}

//插入数据
func InsertStudent(maps map[string]interface{}) bool {
	db.Model(Studentgrade{}).Create(maps)
	if db.NewRecord(maps) {
		return true
	}
	return false
}
