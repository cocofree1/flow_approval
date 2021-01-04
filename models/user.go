package models

import (
	"flow_approval/dao"
	"flow_approval/lib"
	"flow_approval/request"
	"time"
)

// UserRegister 用户注册
func UserRegister(userInfo request.UserRegister)(r lib.ResultInfo){
	user := &dao.FlowUser{
		UserName:   userInfo.UserName,
		Telephone:  userInfo.Telephone,
		Email:      userInfo.Email,
		Password:   userInfo.Password,
		CreatedAt:  time.Now(),
	}

	_,err :=lib.DbObject.Insert(user)
	if err != nil{
		r = lib.ErrorResulInfo(lib.ERR_INSERT_TABLE,"插入User表失败",nil,err)
	}
	return
}

func UserLogin(loginInfo request.UserLogin)(r lib.ResultInfo){
	//验证
	//查询
	qs := lib.DbObject.QueryTable(&dao.FlowUser{})
	var user dao.FlowUser
	err := qs.Filter("user_name", loginInfo.Regex).
		Filter("password", loginInfo.Password).
		One(&user)
	if err != nil{
		r = lib.ErrorResulInfo(lib.ERR_SEARCH_TABLE,"查询User表失败",nil,err)
	}
	r.Data = user
	return
}