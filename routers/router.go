package routers

import (
	"github.com/astaxie/beego"
	"github.com/sapronlee/goba/controllers"
	"github.com/sapronlee/goba/controllers/admin"
)

func init() {
	beego.Router("/", &controllers.HomeController{}, "get:Index")

	beego.Router("/admin", &admin.HomeController{}, "get:Index")
}
