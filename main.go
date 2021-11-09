package main

import (
	_ "gf-bbs/boot"
	_ "gf-bbs/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
