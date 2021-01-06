package controllers

import (
	"flow_approval/lib"
	"flow_approval/models"
	"flow_approval/request"
	"github.com/astaxie/beego"
)

type FlowController struct {
	beego.Controller
}

func (flow *FlowController) CreateFlow(){
	var r lib.ResultInfo
	var input request.CreateFlow
	r = lib.ParseRequestBody(&input,&flow.Controller)
	if r.IsOk(){
		r = models.CreateFlow(input)
	}
	flow.Data["json"] = r
	flow.ServeJSON()
}


func (flow *FlowController) DeleteFlow(){
	var r lib.ResultInfo
	r = lib.ParseRequestParams(&flow.Controller)
	if r.IsOk(){
		r = models.DeleteFlow(r.Data.(int))
	}
	flow.Data["json"] = r
	flow.ServeJSON()
}

func (flow *FlowController) UpdateFlow(){
	var r lib.ResultInfo
	var input request.UpdateFlow
	r = lib.ParseRequestBody(&input,&flow.Controller)
	if r.IsOk(){
		r = models.UpdateFlow(input)
	}
	flow.Data["json"] = r
	flow.ServeJSON()
}

func (flow *FlowController) GetAllFlows(){
	r := models.GetAllFlows()
	flow.Data["json"] = r
	flow.ServeJSON()
}

func (flow *FlowController) GetFlowById(){
	var r lib.ResultInfo
	r = lib.ParseRequestParams(&flow.Controller)
	if r.IsOk(){
		r = models.GetFlowById(r.Data.(int))
	}
	flow.Data["json"] = r
	flow.ServeJSON()
}

func (flow *FlowController) CreateFlowInstance(){
	var r lib.ResultInfo
	var input request.CreateFlowInstance
	r = lib.ParseRequestBody(&input,&flow.Controller)
	if r.IsOk(){
		r = models.CreateFlowInstance(input)
	}
	flow.Data["json"] = r
	flow.ServeJSON()
}

func (flow *FlowController) DeleteFlowInstance(){
	var r lib.ResultInfo
	r = lib.ParseRequestParams(&flow.Controller)
	if r.IsOk(){
		r = models.DeleteFlowInstance()
	}
	flow.Data["json"] = r
	flow.ServeJSON()
}

func (flow *FlowController) UpdateFlowInstance(){
	var r lib.ResultInfo
	var input request.UpdateFlowInstance
	r = lib.ParseRequestBody(&input,&flow.Controller)
	if r.IsOk(){
		r = models.UpdateFlowInstance(input)
	}
	flow.Data["json"] = r
	flow.ServeJSON()
}

func (flow *FlowController) GetAllFlowInstances(){
	r := models.GetAllFlowInstances()
	flow.Data["json"] = r
	flow.ServeJSON()
}

func (flow *FlowController) GetFlowInstanceById(){
	var r lib.ResultInfo
	r = lib.ParseRequestParams(&flow.Controller)
	if r.IsOk(){
		r = models.GetFlowInstanceById()
	}
	flow.Data["json"] = r
	flow.ServeJSON()
}

