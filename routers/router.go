package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	flowRouter := beego.NewNamespace("/flow",
		UserRouter(),
		FlowRouter(),
		)
	beego.AddNamespace(flowRouter)
}
