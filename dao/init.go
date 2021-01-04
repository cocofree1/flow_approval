package dao

import "github.com/astaxie/beego/orm"

func init(){
	orm.RegisterModel(
		new(FlowUser),
	)
}