package routers

import (
	"github.com/astaxie/beego"
	"github.com/coderminer/blog/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/editor", &controllers.EditorController{})
	beego.Router("/blog", &controllers.BlogController{})
	beego.Router("/blog/:id", &controllers.BlogController{}, "get:Detail")
}
