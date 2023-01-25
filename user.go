package WhySingletonDao

import (
	"fmt"
	"gorm.io/gorm"
	"sync"
)

type User struct {
	Id    int64  `gorm:"column:id"`
	Name  string `gorm:"column:name"`
	Value string `gorm:"column:value"`
}

func (User) TableName() string {
	return "t_user1"
}

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}
func (*UserDao) InsertUser(name string, value string) error {
	user := User{Name: name, Value: value}
	err := db.Create(&user).Error
	if err != nil {
		fmt.Println("find user by id err:" + err.Error())
		return err
	}
	return nil
}
func (*UserDao) QueryUserById(username string) (*User, error) {
	var user User
	err := db.Where("name = ?", username).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		fmt.Println("find user by id err:" + err.Error())
		return nil, err
	}
	return &user, nil
}
