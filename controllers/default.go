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
}



// -------------------------YunWei Page start-------------------------------------------
func (c *MainController) ProcYunwei01() {
	c.Layout = "layout.html"
	c.TplName = "yunwei.html"

	c.Data["CrtTitle"] = "应急服务"
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

	c.Data["CrtTitle"] = "驻场服务"
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

	c.Data["CrtTitle"] = "应急服务"
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

	c.Data["CrtTitle"] = "远程服务"
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

	c.Data["CrtTitle"] = "综合布线" 
	c.Data["TitleDesc"] = "综合布线系统是企业信息化运行的通道，其通畅性至关重要，我们会根据用户的实际需求，设计高性价比的布线系统。综合布线系统一个能够支持任何用户选择的话音、数据、图形图像应用的电信布线系统。系统应能支持话音、图形、图像、数据多媒体、安全监控、传感等各种信息的传输，支持UTP、光纤、STP、同轴电缆等各种传输载体，支持多用户多类型产品的应用，支持高速网络的应用。"

	var PlanItemsMap map[int]string 
	PlanItemsMap = map[int]string{}
	PlanItemsMap[0] = "综合布线系统是企业信息化运行的通道，其通畅性至关重要，我们会根据用户的实际需求，设计高性价比的布线系统。"  

	var PartsMap map[string]string 
	PartsMap = map[string]string{} 	
	PartsMap["工作区子系统"] = "目的是实现工作区终端设备与水平子系统之间的连接，由终端设备连接到信息插座的连接线缆所组成。由信息插座、插座盒、连接跳线和适配器组成。工作区子系统的设计主要考虑信息插座和适配器两个方面。"
	PartsMap["建筑群主干子系统"] = "提供外部建筑物与大楼内布线的连接点。EIA/TIA569标准规定了网络接口的物理规格，实现建筑群之间的连接"
	PartsMap["设备子系统"] = "EIA/TIA569标准规定了设备间的设备布线。它是布线系统最主要的管理区域，所有楼层的资料都由电缆或光纤电缆传送至此。通常，此系统安装在计算机系统、网络系统和程控机系统的主机房内。"
	PartsMap["垂直主干子系统"] = "它连接通讯室、设备间和入口设备，包括主干电缆、中间交换和主交接、机械终端和用于主干到主干交换的接插线或插头。主干布线要采用星形拓扑结构，接地应符合EIA/TIA607规定的要求。"
	PartsMap["管理子系统"] = "管理子系统由交连、互连配线架组成。管理点为连接其它子系统提供连接手段。交连和互连允许将通讯线路定位或重定位到建筑物的不同部分，以便能更容易地管理通信线路，使在移动终端设备时能方便地进行插拔。"
	PartsMap["水平支干线子系统"] = "目的是实现信息插座和管理子系统（跳线架）间的连接，将用户工作区引至管理子系统，并为用户提供一个符合国际标准，满足语音及高速数据传输要求的信息点出口。该子系统由一个工作区的信息插座开始，经水平布置到管理区的内侧配线架的线缆所组成。"

	var TroubleItems map[string]string 
	TroubleItems = map[string]string{} 	 
	TroubleItems["01"] = "机房像蛛网一样错综复杂的线缆难以理清头绪？" 
	TroubleItems["02"] = "某线路出问题了却无法定位故障点？" 
	TroubleItems["03"] = "网络设备连接混乱，一旦出故障整个网络全部陷入瘫痪。" 
	TroubleItems["04"] = "综合布线点位不多，只要能满足办公，随便做下就可以？" 

	c.Data["PlanItemsMap"] = PlanItemsMap 
	c.Data["PartsMap"] = PartsMap 
	c.Data["TroubleTitle"] = "机房工程是否一直被以下问题困扰？" 
	c.Data["TroubleItems"] = TroubleItems 
} 

func (c *MainController) ProcPageJifangjianshe() {
	c.Layout = "layout.html" 
	c.TplName = "ruodian.html"

	c.Data["CrtTitle"] = "机房建设" 
	var TroubleItems map[string]string 
	TroubleItems = map[string]string{} 	
	TroubleItems["01"] = "机房改造需要规划多大的面积才适合？" 
	TroubleItems["02"] = "需要给机房留多少配电才能够满足机房设备用电需求？" 
	TroubleItems["03"] = "不知道机房设备放在那里合适，担心楼板承重不安全？" 
	TroubleItems["04"] = "不清楚机房按照什么标准建设合理？"
	TroubleItems["05"] = "机房建设需要做哪些系统工程？" 
	TroubleItems["06"] = "不知道如何规划设计机房？" 
	TroubleItems["07"] = "不知道机房建设预算？" 
	TroubleItems["08"] = "找不到专业的机房建设公司？"

	var PlanItemsMap map[int]string 
	PlanItemsMap = map[int]string{}
	PlanItemsMap[0] = "我们通过专业的软件（Teamviewer）为用户提供快速、安全、稳定的远程服务（5分钟内相应）。" 
	c.Data["PlanItemsMap"] = PlanItemsMap 
	c.Data["TroubleTitle"] = "机房工程是否一直被以下问题困扰？" 
	c.Data["TroubleItems"] = TroubleItems 
}

func (c *MainController) ProcPageMenjinkaoqing() { 
	c.Layout = "layout.html" 
	c.TplName = "ruodian.1.html"

	var CompriseItems map[string]string 
	CompriseItems = map[string]string{}
	CompriseItems["硬件功能"] = "门禁考勤系统门禁系统按联接方式可分为多门联网门禁系统和单门单机门禁系统两种；按数据读取方式可分为感应式、密码式、感应+密码式；刷卡式、刷卡+感应式、刷卡+密码式；指纹、指纹+密码（联机和脱机型）等多种方式。经授权用户可利用使用卡开启各房间门的电锁。用户权限可由系统管理员设定。当系统出现故障时管理人员可利用后备卡，开启通道门。系统可及时将各部分的工作状态向监控中心或用户指定地点报告。当系统中某个组成部分出现故障时，可及时将故障情况向监控中心或用户指定地点报告。可利用读卡器对公司内的工作人员进行考勤管理。大门利用密码+使用卡才能进出。管理人员可根据需求随意增加或删除使用卡。前台工作人员可与访客利用对讲系统通话并通过控制电脑在远端开启/关闭各通道门。各大门安装玻璃破碎开关，在紧急情况时，可按下玻璃破碎开关直接开门。" 
	

	c.Data["CrtTitle"] = "门禁考勤" 
	c.Data["TitleDesc"] = "门禁考勤系统[出入口自动控制系统]是企业弱电智能化的基础组建，它具有对人员进出、授权、查询、统计和防盗报警保安等, 多种功能，还可作为人事管理、考勤管理，可于任何机电设备产品及控制系统联动，既方便内部人员或用户的自由出入，又杜绝外来人员随意进出，提高安全防范能力。它将结束钥匙时代，成为现代社会进出管理的新潮流。"
	c.Data["CompriseItems"] = CompriseItems 
	
}

// -------------------------RuoDian Page end-------------------------------------------



