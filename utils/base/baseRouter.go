package base

import (
	"github.com/gogf/gf/frame/g"
	"mysql4promUI/utils/resp"
	"github.com/gogf/gf/net/ghttp"
)


// baseRouter implemented global settings for all other routers.
type BaseRouter struct {
}


func Succ(r *ghttp.Request, data interface{}) {
	r.Response.WriteJson(resp.Succ(data))
	r.Exit()
}

func Fail(r *ghttp.Request, msg string) {
	r.Response.WriteJson(resp.Fail(msg))
	r.Exit()
}

func Error(r *ghttp.Request, msg string) {
	r.Response.WriteJson(resp.Error(msg))
	r.Exit()
}
func Resp(r *ghttp.Request, code int, msg string, data interface{}) {
	r.Response.WriteJson(resp.Resp{
		Code: code,
		Msg:  msg,
		Data: data,
	})
	r.Exit()
}

//带用户写入
func WriteTpl(r *ghttp.Request, tpl string, params g.Map) error {
	data := g.Map{
		//"UserName": GetUser(r).Username,
	}
	return r.Response.WriteTpl(tpl, data)
}
