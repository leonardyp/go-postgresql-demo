package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"tantan/logger"
	"tantan/logic"
	"tantan/models"
)

// Operations about user
type UserController struct {
	beego.Controller
}

// @Title create
// @Description 使用用户名创建用户,并返回新创建的用户
// @Param	body	body 	models.User	 true  "The user content"
// @Success 200
// @Failure 500 422
// @router / [post]
func (this *UserController) CreateUserAction() {
	var user models.User
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &user); err != nil {
		logger.Error("%v", err)
		logger.ErrorStd("%v", err)
		this.Ctx.Output.SetStatus(422)
		this.Data["json"] = map[string]string{"err": err.Error()}
		this.ServeJSON()
		return
	}
	logger.Debug("request:...%#v", user)
	logger.DebugStd("request:...%#v", user)
	if err := logic.CreateUserLogic(&user); err != nil {
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = map[string]string{"err": err.Error()}
		this.ServeJSON()
		return
	}
	this.Data["json"] = user
	this.ServeJSON()
}

// @Title AllUserAction
// @Description 获取所有用户
// @Success 200 用户列表
// @Failure 500
// @router / [get]
func (this *UserController) AllUserAction() {
	users := &[]models.User{}
	if err := logic.AllUserLogic(users); err != nil {
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = map[string]string{"err": err.Error()}
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = users
		this.ServeJSON()
	}
}

// @Title create or update
// @Description create or update
// @Param user_id  		   path	    int  true	"用户id"
// @Param other_user_id	   path 	int	 true	"其他用户id"
// @Param body 			   body     body true   "state"
// @Success 200 success
// @Failure 500 422
// @router /:user_id/relationships/:other_user_id [put]
func (this *UserController) RelationshipAction() {
	state := struct {
		State string
	}{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &state); err != nil {
		logger.Error("%v", err)
		logger.ErrorStd("%v", err)
		this.Ctx.Output.SetStatus(422)
		this.Data["json"] = map[string]string{"err": err.Error()}
		this.ServeJSON()
		return
	}

	if userId, err := strconv.Atoi(this.Ctx.Input.Param(":user_id")); err != nil {
		logger.Error("%v", err)
		logger.ErrorStd("%v", err)
		this.Ctx.Output.SetStatus(422)
		this.Data["json"] = map[string]string{"err": err.Error()}
		this.ServeJSON()
		return
	} else if otherUserId, err := strconv.Atoi(this.Ctx.Input.Param(":other_user_id")); err != nil {
		logger.Error("%v", err)
		logger.ErrorStd("%v", err)
		this.Ctx.Output.SetStatus(422)
		this.Data["json"] = map[string]string{"err": err.Error()}
		this.ServeJSON()
		return
	} else {
		logger.Debug("req:...userId:%v,otherUserId:%v,body:%#v", userId, otherUserId, state)
		logger.DebugStd("req:...userId:%v,otherUserId:%v,body:%#v", userId, otherUserId, state)
		relationship, err := logic.RelationshipLogic(int64(userId), int64(otherUserId), state.State)
		if err != nil {
			logger.Error("%v", err)
			logger.ErrorStd("%v", err)
			this.Ctx.Output.SetStatus(500)
			this.Data["json"] = map[string]string{"err": err.Error()}
			this.ServeJSON()
			return
		}
		this.Data["json"] = relationship
		this.ServeJSON()
		return
	}
}

// @Title user all relationships
// @Description user all relationships
// @Param	user_id	path 	int	 true  "用户id"
// @Success 200 success
// @Failure 500 422
// @router /:user_id/relationships [get]
func (this *UserController) UserAllRelationshipsAction() {
	if userId, err := strconv.Atoi(this.Ctx.Input.Param(":user_id")); err != nil {
		logger.Error("req:...userId:%v", err)
		logger.ErrorStd("req:...userId:%v", err)
		this.Ctx.Output.SetStatus(422)
		this.Data["json"] = map[string]string{"err": err.Error()}
		this.ServeJSON()

	} else {
		logger.Debug("req:...userId:%v", userId)
		logger.DebugStd("req:...userId:%v", userId)
		relationships, err := logic.UserAllRelationshipLogic(int64(userId))
		if err != nil {
			logger.Error("%v", err)
			logger.ErrorStd("%v", err)
			this.Ctx.Output.SetStatus(500)
			this.Data["json"] = map[string]string{"err": err.Error()}
			this.ServeJSON()
			return
		}
		this.Data["json"] = relationships
		this.ServeJSON()
		return
	}
}
