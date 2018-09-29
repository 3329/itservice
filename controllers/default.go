package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Layout = "layout.html"
	c.TplName = "index.html"

	/*
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "blogs/html_head.tpl"
	*/
}



func (c *MainController) ProcService() {
	c.Layout = "layout.html"
	c.TplName = "service.html"
}

func (c *MainController) ProcExcellent() {
	c.Layout = "layout.html"
	c.TplName = "excellent.html"
}

func (c *MainController) ProcNews() {
	c.Layout = "layout.html"
	c.TplName = "news.html"
}

func (c *MainController) ProcAboutus() {
	c.Layout = "layout.html"
	c.TplName = "aboutus.html"

	c.Data["CurTitle"] = "公司简介"
}

func (c *MainController) ProcJoinus() {
	c.Layout = "layout.html"
	c.TplName = "aboutus.html"

	c.Data["CurTitle"] = "加入我们"
}

func (c *MainController) ProcContactus() {
	c.Layout = "layout.html"
	c.TplName = "aboutus.html"

	c.Data["CurTitle"] = "联系我们"
}


// -------------------------YunWei Page start-------------------------------------------
func (c *MainController) ProcYunwei01() {
	c.Layout = "layout.html"
	c.TplName = "yunwei.html"

	c.Data["CurTitle"] = "应急服务"
	c.Data["TitleDesc"] = "本服务适合于用户电脑数量50台以上或对IT依赖性比较高或IT工作量比较集中的企业。 这种服务方式优点就是服务方式非常灵活。"
	c.Data["PlanDesc"] = "我们提供桌面支持、项目管理、项目实施等各类IT人员派驻服务，通过流程化的预防性体系和严格的工程师绩效考核体系，控制服务质量和安全保密等客户关注点,具体方式：" 
	
	var PlanItemsMap map[int]string 
	PlanItemsMap = map[int]string{}
	PlanItemsMap[0] = "每周5天（周一~周五），按照用户作息时间上门提供人员派驻服务" 
	PlanItemsMap[1] = "每周3天（如周一、三、五），按照用户作息时间上门提供人员派驻服务" 
	PlanItemsMap[2] = "每周2天（如周二、四），按照用户作息时间上门提供人员派驻服务"
	PlanItemsMap[3] = "每周1天（如周三），按照用户作息时间上门提供人员派驻服务"
	c.Data["PlanItemsMap"] = PlanItemsMap 
}

func (c *MainController) ProcPageZhuchang() {
	c.Layout = "layout.html"
	c.TplName = "yunwei.html"

	c.Data["CurTitle"] = "驻场服务"
	c.Data["TitleDesc"] = "本服务适合于用户电脑数量50台以上或对IT依赖性比较高或IT工作量比较集中的企业。 这种服务方式优点就是服务方式非常灵活。"
	c.Data["PlanDesc"] = "我们提供桌面支持、项目管理、项目实施等各类IT人员派驻服务，通过流程化的预防性体系和严格的工程师绩效考核体系，控制服务质量和安全保密等客户关注点,具体方式：" 
	
	var PlanItemsMap map[int]string 
	PlanItemsMap = map[int]string{}
	PlanItemsMap[0] = "每周5天（周一~周五），按照用户作息时间上门提供人员派驻服务" 
	PlanItemsMap[1] = "每周3天（如周一、三、五），按照用户作息时间上门提供人员派驻服务" 
	PlanItemsMap[2] = "每周2天（如周二、四），按照用户作息时间上门提供人员派驻服务"
	PlanItemsMap[3] = "每周1天（如周三），按照用户作息时间上门提供人员派驻服务"
	c.Data["PlanItemsMap"] = PlanItemsMap 
}

func (c *MainController) ProcPageYinji() {
	c.Layout = "layout.html"
	c.TplName = "yunwei.html"

	c.Data["CurTitle"] = "应急服务"
	c.Data["TitleDesc"] = "本服务适用于用户电脑台数量50台以内的企业。 这种服务方式优点是方便、灵活、随叫随到, 5分钟响应、2小时内到场服务。 专属工程师、拒绝陌生面孔，而且相对固定工程师驻场服务，应急响应服务成本相对比较低。"
	c.Data["PlanDesc"] = "具体方式：" 
	
	var PlanItemsMap map[int]string 
	PlanItemsMap = map[int]string{}
	PlanItemsMap[0] = "7天 x 24小时 x 2小时服务（随时可以提请故障上门，2小时到场，不分休息日和工作时间）" 
	PlanItemsMap[1] = "5天 x 8小时 x 2小时服务（每周5天，8小时工作时间内，2小时到场）" 
	PlanItemsMap[2] = "5天 x 8小时 x 4小时服务（每周5天，8小时工作时间内，4小时到场）"
	PlanItemsMap[3] = "5天 x 8小时 x 8小时服务（每周5天，故障提请隔日上门）"
	c.Data["PlanItemsMap"] = PlanItemsMap 
}

func (c *MainController) ProcPageYuanchen() {
	c.Layout = "layout.html"
	c.TplName = "yunwei.html"

	c.Data["CurTitle"] = "远程服务"
	c.Data["TitleDesc"] = "此服务适用于用户的电脑数量10台内或者 移动办公人员多，地域分散或需要支持的次数少，响应要求却很高的企业。"
	c.Data["PlanDesc"] = "具体方式：" 
	
	var PlanItemsMap map[int]string 
	PlanItemsMap = map[int]string{}
	PlanItemsMap[0] = "我们通过专业的软件（Teamviewer）为用户提供快速、安全、稳定的远程服务（5分钟内相应）。" 
	c.Data["PlanItemsMap"] = PlanItemsMap 
}
// -------------------------YunWei Page end-------------------------------------------



// -------------------------RuoDian Page start-------------------------------------------
func (c *MainController) ProcPageBuxian() {
	c.Layout = "layout.html"
	c.TplName = "ruodian.html"

	c.Data["CurTitle"] = "综合布线"

	var PlanItemsMap map[int]string 
	PlanItemsMap = map[int]string{}
	PlanItemsMap[0] = "我们通过专业的软件（Teamviewer）为用户提供快速、安全、稳定的远程服务（5分钟内相应）。" 
	c.Data["PlanItemsMap"] = PlanItemsMap 
}
// -------------------------RuoDian Page end-------------------------------------------