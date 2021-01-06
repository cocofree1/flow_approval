package dao

import "time"

// 流程实例
type FlowInstance struct{
	Id           int       `orm:"column(id);auto;pk" json:"id"`
	FlowId       int       `orm:"column(flow_id);size(64)"`
	FlowInfo     string    `orm:"column(flow_info)" description:"申请信息"`
	Applicant    string    `orm:"column(applicant)" description:"申请人"`
	ApplyReason  string    `orm:"column(apply_reason)" description:"申请理由"`
	CurrentNode  int       `orm:"column(current_node)"`
	FlowStatus   int       `orm:"column(flow_status)" description:"流程状态"`
	DFlag        int       `orm:"column(dflag);size(11)" description:"删除标志"`
	CreatedAt    time.Time `orm:"column(created_at)" json:"created_at" description:"创建日期"`
	UpdatedAt    time.Time `orm:"column(updated_at);type(timestamp);auto_now" json:"updated_at" description:"最后创建日期"`
}
