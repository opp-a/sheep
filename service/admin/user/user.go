package user

import (
	"errors"
	"sheep/models"

	"github.com/astaxie/beego"
)

func Authenticate(user models.User) error {
	err, myuser := models.GetUser(user.Name)
	if err != nil {
		beego.Error(err)
		return err
	}
	if myuser == nil {
		beego.Error("user ", user.Name, " not exist!")
		return errors.New("user " + user.Name + " not exist!")
	}
	if user.Password != myuser.Password {
		beego.Error("password incorrect!")
		return errors.New("password incorrect!")
	}
	return nil
}
