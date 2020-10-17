package models

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"time"

	"github.com/astaxie/beego"
)

type User struct {
	ID        uint      `gorm:"PRIMARY_KEY" form:"_" json:"id"`
	Name      string    `gorm:"Type:varchar(64);UNIQUE;NOT NULL" form:"username" json:"username"` // string默认长度为255
	Password  string    `gorm:"Type:varchar(128);not null;default:'admin'" form:"password" json:"password"`
	CreatedAt time.Time `form:"_" json:"createat"`
}

func InitUsers() {
	getusers := make([]User, 0)
	result := db.Where("name=?", "admin").Find(&getusers)
	if result.RowsAffected <= 0 {
		// md5 of password
		md5Ctx := md5.New()                            //md5 init
		md5Ctx.Write([]byte("password"))               //md5 updata
		cipherStr := md5Ctx.Sum(nil)                   //md5 final
		encryptedData := hex.EncodeToString(cipherStr) //hex_digest

		user := User{Name: "admin", Password: encryptedData}
		if count := db.Create(&user).RowsAffected; count != 1 {
			beego.Error("add user fail!")
		}
	}
}

func GetUser(name string) (error, *User) {
	getuser := User{}
	result := db.Where("name=?", name).First(&getuser)
	if result.Error != nil {
		beego.Error(result.Error)
		return result.Error, nil
	}
	if result.RowsAffected != 1 {
		beego.Error(errors.New("user not found!"))
		return errors.New("user not found!"), nil
	}
	return nil, &getuser
}

func AddUser(username, password string) error {
	if username == "" || password == "" {
		return errors.New("params is empty!")
	}
	getusers := make([]User, 0)
	result := db.Where("name=?", username).Find(&getusers)
	if result.RowsAffected <= 0 {
		user := User{Name: username, Password: password}
		if count := db.Create(&user).RowsAffected; count != 1 {
			beego.Error("add user fail!")
			return errors.New("add user fail!")
		}
		return nil
	}
	return errors.New("user is exist!")
}
