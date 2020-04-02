package routers

import (
	"github.com/astaxie/beego"
	"github.com/hetianyi/gift/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
