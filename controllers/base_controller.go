package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
	c.Data["AppName"] = beego.AppName
	c.Layout = "layout/base.tpl"
	c.RunInheritPrepare()
}

type NestPreparer interface {
	NestPrepare()
}

type AdminNestPreparer interface {
	AdminNestPrepare()
}

func (c *BaseController) RunInheritPrepare() {
	if app, ok := c.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}
