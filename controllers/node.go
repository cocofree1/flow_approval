package controllers

import (
	"flow_approval/lib"
	"flow_approval/models"
	"flow_approval/request"
	"github.com/astaxie/beego"
)

type NodeController struct {
	beego.Controller
}

func (node *NodeController) CreateNode(){
	var r lib.ResultInfo
	var input request.CreateNode
	r = lib.ParseRequestBody(&input,&node.Controller)
	if r.IsOk(){
		r = models.CreateNode(input)
	}
	node.Data["json"] = r
	node.ServeJSON()
}

func (node* NodeController) DeleteNode(){
	var r lib.ResultInfo
	r = lib.ParseRequestParams(&node.Controller)
	if r.IsOk(){
		r = models.DeleteNode()
	}
	node.Data["json"] = r
	node.ServeJSON()
}

func (node* NodeController) UpdateNode(){
	var r lib.ResultInfo
	var input request.UpdateNode
	r = lib.ParseRequestBody(&input,&node.Controller)
	if r.IsOk(){
		r = models.UpdateNode(input)
	}
	node.Data["json"] = r
	node.ServeJSON()
}

func (node* NodeController) GetAllNodes(){
	r := models.GetAllNodes()
	node.Data["json"] = r
	node.ServeJSON()
}

func (node* NodeController) GetNodeById(){
	var r lib.ResultInfo
	r = lib.ParseRequestParams(&node.Controller)
	if r.IsOk(){
		r = models.GetNodeById()
	}
	node.Data["json"] = r
	node.ServeJSON()
}
