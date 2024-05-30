package utils

import (
	"log"
	"questionplatform/global"

	"golang.org/x/crypto/bcrypt"
)

func GetPwd(password string) (hashpass []byte, err error) {
	pass := password + global.Config.Salt
	hashpass, err = bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return
}

func ComparePwd(hashpwd, password string) (flag bool) {
	pass := password + global.Config.Salt
	err := bcrypt.CompareHashAndPassword([]byte(hashpwd), []byte(pass))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
