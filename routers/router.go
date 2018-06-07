package routers

import (
	"proj/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/api/imgupload", &controllers.ImgUploadController{})
    beego.Router("/api", &controllers.MainController{})
}
