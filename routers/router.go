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

	// ------------------------- YunWei start ---------------------------------
	beego.Router("/yunwei01", &controllers.MainController{}, "*:ProcYunwei01")
	beego.Router("/zhuchang", &controllers.MainController{}, "*:ProcPageZhuchang")
	beego.Router("/yinji", &controllers.MainController{}, "*:ProcPageYinji")
	beego.Router("/yuanchen", &controllers.MainController{}, "*:ProcPageYuanchen")
	// ------------------------- YunWei end ---------------------------------


	// ------------------------- RuoDian start ---------------------------------
	beego.Router("/buxian", &controllers.MainController{}, "*:ProcPageBuxian") 
	beego.Router("/jifangjianshe", &controllers.MainController{}, "*:ProcPageJifangjianshe")
	beego.Router("/menjinkaoqing", &controllers.MainController{}, "*:ProcPageMenjinkaoqing")
	beego.Router("/shipinjiankong", &controllers.MainController{}, "*:ProcPageShipinjiankong")
	beego.Router("/fangdaobaojing", &controllers.MainController{}, "*:ProcPageFangdaobaojing")
	beego.Router("/duomeitihuiyishi", &controllers.MainController{}, "*:ProcPageDuomeitihuiyishi")
	beego.Router("/yikakong", &controllers.MainController{}, "*:ProcPageYikakong")
	beego.Router("/pinjieping", &controllers.MainController{}, "*:ProcPagePinjieping")
	// ------------------------- RuoDian end ---------------------------------


	// ------------------------- XiTongJiChen start---------------------------------
	beego.Router("/wangluogaizhao", &controllers.MainController{}, "*:ProcPageWangluogaizhao")
	beego.Router("/buxian", &controllers.MainController{}, "*:ProcPageBuxian")
	beego.Router("/buxian", &controllers.MainController{}, "*:ProcPageBuxian")
	beego.Router("/buxian", &controllers.MainController{}, "*:ProcPageBuxian")
	beego.Router("/buxian", &controllers.MainController{}, "*:ProcPageBuxian")
	beego.Router("/buxian", &controllers.MainController{}, "*:ProcPageBuxian")
	beego.Router("/buxian", &controllers.MainController{}, "*:ProcPageBuxian")
	beego.Router("/buxian", &controllers.MainController{}, "*:ProcPageBuxian")
	beego.Router("/buxian", &controllers.MainController{}, "*:ProcPageBuxian")
	beego.Router("/buxian", &controllers.MainController{}, "*:ProcPageBuxian")
	beego.Router("/buxian", &controllers.MainController{}, "*:ProcPageBuxian")
	beego.Router("/buxian", &controllers.MainController{}, "*:ProcPageBuxian")
	beego.Router("/buxian", &controllers.MainController{}, "*:ProcPageBuxian")
	beego.Router("/buxian", &controllers.MainController{}, "*:ProcPageBuxian")
	beego.Router("/buxian", &controllers.MainController{}, "*:ProcPageBuxian") 	
	// ------------------------- XiTongJiChen end---------------------------------

}
