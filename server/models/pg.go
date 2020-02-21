package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func init() {
	var err error
	for trytimes := 3; trytimes > 0; trytimes-- {
		beego.Info("try connect to sql")
		// db, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=root dbname=sheep sslmode=disable password=sheeppwd")
		db, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=root dbname=sheep password=sheeppwd sslmode=disable")
		if err != nil {
			beego.Error(err)
			time.Sleep(time.Second)
			continue
		}
		break
	}
	if err != nil {
		beego.Error(err)
		panic("connect db fail! msg:" + err.Error())
	}
	db.DB().SetMaxIdleConns(30)
	db.DB().SetMaxOpenConns(60)
	db.SingularTable(true)

	db.AutoMigrate(&User{})
	db.AutoMigrate(&File{})
	db.AutoMigrate(&Shop{})

	// tshop := Shop{
	// 	Name: "iam shop",
	// 	Icons: []File{
	// 		{Name: "file1", Content: "xxx1"},
	// 		{Name: "file2", Content: "xxxxxx1"},
	// 	},
	// }
	// AddShop([]Shop{tshop})

	InitUsers()
}
