package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	PassWordCost = 12 //密码加密难度
)

type User struct {
	gorm.Model
	UserName       string `gorm:"Unique"`
	PasswordDigest string //密文
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

func FindUser(UserName string) int64 {
	var count int64
	DB.Model(&User{}).Where("user_name=?", UserName).First(&User{}).Count(&count)
	return count
}
