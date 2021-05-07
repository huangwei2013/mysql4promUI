/**
* @Author: Rocks
* @Email: 451360994@qq.com
* @Description:
* @File:  t_rule_labels_controller
* @Version: 1.0.0
* @Date: 2021-04-22
*/

package controller


import (
   "mysql4promUI/utils/base"
    "mysql4promUI/app/model"
   "github.com/gogf/gf/frame/g"
   "github.com/gogf/gf/net/ghttp"
   "github.com/gogf/gf/os/glog"
   "github.com/gogf/gf/util/gconv"
)

type TRuleLabelsController struct {
    base.BaseRouter
}

var (
    controllerNameTRuleLabels = "TRuleLabelsController"
)

//TRuleLabels页面信息
func (controller *TRuleLabelsController) Index(r *ghttp.Request) {
    base.WriteTpl(r, "rule/labels.html", g.Map{})
}

//获取TRuleLabels单条信息
func (controller *TRuleLabelsController) Get(r *ghttp.Request) {
    id := r.GetInt("id")
    model := model.TRuleLabels{Id: id}.Get()
    if model.Id <= 0 {
       base.Fail(r, controllerNameTRuleLabels+" get fail")
    }
    base.Succ(r, model)
}

//根据id或者ids删除TRuleLabels
func (controller *TRuleLabelsController) Delete(r *ghttp.Request) {
   ids := r.GetInts("ids")
    for _, id := range ids {
        model := model.TRuleLabels{Id: id}
        model.Delete()
    }
    base.Succ(r, "")
}


//创建一条TRuleLabels
func (controller *TRuleLabelsController) Save(r *ghttp.Request) {
    model := model.TRuleLabels{}
    err := gconv.Struct(r.GetPostMap(), &model)
    if err != nil {
      glog.Error( controllerNameTRuleLabels+" save struct error", err)
      base.Error(r, "save error")
    }
    var num int64
    if model.Id <= 0 {
      num = model.Insert()
    } else {
      num = model.Update()
    }

    if num <= 0 {
       base.Fail(r, controllerNameTRuleLabels+" save fail")
    }

    base.Succ(r, "")
}

//更新一条TRuleLabels
func (controller *TRuleLabelsController) Update(r *ghttp.Request) {
    model := model.TRuleLabels{}
    err := gconv.Struct(r.GetPostMap(), &model)
    if err != nil {
    glog.Error( controllerNameTRuleLabels+" save struct error", err)
    base.Error(r, "save error")
    }
    var num int64
    if model.Id <= 0 {
    num = model.Insert()
    } else {
    num = model.Update()
    }

    if num <= 0 {
    base.Fail(r, controllerNameTRuleLabels+" save fail")
    }

    base.Succ(r, "")
}

//分页列表TRuleLabels
func (controller *TRuleLabelsController) Page(r *ghttp.Request) {
   form := base.NewForm(r.GetQueryMap())
   model := model.TRuleLabels{}
   page := model.Page(&form)
   base.Succ(r, g.Map{"list": page, "form": form})
}