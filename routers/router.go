package routers

import (
	"github.com/sapronlee/goba/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
