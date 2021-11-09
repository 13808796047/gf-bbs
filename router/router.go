package router

import (
	"gf-bbs/app/api"

	"github.com/gogf/gf/frame/g"
)

func init() {
	s := g.Server()
	HomeController := new(api.HomeController)
	s.BindObject("/", HomeController, "Index")
}
