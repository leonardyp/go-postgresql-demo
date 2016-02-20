package logic

import (
	"errors"
	"go-postgresql-demo/dao"
	"go-postgresql-demo/logger"
	"go-postgresql-demo/models"
	"go-postgresql-demo/utils"
)

var ERR_ILLEGAL_STATE = errors.New("illegal state")

// @Description 关系存在不一致则更新,不存在则添加
func RelationshipLogic(userId, otherUserId int64, state string) (models.Relationship, error) {
	var relationship = models.Relationship{}

	if state != "liked" && state != "disliked" {
		return relationship, ERR_ILLEGAL_STATE
	}
	user, err := dao.GetUserById(userId)
	if err != nil {
		return relationship, err
	}
	otherUser, err := dao.GetUserById(otherUserId)
	logger.DebugStd("%#v", user)
	if user.Likes == nil || len(user.Likes) == 0 {
		if state == "liked" {
			user.Likes = append(user.Likes, otherUserId)
		}
		if otherUser.Likes != nil && utils.InInts(userId, otherUser.Likes) {
			state = "match"
		}
	} else {
		if state == "disliked" {
			user.Likes = utils.RemoveIntFromInts(otherUserId, user.Likes)
		} else {
			if otherUser.Likes != nil && utils.InInts(userId, otherUser.Likes) {
				state = "match"
			}
			if !utils.InInts(otherUserId, user.Likes) {
				user.Likes = append(user.Likes, otherUserId)
			}
		}
	}

	relationship.Id = otherUserId
	relationship.State = state
	relationship.Type = "relationship"

	return relationship, dao.UpdateUserLikes(user)
}

// @Description 用户所有关系
func UserAllRelationshipLogic(userId int64) ([]models.Relationship, error) {
	user, err := dao.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	var state string
	relationships := []models.Relationship{}
	if user.Likes == nil {
		user.Likes = []int64{}
	}
	for _, v := range user.Likes {
		u, err := dao.GetUserById(v)
		if err != nil {
			logger.Error("%v", err)
			logger.ErrorStd("%v", err)
			continue
		}
		if utils.InInts(userId, u.Likes) {
			state = "match"
		} else {
			state = "liked"
		}
		relationships = append(relationships, models.Relationship{
			Id:    v,
			State: state,
			Type:  "relationship",
		})
	}

	users, _ := dao.UsersFilterByIds(append(user.Likes, int64(userId)))
	for _, v := range users {
		relationships = append(relationships, models.Relationship{
			Id:    v.Id,
			State: "disliked",
			Type:  "relationship",
		})
	}
	return relationships, err
}
