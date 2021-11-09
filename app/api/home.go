package api

import (
	"gf-bbs/app/dao"
	"gf-bbs/app/model"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type HomeController struct{}

func (c *HomeController) Index(r *ghttp.Request) {
	var category model.Categories
	dao.Categories.CategoriesDao.Ctx(r.Context()).WithAll().Where("id", 1).Scan(&category)
	r.Response.WriteTpl("layouts/app.html", g.Map{
		"title":   "首页",
		"content": "pages/index.html",
	})
}
