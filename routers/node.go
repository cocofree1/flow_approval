package routers

import (
	"flow_approval/controllers"
	"github.com/astaxie/beego"
)

func NodeRouter() beego.LinkNamespace{
	userRouter := beego.NSNamespace("/node",
		beego.NSRouter("/",&controllers.NodeController{},"post:CreateNode"),
		beego.NSRouter("/:id",&controllers.NodeController{}, "delete:DeleteNode"),
		beego.NSRouter("/",&controllers.NodeController{},"put:UpdateNode"),
		beego.NSRouter("/", &controllers.NodeController{},"get:GetAllNodes"),
		beego.NSRouter("/:id", &controllers.NodeController{},"get:GetNodeById"),
	)
	return userRouter
}
