package routers

import (
	"flow_approval/controllers"
	"github.com/astaxie/beego"
)

func FlowRouter() beego.LinkNamespace{
	userRouter := beego.NSNamespace("/process",
		beego.NSRouter("/",&controllers.FlowController{},"post:CreateFlow"),
		beego.NSRouter("/:id",&controllers.FlowController{}, "delete:DeleteFlow"),
		beego.NSRouter("/",&controllers.FlowController{},"put:UpdateFlow"),
		beego.NSRouter("/", &controllers.FlowController{},"get:GetAllFlows"),
		beego.NSRouter("/:id", &controllers.FlowController{},"get:GetFlowById"),

		// instance
		//---------------------------
		beego.NSRouter("/instance",&controllers.FlowController{},"post:CreateFlowInstance"),
		beego.NSRouter("/instance/:id",&controllers.FlowController{}, "delete:DeleteFlowInstance"),
		beego.NSRouter("/instance",&controllers.FlowController{},"put:UpdateFlowInstance"),
		beego.NSRouter("/instance", &controllers.FlowController{},"get:GetAllFlowInstances"),
		beego.NSRouter("/instance/:id", &controllers.FlowController{},"get:GetFlowInstanceById"),
	)
	return userRouter
}
