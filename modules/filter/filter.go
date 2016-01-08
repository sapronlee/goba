package filter

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/shaoshing/train"
)

func AddHandleFilter() error {
	beego.InsertFilter("/assets/*", beego.BeforeStatic, assetFilter)
	return nil
}

func assetFilter(ctx *context.Context) {
	train.ServeRequest(ctx.ResponseWriter, ctx.Request)
}
