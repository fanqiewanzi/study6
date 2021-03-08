package models

type User struct {
	Model
	UserName string `json:"userName" gorm:"column:userName"`
	Password string `json:"password" gorm:"column:password"`
}

//通过用户名查询用户
func GetUserByName(user User) (u User) {
	db.Table("users").Where("userName=?", user.UserName).Find(&u)
	return u
}

//插入新记录
func InsertUser(user User) bool {
	var u []User

	//查找是否存在相同用户名的
	db.Table("users").Where("userName=?", user.UserName).Find(&u)
	if len(u) != 0 {
		return false
	}
	//创建新用户
	db.Table("users").Create(&user)
	//判断是否新增成功
	if !db.Table("users").NewRecord(user) {
		return true
	}
	return false
}
