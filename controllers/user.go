package controllers

import (
	"flow_approval/lib"
	"flow_approval/models"
	"flow_approval/request"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func(user *UserController)UserLogin(){
	var r lib.ResultInfo
	var input request.UserLogin
	r = lib.ParseRequestBody(&input,&user.Controller)
	if r.IsOk(){
		r = models.UserLogin(input)
	}
	user.Data["json"] = r
	user.ServeJSON()
}

func(user *UserController) UserRegister(){
	var r lib.ResultInfo
	var input request.UserRegister
	r = lib.ParseRequestBody(&input,&user.Controller)
	if r.IsOk(){
		r = models.UserRegister(input)
	}
	user.Data["json"] = r
	user.ServeJSON()
}