package admin

import (
	"github.com/sapronlee/goba/controllers"
)

type BaseController struct {
	controllers.BaseController
}

func (c *BaseController) Prepare() {
	c.BaseController.Prepare()
	c.RunInheritPrepare()
}

func (c *BaseController) RunInheritPrepare() {
	if app, ok := c.AppController.(controllers.AdminNestPreparer); ok {
		app.AdminNestPrepare()
	}
}
