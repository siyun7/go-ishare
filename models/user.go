package models

type User struct {
	Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Status   int    `json:"status"`
}

type JwtToken struct {
	Token string `json:"token"`
}

func ExistUserByName(name string) bool {
	var tag User

	db.Select("id").Where("name = ?", name).First(&tag)

	if tag.ID > 0 {
		return true
	}

	return false
}

func ExistUserById(id int) bool {
	var tag User

	db.Select("id").Where("id = ?", id).First(&tag)

	if tag.ID > 0 {
		return true
	}

	return false
}

func AddUser(name string, password string) bool {
	db.Create(&User{
		Name:     name,
		Password: password,
	})

	return true
}

func EditUser(id int, data interface{}) bool {
	db.Model(User{}).Where("id = ?", id).Update(data)
	return true
}

func DeleteUser(id int) bool {
	db.Model(User{}).Where("id = ?", id).Delete(User{})
	return true
}

func GetUsers(pageNum int, pageSize int, maps interface{}) (users []User) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&users)
}

