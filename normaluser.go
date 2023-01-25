package WhySingletonDao

import (
	"fmt"
	"gorm.io/gorm"
)

type NUser struct {
	Id    int64  `gorm:"column:id"`
	Name  string `gorm:"column:name"`
	Value string `gorm:"column:value"`
}

func (NUser) TableName() string {
	return "t_user2"
}

func InsertUser(name string, value string) error {
	user := NUser{Name: name, Value: value}
	err := db.Create(&user).Error
	if err != nil {
		fmt.Println("find user by id err:" + err.Error())
		return err
	}
	return nil
}
func QueryUserById(username string) (*NUser, error) {
	var user NUser
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
