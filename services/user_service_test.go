package services

import (
	"testing"
	"fmt"
	"carp.cn/whale/kit"
	"carp.cn/whale/config"
	"carp.cn/whale/zaplogger"
	"carp.cn/whale/db"
)

func init(){
	config.LoadFromFile("../config.json")

	zaplogger.SetLogger("../logs", "test.go", "debug", true)


	db.InitDB()
	kit.Init()
}
func TestUserService_GetUserByID(t *testing.T) {
	user := DefaultUserService.GetUserByID(1)
	user.Nickname = "snlan"
	err := DefaultUserService.UpdateUser(user)
	DefaultUserService.GetUserByID(1)
	fmt.Println(err)
	fmt.Printf("%+v", user)
}
