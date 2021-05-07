/**
* @Author: Rocks
* @Email: 451360994@qq.com
* @Description:
* @File:  t_rule_annotations_controller
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

type TRuleAnnotationsController struct {
    base.BaseRouter
}

var (
    controllerNameTRuleAnnotations = "TRuleAnnotationsController"
)

//TRuleAnnotations页面信息
func (controller *TRuleAnnotationsController) Index(r *ghttp.Request) {
    base.WriteTpl(r, "rule/annotations.html", g.Map{})
}

//获取TRuleAnnotations单条信息
func (controller *TRuleAnnotationsController) Get(r *ghttp.Request) {
    id := r.GetInt("id")
    model := model.TRuleAnnotations{Id: id}.Get()
    if model.Id <= 0 {
       base.Fail(r, controllerNameTRuleAnnotations+" get fail")
    }
    base.Succ(r, model)
}

//根据id或者ids删除TRuleAnnotations
func (controller *TRuleAnnotationsController) Delete(r *ghttp.Request) {
   ids := r.GetInts("ids")
    for _, id := range ids {
        model := model.TRuleAnnotations{Id: id}
        model.Delete()
    }
    base.Succ(r, "")
}


//创建一条TRuleAnnotations
func (controller *TRuleAnnotationsController) Save(r *ghttp.Request) {
    model := model.TRuleAnnotations{}
    err := gconv.Struct(r.GetPostMap(), &model)
    if err != nil {
      glog.Error( controllerNameTRuleAnnotations+" save struct error", err)
      base.Error(r, "save error")
    }
    var num int64
    if model.Id <= 0 {
      num = model.Insert()
    } else {
      num = model.Update()
    }

    if num <= 0 {
       base.Fail(r, controllerNameTRuleAnnotations+" save fail")
    }

    base.Succ(r, "")
}

//更新一条TRuleAnnotations
func (controller *TRuleAnnotationsController) Update(r *ghttp.Request) {
    model := model.TRuleAnnotations{}
    err := gconv.Struct(r.GetPostMap(), &model)
    if err != nil {
    glog.Error( controllerNameTRuleAnnotations+" save struct error", err)
    base.Error(r, "save error")
    }
    var num int64
    if model.Id <= 0 {
    num = model.Insert()
    } else {
    num = model.Update()
    }

    if num <= 0 {
    base.Fail(r, controllerNameTRuleAnnotations+" save fail")
    }

    base.Succ(r, "")
}

//分页列表TRuleAnnotations
func (controller *TRuleAnnotationsController) Page(r *ghttp.Request) {
   form := base.NewForm(r.GetQueryMap())
   model := model.TRuleAnnotations{}
   page := model.Page(&form)
   base.Succ(r, g.Map{"list": page, "form": form})
}