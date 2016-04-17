package routers

import (
	"gochatserver/controllers"
	"github.com/astaxie/beego"

)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/test",&controllers.TestController{})
//	beego.Router("/test",&controllers.identifyController{})

}