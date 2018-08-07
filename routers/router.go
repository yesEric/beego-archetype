package routers

import (
	"beego-archetype/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{})
	beego.AutoRouter(&controllers.RoleController{})
}
