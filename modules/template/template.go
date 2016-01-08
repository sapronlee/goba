package template

import (
	"github.com/astaxie/beego"
	"github.com/shaoshing/train"
)

func AddTemplateFunc() error {
	beego.AddFuncMap("javascriptTag", train.JavascriptTag)
	beego.AddFuncMap("stylesheetTag", train.StylesheetTag)
	beego.AddFuncMap("stylesheetTagWithParam", train.StylesheetTagWithParam)
	return nil
}
