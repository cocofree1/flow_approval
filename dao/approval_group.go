package dao

import "time"

// 每个节点对应的审批组
type ApprovalGroup struct{
	Id           int       `orm:"column(id);auto;pk" json:"id"`
	GroupName    string    `orm:"column(group_name)"`
	UserIds      string    `orm:"column(user_ids)" description:"用户IDs"`
	Description  string    `orm:"column(description)"`
	DFlag        int       `orm:"column(dflag);size(11)" description:"删除标志"`
	CreatedAt    time.Time `orm:"column(created_at)" json:"created_at" description:"创建日期"`
	UpdatedAt    time.Time `orm:"column(updated_at);type(timestamp);auto_now" json:"updated_at" description:"最后创建日期"`
}