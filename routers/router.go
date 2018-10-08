package routers

import (
	"itservice/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}) 

	beego.Router("/aboutus", &controllers.MainController{}, "*:ProcAboutus")
	beego.Router("/service", &controllers.MainController{}, "*:ProcService")
	beego.Router("/excellent", &controllers.MainController{}, "*:ProcExcellent")
	beego.Router("/news", &controllers.MainController{}, "*:ProcNews")

	// -------------------------YunWei---------------------------------
	beego.Router("/yunwei01", &controllers.MainController{}, "*:ProcYunwei01")
	beego.Router("/zhuchang", &controllers.MainController{}, "*:ProcPageZhuchang")
	beego.Router("/yinji", &controllers.MainController{}, "*:ProcPageYinji")
	beego.Router("/yuanchen", &controllers.MainController{}, "*:ProcPageYuanchen")

	// -------------------------RuoDian---------------------------------
	beego.Router("/buxian", &controllers.MainController{}, "*:ProcPageBuxian")
}
