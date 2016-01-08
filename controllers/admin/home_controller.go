package admin

import "github.com/astaxie/beego"

type HomeController struct {
	BaseController
}

func (c *HomeController) AdminNestPrepare() {
	beego.BeeLogger.Error("Admin HomeController Prepare")
}

func (c *HomeController) Index() {
	c.TplNames = "index.tpl"
}
