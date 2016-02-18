package dao

import (
	"gopkg.in/pg.v3"
	"tantan/logger"
	"tantan/models"
)

// @Description 使用结构体创建用户
func CreateUserDao(user *models.User) error {
	_, err := models.DB().ExecOne(`
	      INSERT INTO users VALUES (?id,?name,?type)
	  `, *user)
	if err != nil {
		logger.Error("%v", err)
		logger.ErrorStd("%v", err)
	}
	return err
}

// @Description 获取用户最大id
func GetMaxUserId() (int64, error) {
	var id int64
	_, err := models.DB().QueryOne(pg.LoadInto(&id), `select max(id) from users`)
	if err != nil {
		logger.Error("%v", err)
		logger.ErrorStd("%v", err)
	}
	return id, err
}

// @Description 获取所有用户列表
func AllUserDao(users *[]models.User) error {
	_, err := models.DB().Query(users, `
		SELECT * FROM users
	`)
	if err != nil {
		logger.Error("%v", err)
		logger.ErrorStd("%v", err)
	}
	return err
}

// @Description 根据用户id获取用户信息
func GetUserById(userId int64) (models.User, error) {
	var user models.User
	_, err := models.DB().QueryOne(&user, `
		SELECT * FROM users
		WHERE id = ?
	`, userId)
	if err != nil {
		logger.Error("%v", err)
		logger.ErrorStd("%v", err)
	}
	return user, err
}

// @Description 更新用户喜好
func UpdateUserLikes(user models.User) error {
	_, err := models.DB().Exec("update users set likes=? where id=?", user.Likes, user.Id)
	if err != nil {
		logger.Error("%v", err)
		logger.ErrorStd("%v", err)
	}
	return err
}

// @Description 根据用户id切片过滤后的用户
func UsersFilterByIds(ids []int64) ([]models.User, error) {
	var users = []models.User{}
	_, err := models.DB().Query(&users, `SELECT * FROM users WHERE id not IN (?)`, pg.Ints(ids))
	if err != nil {
		logger.Error("%v", err)
		logger.ErrorStd("%v", err)
	}
	return users, err
}
