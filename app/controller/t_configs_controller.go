/**
* @Author: Rocks
* @Email: 451360994@qq.com
* @Description:
* @File:  t_configs_controller
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

type TConfigsController struct {
    base.BaseRouter
}

var (
    controllerNameTConfigs = "TConfigsController"
)

//TConfigs页面信息
func (controller *TConfigsController) Index(r *ghttp.Request) {
    base.WriteTpl(r, "configs.html", g.Map{})
}

//获取TConfigs单条信息
func (controller *TConfigsController) Get(r *ghttp.Request) {
    id := r.GetInt("id")
    model := model.TConfigs{Id: id}.Get()
    if model.Id <= 0 {
       base.Fail(r, controllerNameTConfigs+" get fail")
    }
    base.Succ(r, model)
}

//根据id或者ids删除TConfigs
func (controller *TConfigsController) Delete(r *ghttp.Request) {
   ids := r.GetInts("ids")
    for _, id := range ids {
        model := model.TConfigs{Id: id}
        model.Delete()
    }
    base.Succ(r, "")
}


//创建一条TConfigs
func (controller *TConfigsController) Save(r *ghttp.Request) {
    model := model.TConfigs{}
    err := gconv.Struct(r.GetPostMap(), &model)
    if err != nil {
      glog.Error( controllerNameTConfigs+" save struct error", err)
      base.Error(r, "save error")
    }
    var num int64
    if model.Id <= 0 {
      num = model.Insert()
    } else {
      num = model.Update()
    }

    if num <= 0 {
       base.Fail(r, controllerNameTConfigs+" save fail")
    }

    base.Succ(r, "")
}

//更新一条TConfigs
func (controller *TConfigsController) Update(r *ghttp.Request) {
    model := model.TConfigs{}
    err := gconv.Struct(r.GetPostMap(), &model)
    if err != nil {
    glog.Error( controllerNameTConfigs+" save struct error", err)
    base.Error(r, "save error")
    }
    var num int64
    if model.Id <= 0 {
    num = model.Insert()
    } else {
    num = model.Update()
    }

    if num <= 0 {
    base.Fail(r, controllerNameTConfigs+" save fail")
    }

    base.Succ(r, "")
}

//分页列表TConfigs
func (controller *TConfigsController) Page(r *ghttp.Request) {
   form := base.NewForm(r.GetQueryMap())
   model := model.TConfigs{}
   page := model.Page(&form)
   base.Succ(r, g.Map{"list": page, "form": form})
}