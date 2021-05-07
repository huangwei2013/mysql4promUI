package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/text/gstr"
	_ "mysql4promUI/boot"
	"strings"
)


func CamelCase(name string) string {
	return gstr.CamelCase(name)
}
func ScamelCase(name string) string {
	return gstr.CamelLowerCase(name)
}
func CameTable(name string) string {
	return strings.ToLower(gstr.Replace(name, "_", "/"))
}
func SubTable(name string) string {
	tables := gstr.Split(name, "_")
	return strings.Join(tables[1:],"_")
}

func main() {

	v := g.View()
	v.BindFunc("CamelCase", CamelCase)
	v.BindFunc("ScamelCase", ScamelCase)
	v.BindFunc("CameTable", CameTable)
	v.BindFunc("SubTable", SubTable)

	v.Assign("baseURL", "http://localhost:18082/")

	g.Server().Run()
}
