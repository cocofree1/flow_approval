package dao

import "time"

type FlowUser struct{
	Id         int     	 `orm:"column(id);auto;pk" json:"id"`
	UserName   string  	 `orm:"column(user_name);size(64)" json:"user_name" description:"用户名"`
	Telephone  string  	 `orm:"column(telephone);size(11)" json:"telephone" description:"电话号码"`
	Email      string  	 `orm:"column(email);size(64)" json:"email" description:"邮箱"`
	Password   string    `orm:"column(password);size(64)" json:"password" description:"密码"`
	DFlag      int       `orm:"column(dflag);size(11)" description:"删除标志"`
	CreatedAt  time.Time `orm:"column(created_at)" json:"created_at" description:"创建日期"`
	UpdatedAt  time.Time `orm:"column(updated_at);type(timestamp);auto_now" json:"updated_at" description:"最后创建日期"`
}
