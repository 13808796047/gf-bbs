package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type HomeController struct{}

func (c *HomeController) Index(r *ghttp.Request) {
	r.Response.WriteTpl("layouts/app.html", g.Map{
		"title":   "首页",
		"content": "pages/index.html",
	})
}
