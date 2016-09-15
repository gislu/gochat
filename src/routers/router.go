package routers

import (
	"controllers"
	"github.com/astaxie/beego"

)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/corp",&controllers.CorpController{})
	beego.Router("/enroll",&controllers.EnrollController{})
	beego.Router("/pub",&controllers.PubController{})
}