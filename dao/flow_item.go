package dao

import "time"

// 流程
type FlowItem struct{
	Id           int       `orm:"column(id);auto;pk" json:"id"`
	FlowName     string    `orm:"column(flow_name);size(64)"`
	FlowType     int       `orm:"column(flow_type)"`
	Description  string    `orm:"column(description)"`
	DFlag        int       `orm:"column(dflag);size(11)" description:"删除标志"`
	CreatedAt    time.Time `orm:"column(created_at)" json:"created_at" description:"创建日期"`
	UpdatedAt    time.Time `orm:"column(updated_at);type(timestamp);auto_now" json:"updated_at" description:"最后创建日期"`
}