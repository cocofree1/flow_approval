package models

import (
	"flow_approval/dao"
	"flow_approval/lib"
	"flow_approval/request"
	"time"
)

func CreateFlow(createInfo request.CreateFlow)(r lib.ResultInfo){
	now := time.Now()
	flowItem := &dao.FlowItem{
		FlowName:     createInfo.FlowName,
		FlowType:     createInfo.FlowType,
		Description:  createInfo.Description,
		CreatedAt:    now,
	}
	_,err :=lib.DbObject.Insert(flowItem)
	if err != nil{
		r = lib.ErrorResulInfo(lib.ERR_INSERT_TABLE,"插入FlowItem表失败",nil,err)
		return
	}
	r.Data = "插入表成功"
	return
}

func UpdateFlow(updateInfo request.UpdateFlow)(r lib.ResultInfo){
	flowItem := &dao.FlowItem{
		Id:           updateInfo.FlowId,
		FlowName:     updateInfo.FlowName,
		FlowType:     updateInfo.FlowType,
		Description:  updateInfo.Description,
	}
	_,err := lib.DbObject.Update(flowItem,"FlowName","FlowType","Description")
	if err != nil{
		r = lib.ErrorResulInfo(lib.ERR_UPDATE_TABLE,"更新FlowItem表失败",nil,err)
		return
	}
	r.Data = "更新表成功"
	return
}

func DeleteFlow(id int)(r lib.ResultInfo){
	_,err := lib.DbObject.Delete(&dao.FlowItem{Id: id})
	if err != nil{
		r = lib.ErrorResulInfo(lib.ERR_DELETE_TABLE,"删除FlowItem表失败",nil,err)
		return
	}
	r.Data = "删除表失败"
	return
}

func GetAllFlows()(r lib.ResultInfo){
	var flowItems []dao.FlowItem
	_,err := lib.DbObject.QueryTable(&dao.FlowItem{}).All(&flowItems)
	if err != nil{
		r = lib.ErrorResulInfo(lib.ERR_SEARCH_TABLE,"查询FlowItem表失败",nil,err)
		return
	}
	r.Data = flowItems
	return
}

func GetFlowById(id int)(r lib.ResultInfo){
	flowItem := &dao.FlowItem{Id:id}
	err := lib.DbObject.Read(flowItem)
	if err != nil{
		r = lib.ErrorResulInfo(lib.ERR_SEARCH_TABLE,"查询FlowItem表失败",nil,err)
		return
	}
	r.Data = flowItem
	return
}