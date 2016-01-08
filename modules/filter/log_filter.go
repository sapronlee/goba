package filter

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strings"
)

type logFilter struct {
	UrlPrefix []string
	UrlSuffix []string
}

func AddLogFilter() error {
	beego.DefaultLogFilter = &logFilter{
		UrlPrefix: []string{"/assets/"},
		UrlSuffix: []string{},
	}
	return nil
}

func (f *logFilter) Filter(ctx *context.Context) bool {
	url := ctx.Request.RequestURI

	if f.filterUrlPrefix(url) || f.filterUrlSuffix(url) {
		return true
	}

	return false
}

func (f *logFilter) filterUrlPrefix(url string) bool {
	for _, prefix := range f.UrlPrefix {
		if strings.HasPrefix(url, prefix) {
			return true
		}
	}
	return false
}

func (f *logFilter) filterUrlSuffix(url string) bool {
	for _, suffix := range f.UrlSuffix {
		if strings.HasSuffix(url, suffix) {
			return true
		}
	}
	return false
}
