package dao

import "time"

type NodeInstance struct{
	Id              int       `orm:"column(id);auto;pk" json:"id"`
	FlowInstanceId  int       `orm:"column(flow_instance_id);size(64)"`
	NodeId          int       `orm:"column(node_id);size(64)"`
	ApprovalInfo    string    `orm:"column(approval_info)" description:"审批信息"`
	ApprovalUser    string    `orm:"column(approval_user)" description:"审批人"`
	NodeStatus      int       `orm:"column(node_status)" description:"节点状态"`
	DFlag           int       `orm:"column(dflag);size(11)" description:"删除标志"`
	CreatedAt       time.Time `orm:"column(created_at)" json:"created_at" description:"创建日期"`
	UpdatedAt       time.Time `orm:"column(updated_at);type(timestamp);auto_now" json:"updated_at" description:"最后创建日期"`
}
