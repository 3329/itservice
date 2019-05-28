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
	c.TplName = "ruodian1.html"

	var CompriseItems, FunctionItems map[string]string 
	CompriseItems = map[string]string{} 
	FunctionItems = map[string]string{} 
	CompriseItems["硬件介绍"] = "门禁考勤系统门禁系统按联接方式可分为多门联网门禁系统和单门单机门禁系统两种；按数据读取方式可分为感应式、密码式、感应+密码式；刷卡式、刷卡+感应式、刷卡+密码式；指纹、指纹+密码（联机和脱机型）等多种方式。经授权用户可利用使用卡开启各房间门的电锁。用户权限可由系统管理员设定。当系统出现故障时管理人员可利用后备卡，开启通道门。系统可及时将各部分的工作状态向监控中心或用户指定地点报告。当系统中某个组成部分出现故障时，可及时将故障情况向监控中心或用户指定地点报告。可利用读卡器对公司内的工作人员进行考勤管理。大门利用密码+使用卡才能进出。管理人员可根据需求随意增加或删除使用卡。前台工作人员可与访客利用对讲系统通话并通过控制电脑在远端开启/关闭各通道门。各大门安装玻璃破碎开关，在紧急情况时，可按下玻璃破碎开关直接开门。" 
	
	FunctionItems["实时出勤查询"] = "在任何情况下，不进行任何考勤设置，都能实时查询各员工的刷卡记录（包括刷卡人、刷卡时间、刷卡性质、窗口号）及未打员工，也可进行按条件查询部分或全体人员的刷卡记录（如：输入工号、日期及时间），便可查询该员工此时间段内的打卡明细。" 
	FunctionItems["基本出勤处理"] = "根据事先的定义的考勤设置能对各种出勤数据进行处理，自动判断迟到、早退、缺勤的人员及有效打卡时间内的打卡明细情况。如：可单独查询迟到、早退情况、缺勤情况及有效打卡时间的打卡明细，不正常的打卡数据等，在这些查询功能中都可按条件进行查询；迟到早退查询功能中还可输入相应的分钟数，以便查询大于此时间数的人员。"
	FunctionItems["多种的考勤统计"] = "能实时统计员工的上下班、加班、迟到、早退、请假、缺勤等相关出勤信息（上面记录着员工不正常出勤的次数及分钟数）形成一张综合性汇总报表，也可按条件（日期、工号、部门等）进行统计工作。"
	FunctionItems["自动扣款统计"] = "根据出勤情况及考勤设置，能自动统计所有员工某时间范围内的出勤应扣款额，方便了薪资计算。加班自动统计：根据员工加班打卡时间及企业安排加班时间，自动统计员工的实际加班及安排加班时间内的加班时间，灵活适应各种企业的加班时间统计。" 
	FunctionItems["异常事项处理"] = "对因公事耽误打卡的人员可进行补卡处理，对请假员工可进行请假处理，对某人员、某班次的时间变动可进行工时调整处理，月底统计进将自动计算请假天数。灵活的班次设置：根据企业实际情况可灵活设置班次数目，各班次标准上下班时间，轮班、换班等设置工作。灵活的打卡限制：可灵活设置各班次的上下班有效打卡时间，杜绝员工随意打卡，使管理更方便、合理，员工有组织、有纪律。" 
	FunctionItems["灵活的考勤设置"] = "可根据企业内部实际情况任意设置考核的字段数（如：迟到、严重迟到、早退、严重早退、缺勤等），各字段对应考核时间范围及考核对象；并可对各字段、各班次（管理人员、工人）进行不同处罚金额。自动判断上下班卡：系统根据设置情况自动判断员工的打卡数据是上班卡还是下班卡，无须人为干预。"
	
	c.Data["CrtTitle"] = "门禁考勤" 
	c.Data["TitleDesc"] = "门禁考勤系统[出入口自动控制系统]是企业弱电智能化的基础组建，它具有对人员进出、授权、查询、统计和防盗报警保安等, 多种功能，还可作为人事管理、考勤管理，可于任何机电设备产品及控制系统联动，既方便内部人员或用户的自由出入，又杜绝外来人员随意进出，提高安全防范能力。它将结束钥匙时代，成为现代社会进出管理的新潮流。"
	c.Data["CompriseItems"] = CompriseItems 
	c.Data["FunctionItems"] = FunctionItems 
}

func (c *MainController) ProcPageShipinjiankong() { 
	c.Layout = "layout.html" 
	c.TplName = "ruodian1.html"

	var CompriseItems, FunctionItems map[string]string 
	CompriseItems = map[string]string{} 
	FunctionItems = map[string]string{} 
	CompriseItems["系统组成介绍"] = "视频监控系统包括前端摄像机、传输线缆、视频监控平台。摄像机可分为网络数字摄像机和模拟摄像机，可作为前端视频图像信号的采集。完整的视频监控系统是由摄像、传输、控制、显示、记录登记5大部分组成。摄像机通过网络线缆或同轴视频电缆将视频图像传输到控制主机，控制主机再将视频信号分配到各监视器及录像设备，同时可将需要传输的语音信号同步录入到录像机内。 通过控制主机，操作人员可发出指令，对云台的上、下、左、右的动作进行控制及对镜头进行调焦变倍的操作，并可通过视频矩阵实现在多路摄像机的切换。利用特殊的录像处理模式，可对图像进行录入、回放、调出及储存等操作。" 
	
	FunctionItems["远程访问"] = "传统的视频监控一般是在小范围内进行，而用户普遍要求访问地点不受地域限制，能随时随地访问被监控地点。多人同时访问同一个监控点：传统上，一个监控点一般是被一个监控中心（用户）所访问。同一个监控点很可能会同时被多个用户所访问，并且这些用户之间可能毫无关系。用户访问的复杂化将要求系统强化对访问权限的管理。" 
	FunctionItems["监控点趋向分散，同时监控趋向集中"] = "属于同一用户的监控点越来越分散，不受地域所限。而对这些分散的监控点，需要集中的管理与控制。要求监控系统具有开放性和扩展性：同一系统应当支持多种不同类型的监控设备，用户数、被监控点的数量可以方便地增减。"
	FunctionItems["海量数据存储"] = "网络化使得传统的本地录像功能可以转移到远程服务器上来实现，使得海量数据存储成为可能。同时，也要求系统具备更强的存储、检索和备份等功能。"
	FunctionItems["信息安全"] = "系统复杂化，用户的多元化，加上视频监控本身的业务特点必然要求对系统对信息安全提供有力的保证。" 
	FunctionItems["智能视频监控"] = "未来的视频监控系统将不仅仅局限于被动地提供视频画面，更要求系统本身有足够的智能，能够识别不同的物体，发现监控画面中的异常情况，以快速和适合的方式发出警报和提供有用信息，从而更加有效地协助安全人员处理危机。" 
	FunctionItems["降低误报和漏报现象"] = "成为应对袭击和处理突发事件的有力辅助工具。智能视频监控还可以应用在交通管理、客户行为分析、客户服务等多种非安全相关的场景，以提高用户的投资回报。"
	
	c.Data["CrtTitle"] = "视频监控" 
	c.Data["TitleDesc"] = "视频监控是安全防范系统的重要组成部分，它是一种防范能力较强的综合系统。视频监控以其直观、准确、及时和信息内容丰富而广泛应用于许多场合。近年来，随着计算机、网络以及图像处理、传输技术的飞速发展，视频监控技术也有了长足的发展。视频监控技术在安全防范系统中占有重要的地位。"
	c.Data["CompriseItems"] = CompriseItems 
	c.Data["FunctionItems"] = FunctionItems 
}

func (c *MainController) ProcPageFangdaobaojing() { 
	c.Layout = "layout.html" 
	c.TplName = "ruodian1.html"

	var CompriseItems, FunctionItems map[string]string 
	CompriseItems = map[string]string{} 
	FunctionItems = map[string]string{} 
	CompriseItems["报警控制器组成"] = "报警控制器是一台主机（如电脑的主机一样），用来控制包括有线/无线信号的处理，系统本身故障的检测，电源部分，信号输入，信号输出，内置拨号器等这几个方面，一个防盗报警系统中报警控制器是必不可少的。前端探测器包括有：门磁开关、玻璃破碎探测器、红外探测器和红外/微波双鉴器、电子围栏、外对射、紧急呼救按钮等。;" 

	FunctionItems["功能现状分析"] = "使用产品更稳定可靠，如提升产品抗破坏、抗雷击浪涌、抗干扰等性能；增加更多的功能，方便安装调试与使用，如探测器可调频率、可设置探测区域、可防遮挡、防破坏、探测器与主机无线连接等；更方便的安装与调试，探测器与主机之间使用无线连接，简化安装施工工作量，适应各种家庭用户需求；朝更智能化方向发展，如增加区域管理、语音提醒、生物识别技术布撤防、报警自动发彩信图片、远程控制等技术；更多的联网方式，如由以前单一的电话组网扩展到可以IP网络联网、无线化的GSM/GPRS联网、VPN专线联网等;通过网络的加密技术，可以保障数据更加安全，并可以通过有线与无线互相之间的主从备份功能，使警情传送更快、更可靠；更强大的集成与扩展性，多种报警内容的命令输出，联动其它如视频监控、门禁等，实现各系统之间的协调统一，使得报警安防更自动化、智能化，安全防范进一步得到提高" 

	c.Data["CrtTitle"] = "防盗报警系统" 
	c.Data["TitleDesc"] = "在城市的建筑当中，所使用的监控设备逐渐的增加，于是很多的人开始关注监控设备的使用。随着人们在社会中对自身安全度和生活舒适度的认知提高了，财产安全等等都需要得到更大范围的保护，因此各种安全防范系统就应运而生了，其中防盗报警系统是具有代表性的安全防范设施之一了。防盗报警系统的设备一般分为：前端探测器，报警控制器。"
	c.Data["CompriseItems"] = CompriseItems 
	c.Data["FunctionItems"] = FunctionItems 
}

func (c *MainController) ProcPageDuomeitihuiyishi() { 
	c.Layout = "layout.html" 
	c.TplName = "ruodian1.html"

	var CompriseItems, FunctionItems map[string]string 
	CompriseItems = map[string]string{} 
	FunctionItems = map[string]string{} 
	CompriseItems["多媒体会议室组成"] = "多功能会议系统包含对会议室进行会议、音频、视频、摄像、录音录像及远程会议系统设计，实现会议室的多媒体会议功能，满足简洁流畅的会议过程、逼真传神的听觉效果、清晰舒适的视频显示、智能的摄像跟踪、完整的会议记录、方便快捷的远程会议等要求。整个系统大致分为显示与监视、拾音与扩声、信号传输与控制摄像照明灯光、电视电话会议等五个部分。由大屏幕显示、等离子显示、音响、视频监视录像、多媒体音视频信号源矩阵切换和中央集成控制、摄像照明灯光及控制等部分组成。;" 

	FunctionItems["先进的展示功能"] = "能有效表现与展示各个领域的最新内容。这就需要使用当今先进的多媒体表现形式：文字、语言、图片、连续影像、电脑文档等。;" 
	FunctionItems["完善的会议功能"] = "满足各种级别的会议、学术论坛、报告会等各种形式的会议活动。满足新闻发会、庆功会、签字仪式等公共活动。"
	FunctionItems["强大的培训功能"] = "技术培训、教学、产品介绍；大型年会、月会、公司例会。"
	FunctionItems["一般性演出功能"] = "能进行一般非专业性的文艺晚会形式的文艺演出，举行各种晚宴、酒会及舞会等。频信号混合和切换，并将处理结果送回参加会议的终端。"

	c.Data["CrtTitle"] = "多媒体会议室" 
	c.Data["TitleDesc"] = "随着以电脑为中心的多媒体技术的普及和提高，给会议工作带来了新的手段和方法，尤其是近几年，视频会议、远程教学等可视化信息技术在会议室领域得到广泛应用，多媒体会议室以其功能的多样性（如现场会议、学术报告、培训教学等）得到迅速普及。在多媒体会议室里不管是作报告、总结、汇报、介绍产品等等，用电脑互动操作的图、文、声、影、画展示，充分调动了与会者的感官知觉，大大提高了会议效果。可以说，多媒体在办公领域中，也越来越体现出它的优势。"
	c.Data["CompriseItems"] = CompriseItems 
	c.Data["FunctionItems"] = FunctionItems 
}

func (c *MainController) ProcPageYikakong() { 
	c.Layout = "layout.html" 
	c.TplName = "ruodian1.html"

	var CompriseItems, FunctionItems map[string]string 
	CompriseItems = map[string]string{} 
	FunctionItems = map[string]string{} 
	CompriseItems["“一体化”体系结构："] = "系统平台完全按照“一体化”的原则进行设计，采用一个统一的数据平台，统一的管理使用平台，统一的终端基础软硬件平台，统一的卡片规划平台，使用户的总体投资小，售后服务成本低、升级容易、管理简单。更为重要的是一体化的架构使得系统平台结构关系变为简单，各个应用子系统在同一个平台上进行操作，降低了应用系统的扩展和实施的难度。;" 
	CompriseItems["先进的1+X结构："] = "业界先进的“1+X”体系结构，完全支持一个总部，多个分部的管理模式，让管理不再受地域限制。数字化一卡通的终端设备拥有基于一个硬件平台的多个应用程序的功能，数字化一卡通的软件采用1个管理平台加若干（X）个应用子系统的设计模式，充分地体现了平台化的设计理念。"
	CompriseItems["科学的“无关性”设计"] = "业界领先目前先进的“无关性”设计技术，突现了开发性、扩展性和适应性，使得用户单位具有更好的选择权和使用权。无关性”主要表现在以下五个方面：数据平台无关性：支持多种操作平台（Unix、Linux、Windows），保障用户的投资；软件接入无关性：支持多种技术体系的应用软件接入；网络形式无关性：支持TCP/IP、485总线/星型、拨号等多种通信方式；	终端厂家无关性：支持多种厂家终端设备；卡片类型无关性：支持非接触式逻辑加密卡（如Mifare One S50、S70，复旦微电子、清华微电子等多卡种）；支持接触式、非接触式CPU卡（如Mifare Pro系列、NTT系列、MOTOROLA系列等多种卡型）；支持二代身份证。"

	FunctionItems["一卡"] = "在同一张卡上实现多种不同功能的智能管理，一张卡上通行很多的设备，而不是不同功能有不同的卡，不同的机在不同的卡上使用。" 
	FunctionItems["一线"] = "一条线通多种信息，多种不同的设备都挂在一条线上，通过一条线跟PC机一个接口把所有的设备都串起来，进行不同数据的信息交换。"
	FunctionItems["一库"] = "在同一个软件、同一台PC机上、同一个数据库内、实现卡的发行、卡的注销、卡的报失、卡的资料查询等，准确明了、方便快捷。"

	c.Data["CrtTitle"] = "一卡通" 
	c.Data["TitleDesc"] = "一卡通系统是通过一张卡实现多种不同功能的智能管理。一张卡上通行很多的设备，而不是不同功能有不同的卡，不同的卡在不同的设备上使用。多种不同的设备都挂在一条数据线上，通过一条数据线跟管理计算机通讯，在同一套系统软件，同一台计算机上，同一个数据库内，进行不同数据的信息交换。实现卡的发行、取消、报失、卡的资料查询等。各种使用功能准确明了，方便快捷。广泛应用于城市公共交通、高速公路自动收费、智能大厦、各种公共收费、智能小区、物业管理、考勤门禁管理、校园和厂区一卡通系统中。集RF射频技术、智能卡应用技术、计算机网络技术、自动控制技术、于一体，来实现某一管理区域的集中管理，使人们充分享受现代科技给日常工作和生活带来的便利和安全。"
	c.Data["CompriseItems"] = CompriseItems 
	c.Data["FunctionItems"] = FunctionItems 
}

func (c *MainController) ProcPagePinjieping() { 
	c.Layout = "layout.html" 
	c.TplName = "ruodian1.html"

	var CompriseItems, FunctionItems map[string]string 
	CompriseItems = map[string]string{} 
	FunctionItems = map[string]string{} 
	CompriseItems["采集系统"] = "拼接屏技术飞速发展，现在的专用液晶屏（SLCD）是LCD的一个高档衍生品种，采用世界先进的工业级液晶面板，采用色彩饱和度高达97%，亮度可达700cd/㎡，对比度为3000：1，刷新频率120Hz，可视角度高达178度，完整包装可抵抗1米以下跌落冲击，使用寿命长达6万小时。SLCD是一个完整的拼接显示单元，既能单独作为显示器使用，又可以拼接成超大屏幕使用。根据不同使用需求，实现可变大也可变小的百变大屏功能：单屏分割显示、单屏单独显示、任意组合显示、液晶图像拼接、全屏拼接、竖屏显示，图像边框可选补偿或遮盖，支持数字信号的漫游、缩放拉伸、跨屏显示，各种显示预案的设置和运行，全高清信号实时处理。" 

	FunctionItems["超长寿命"] = "随着科学技术的进步，液晶产品的使用寿命已经可以达到50000个小时以上(主要受背光源影响)，目前市场主流的LED背光源技术已经基本摆脱了CCFL背光源时代的限制，其使用寿命可以达到10万小时以上，亮度可以达到1000流明，且不会在长时间的使用后出现光源变暗的情况。" 
	FunctionItems["视角大"] = "对于早期的液晶产品而言，可视角度曾经是制约液晶的一个大问题，但随着液晶技术的不断进步，已经完全解决了这个问题，像使用了DID、IPS技术的拼接屏，其可视角度可以达到史无前例的178度。"
	FunctionItems["分辨率高"] = "按照目前市场主流技术，液晶显示屏的物理分辨率可以轻易达到肉眼无法分辨的视网膜级，液晶的亮度和对比度都很高，色彩鲜艳亮丽，图像稳定不闪烁。"
	FunctionItems["超薄轻巧"] = "液晶具有厚度薄，重量轻的特点，可以方便地拼接和安装。55寸专用液晶屏，重量只有30KG，厚度不到11公分，这是其它显示器件所不能比拟的。"
	FunctionItems["功耗小"] = "液晶显示设备，小功率，低发热一向为人们所称道，待机情况下仅为3W。"
	FunctionItems["故障率低"] = "液晶是目前稳定可靠的显示设备，由于发热量很小，器件很稳定，不会因为元器件温升过高损坏而造成故障。"

	c.Data["CrtTitle"] = "拼接屏" 
	c.Data["TitleDesc"] = "拼接屏是完整的成品，即挂即用，安装就像搭积木一样简单，单个或多个液晶屏的拼接使用及安装都非常简单。拼接屏表面还带钢化玻璃保护层、拼接屏内置智能温控报警电路及特有的“快散”散热系统。"
	c.Data["CompriseItems"] = CompriseItems 
	c.Data["FunctionItems"] = FunctionItems 
}


// ------------------------- RuoDian Page end -------------------------------------------


// ------------------------- XiTongJiChen Page start -------------------------------------------

func (c *MainController) ProcPageWangluogaizhao() { 
	c.Layout = "layout.html" 
	c.TplName = "xitongjicheng.html"

	var CompriseItems, FunctionItems map[string]string 
	CompriseItems = map[string]string{} 
	FunctionItems = map[string]string{} 
	CompriseItems["“一体化”体系结构："] = "我们通过结构化的综合布线系统和计算机网络技术，将各个分离的设备(如个人电脑)、功能和信息等集成到相互关联的、统一和协调的系统之中，使资源达到充分共享，实现集中、高效、便利的管理。系统集成采用功能集成、BSV液晶拼接集成、综合布线、网络集成、软件界面集成等多种集成技术。系统集成实现的关键在于解决系统之间的互连和互操作性问题，它是一个多厂商、多协议和面向各种应用的体系结构。这需要解决各类设备、子系统间的接口、协议、系统平台、应用软件等与子系统、建筑环境、施工配合、组织管理和人员配备相关的一切面向集成的问题。"
	CompriseItems["先进的1+X结构："] = "业界先进的“1+X”体系结构，完全支持一个总部，多个分部的管理模式，让管理不再受地域限制。数字化一卡通的终端设备拥有基于一个硬件平台的多个应用程序的功能，数字化一卡通的软件采用1个管理平台加若干（X）个应用子系统的设计模式，充分地体现了平台化的设计理念。"
	CompriseItems["科学的“无关性”设计"] = "业界领先目前先进的“无关性”设计技术，突现了开发性、扩展性和适应性，使得用户单位具有更好的选择权和使用权。无关性”主要表现在以下五个方面：数据平台无关性：支持多种操作平台（Unix、Linux、Windows），保障用户的投资；软件接入无关性：支持多种技术体系的应用软件接入；网络形式无关性：支持TCP/IP、485总线/星型、拨号等多种通信方式；	终端厂家无关性：支持多种厂家终端设备；卡片类型无关性：支持非接触式逻辑加密卡（如Mifare One S50、S70，复旦微电子、清华微电子等多卡种）；支持接触式、非接触式CPU卡（如Mifare Pro系列、NTT系列、MOTOROLA系列等多种卡型）；支持二代身份证。"

	FunctionItems["一卡"] = "在同一张卡上实现多种不同功能的智能管理，一张卡上通行很多的设备，而不是不同功能有不同的卡，不同的机在不同的卡上使用。" 
	FunctionItems["一线"] = "一条线通多种信息，多种不同的设备都挂在一条线上，通过一条线跟PC机一个接口把所有的设备都串起来，进行不同数据的信息交换。"
	FunctionItems["一库"] = "在同一个软件、同一台PC机上、同一个数据库内、实现卡的发行、卡的注销、卡的报失、卡的资料查询等，准确明了、方便快捷。"

	c.Data["CrtTitle"] = "网络改造" 
    c.Data["TitleDesc"] = "我们的系统集成（SI，System Integration）服务经过多年经验的累积已经形成了蓝盟独有的方法论，确保了系统集成各模块的有效耦合，达到高性价比的交付。"
	c.Data["CompriseItems"] = CompriseItems 
	c.Data["FunctionItems"] = FunctionItems 
}

func (c *MainController) ProcPageWuxianfugai() { 
}

func (c *MainController) ProcPageFangwenkongzhi() { 
	
}

func (c *MainController) ProcPageShangwangguanli() { 
	
}

func (c *MainController) ProcPageFwqxulihua() { 
	
}

func (c *MainController) ProcPageZuomianxulihua() { 
	
}

func (c *MainController) ProcPageVpnservice() { 
	
}

func (c *MainController) ProcPageDatabackup() { 
	
}

func (c *MainController) ProcPageEmailserver() { 
	
}

func (c *MainController) ProcPageFileshare() { 
	
}

func (c *MainController) ProcPageServerLoadBalance() { 
	
}



// ------------------------- XiTongJiChen Page end -------------------------------------------



