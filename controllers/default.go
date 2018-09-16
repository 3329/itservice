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

func (c *MainController) ProcAboutus() {
	c.Layout = "layout.html"
	c.TplName = "aboutus.html"
}

func (c *MainController) ProcService() {
	c.Layout = "layout.html"
	c.TplName = "service.html"
}

func (c *MainController) ProcJoinus() {
	c.Layout = "layout.html"
	c.TplName = "joinus.html"
}

func (c *MainController) ProcContactus() {
	c.Layout = "layout.html"
	c.TplName = "contactus.html"
}


