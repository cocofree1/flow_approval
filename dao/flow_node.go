package dao

import "time"

// 流程节点
type FlowNode struct{
	Id           int       `orm:"column(id);auto;pk" json:"id"`
	FlowId       int       `orm:"column(flow_id);size(64)"`
	NodeName     string    `orm:"column(node_name)"`
	FlowPos      int       `orm:"column(flow_pos)" description:"节点位置"`
	GroupId      int       `orm:"column(group_id)" description:"审批组id"`
	Description  string    `orm:"column(description)"`
	DFlag        int       `orm:"column(dflag);size(11)" description:"删除标志"`
	CreatedAt    time.Time `orm:"column(created_at)" json:"created_at" description:"创建日期"`
	UpdatedAt    time.Time `orm:"column(updated_at);type(timestamp);auto_now" json:"updated_at" description:"最后创建日期"`
}