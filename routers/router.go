package routers

import (
	"github.com/astaxie/beego"
	"github.com/coderminer/blog/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
