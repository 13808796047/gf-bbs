package api

import (
	"gf-bbs/app/dao"
	"gf-bbs/app/model"
	"gf-bbs/packed"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type TopicController struct{}

func (t *TopicController) Index(r *ghttp.Request) {
	var topics []*model.Topics
	dao.Topics.TopicsDao.Ctx(r.Context()).Offset((r.GetInt("page") - 1) * 10).Limit(10).With(model.Topics{}.Categories).Scan(&topics)
	count, _ := dao.Topics.TopicsDao.Ctx(r.Context()).Count()
	page := r.GetPage(count, 10)
	r.Response.WriteTpl("layouts/app.html", g.Map{
		"title":   "话题列表",
		"content": "topics/index.html",
		"topics":  topics,
		"page3":   packed.WrapContent(page),
	})
}
