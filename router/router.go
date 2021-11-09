package router

import (
	"gf-bbs/app/api"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	HomeController := new(api.HomeController)
	s.BindObject("/", HomeController, "Index")
	TopicController := new(api.TopicController)
	s.Group("/topics", func(group *ghttp.RouterGroup) {
		group.GET("/", TopicController, "Index")
	})
}
