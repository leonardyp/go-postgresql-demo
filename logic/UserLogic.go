package logic

import (
	"errors"
	"go-postgresql-demo/dao"
	"go-postgresql-demo/models"
	"strings"
)

// @Description 使用用户名创建用户
func CreateUserLogic(user *models.User) error {
	if strings.TrimSpace(user.Name) == "" {
		return errors.New("empty user name is illegal")
	} else {
		user.Type = "user"
		return dao.CreateUserDao(user)
	}
}

// @Description 获取所有用户列表
func AllUserLogic(users *[]models.User) error {
	return dao.AllUserDao(users)
}
