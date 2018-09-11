package routers

import (
	"itservice/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/aboutus", &controllers.MainController{}, "*:ProcAboutus")
	beego.Router("/joinus", &controllers.MainController{}, "*:ProcJoinus")
	beego.Router("/contactus", &controllers.MainController{}, "*:ProcContactus")
}
