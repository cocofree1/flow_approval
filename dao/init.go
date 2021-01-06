package dao

import "github.com/astaxie/beego/orm"

func init(){
	orm.RegisterModel(
		new(FlowUser),
		new(ApprovalGroup),
		new(FlowInstance),
		new(FlowItem),
		new(FlowNode),
		new(NodeInstance),
	)
}