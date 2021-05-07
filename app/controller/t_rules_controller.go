/**
* @Author: Rocks
* @Email: 451360994@qq.com
* @Description:
* @File:  t_rules_controller
* @Version: 1.0.0
* @Date: 2021-04-22
*/

package controller

import (
    "github.com/gogf/gf/frame/g"
    "github.com/gogf/gf/net/ghttp"
    "github.com/gogf/gf/os/glog"
    "github.com/gogf/gf/util/gconv"
    "mysql4promUI/app/model"
    "mysql4promUI/utils/base"
)

type TRulesController struct {
    base.BaseRouter
}

var (
    controllerNameTRules = "TRulesController"
)

//TRules页面信息
func (controller *TRulesController) Index(r *ghttp.Request) {
    //base.WriteTpl(r, "rules.html", g.Map{})

    err := r.Response.WriteTpl("rules.html", g.Map{})
    if err != nil {
        glog.Error(err)
    }
}

//获取TRules单条信息
func (controller *TRulesController) Get(r *ghttp.Request) {
    id := r.GetInt("id")
    model := model.TRules{Id: id}.Get()
    if model.Id <= 0 {
       base.Fail(r, controllerNameTRules+" get fail")
    }
    base.Succ(r, model)
}

//根据id或者ids删除TRules
func (controller *TRulesController) Delete(r *ghttp.Request) {
   ids := r.GetInts("ids")
    for _, id := range ids {
        model := model.TRules{Id: id}
        model.Delete()
    }
    base.Succ(r, "")
}

//创建一条TRules
func (controller *TRulesController) Save(r *ghttp.Request) {
    model := model.TRules{}
    err := gconv.Struct(r.GetPostMap(), &model)
    if err != nil {
      glog.Error( controllerNameTRules+" save struct error", err)
      base.Error(r, "save error")
    }
    var num int64
    if model.Id <= 0 {
      num = model.Insert()
    } else {
      num = model.Update()
    }

    if num <= 0 {
       base.Fail(r, controllerNameTRules+" save fail")
    }

    base.Succ(r, "")
}

//更新一条TRules
func (controller *TRulesController) Update(r *ghttp.Request) {
    model := model.TRules{}
    err := gconv.Struct(r.GetPostMap(), &model)
    if err != nil {
        glog.Error( controllerNameTRules+" save struct error", err)
        base.Error(r, "save error")
    }
    var num int64
    if model.Id <= 0 {
        num = model.Insert()
    } else {
        num = model.Update()
    }

    if num <= 0 {
        base.Fail(r, controllerNameTRules+" save fail")
    }

    base.Succ(r, "")
}

//分页列表TRules
func (controller *TRulesController) Page(r *ghttp.Request) {
   form := base.NewForm(r.GetQueryMap())
   model := model.TRules{}
   page := model.Page(&form)
   base.Succ(r, g.Map{"list": page, "form": form})
}