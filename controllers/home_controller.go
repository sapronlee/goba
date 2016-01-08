package controllers

type HomeController struct {
	BaseController
}

func (c *HomeController) NestPrepare() {
}

func (c *HomeController) Index() {
	c.TplNames = "index.tpl"
}
