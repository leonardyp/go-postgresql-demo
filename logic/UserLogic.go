package logic

import (
	"errors"
	"strings"
	"tantan/dao"
	"tantan/logger"
	"tantan/models"
)

// @Description 使用用户名创建用户
func CreateUserLogic(user *models.User) error {
	if strings.TrimSpace(user.Name) == "" {
		return errors.New("empty user name is illegal")
	} else {
		user.Type = "user"
		maxUserId, err := dao.GetMaxUserId()
		if err != nil {
			logger.Error("%v", err)
			logger.ErrorStd("%v", err)
			return err
		}
		logger.Debug("最大用户id:%v", maxUserId)
		logger.DebugStd("最大用户id:%v", maxUserId)
		user.Id = maxUserId + 1
		return dao.CreateUserDao(user)
	}
}

// @Description 获取所有用户列表
func AllUserLogic(users *[]models.User) error {
	return dao.AllUserDao(users)
}
