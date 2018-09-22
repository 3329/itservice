package routers

import (
	"itservice/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/aboutus", &controllers.MainController{}, "*:ProcAboutus")
	beego.Router("/service", &controllers.MainController{}, "*:ProcService")
	beego.Router("/joinus", &controllers.MainController{}, "*:ProcJoinus")
	beego.Router("/contactus", &controllers.MainController{}, "*:ProcContactus")

	beego.Router("/yunwei01", &controllers.MainController{}, "*:ProcYunwei01")
	beego.Router("/zhuchang", &controllers.MainController{}, "*:ProcPageZhuchang")
	beego.Router("/yinji", &controllers.MainController{}, "*:ProcPageYinji")
	beego.Router("/yuanchen", &controllers.MainController{}, "*:ProcPageYuanchen")
}
