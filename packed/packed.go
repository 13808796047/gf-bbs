package packed

import (
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gpage"
)

func WrapContent(page *gpage.Page) string {
	content := page.GetContent(4)
	content = gstr.ReplaceByMap(content, map[string]string{
		"<span":  "<li class='page-item disabled'><a ",
		"/span>": "/a></li>",
		"<a":     "<li class='page-item'><a",
		"/a>":    "/a></li>",
	})
	return "<nav aria-label='Page navigation example'> <style>.GPageSpan{background:#007bff;color:#fff}.pagination a{display:inline-block;padding:5px 10px;margin-left:-1px;border:1px solid #eee;}</style><ul class='pagination justify-content-center'>" + content + "</ul></nav>"
}
