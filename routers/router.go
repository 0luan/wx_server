package routers

import (
	"proj/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/api", &controllers.ListController{})
}
