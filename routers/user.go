package routers

import (
	"flow_approval/controllers"
	"github.com/astaxie/beego"
)

func UserRouter() beego.LinkNamespace{
	userRouter := beego.NSNamespace("/user",
		beego.NSRouter("/login",&controllers.UserController{},"post:UserLogin"),
		beego.NSRouter("/register",&controllers.UserController{},"post:UserRegister"),
		)
	return userRouter
}