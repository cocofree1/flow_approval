package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	flowRouter := beego.NewNamespace("/flow",
		UserRouter(),
		)
	beego.AddNamespace(flowRouter)
}
