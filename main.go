package main

import (
	"github.com/astaxie/beego"
	"github.com/sapronlee/goba/modules/filter"
	"github.com/sapronlee/goba/modules/plugin"
	"github.com/sapronlee/goba/modules/template"
	_ "github.com/sapronlee/goba/routers"
)

func main() {
	beego.AddAPPStartHook(plugin.InitAsset)
	beego.AddAPPStartHook(filter.AddLogFilter)
	beego.AddAPPStartHook(filter.AddHandleFilter)
	beego.AddAPPStartHook(template.AddTemplateFunc)

	beego.Run()
}
