package model

import (
    "github.com/gogf/gf/util/gconv"
    "mysql4promUI/utils/base"
    "github.com/gogf/gf/database/gdb"
    "github.com/gogf/gf/frame/g"
    "github.com/gogf/gf/os/glog"

    "github.com/gogf/gf/os/gtime"
)


type TRules struct {
    Id           int         `orm:"id,primary"    json:"Id,omitempty" gconv:"Id,omitempty"`                      
    RuleName     string      `orm:"rule_name"     json:"RuleName,omitempty" gconv:"RuleName,omitempty"`          
    RuleFn       string      `orm:"rule_fn"       json:"RuleFn,omitempty" gconv:"RuleFn,omitempty"`              
    RuleInterval int         `orm:"rule_interval" json:"RuleInterval,omitempty" gconv:"RuleInterval,omitempty"`  
    RuleAlert    string      `orm:"rule_alert"    json:"RuleAlert,omitempty" gconv:"RuleAlert,omitempty"`        
    RuleExpr     string      `orm:"rule_expr"     json:"RuleExpr,omitempty" gconv:"RuleExpr,omitempty"`          
    RuleFor      string      `orm:"rule_for"      json:"RuleFor,omitempty" gconv:"RuleFor,omitempty"`            
    Note         string      `orm:"note"          json:"Note,omitempty" gconv:"Note,omitempty"`                  
    State        int         `orm:"state"         json:"State,omitempty" gconv:"State,omitempty"`                
    CreatedAt    *gtime.Time `orm:"created_at"    json:"CreatedAt,omitempty" gconv:"CreatedAt,omitempty"`        
    UpdatedAt    *gtime.Time `orm:"updated_at"    json:"UpdatedAt,omitempty" gconv:"UpdatedAt,omitempty"`        
}

//创建mTRules
func (model *TRules) Insert() int64 {
    //model.CreateTime = gtime.Now()
    //model.UpdateTime = gtime.Now()
    //model.Status=1
    r, err := model.dbModel().Data(model).Insert()
    if err != nil {
        glog.Error(model.TableName()+" insert error", err)
        return 0
    }

    res, err2 := r.RowsAffected()
    if err2 != nil {
        glog.Error(model.TableName()+" insert res error", err2)
        return 0
    } else if res > 0 {
        lastId, err2 := r.LastInsertId()
        if err2 != nil {
            glog.Error(model.TableName()+" LastInsertId res error", err2)
            return 0
        } else {
           model.Id = gconv.Int(lastId)
        }
    }
    return res
}
//删除TRules
func (model TRules) Delete() int64 {
    if model.Id <= 0 {
        glog.Error(model.TableName() + " delete id error")
        return 0
    }
    r, err := model.dbModel().Where(" id = ?", model.Id).Delete()
    if err != nil {
        glog.Error(model.TableName()+" delete error", err)
        return 0
    }
    res, err2 := r.RowsAffected()
    if err2 != nil {
        glog.Error(model.TableName()+" delete res error", err2)
        return 0
    }
    return res
}

//更新TRules
func (model *TRules) Update() int64 {
    //model.UpdateTime = gtime.Now();
    r, err := model.dbModel().Data(model).Where(" id = ?", model.Id).Update()
    if err != nil {
        glog.Error(model.TableName()+" update error", err)
        return 0
    }
    res, err2 := r.RowsAffected()
    if err2 != nil {
        glog.Error(model.TableName()+" update res error", err2)
        return 0
    }
    return res
}

//根据ID查询TRules
func (model TRules) Get() TRules {
    if model.Id <= 0 {
        glog.Error(model.TableName() + " get id error")
        return TRules{}
    }
    var resData TRules
    err := model.dbModel("t").Where(" id = ? and status=1", model.Id).Fields(model.columns()).Struct(&resData)
    if err != nil {
        glog.Error(model.TableName()+" get one error", err)
        return TRules{}
    }
    return resData
}

//分页查询TRules
func (model TRules) Page(form *base.BaseForm) []TRules {
    if form.Page <= 0 || form.Rows <= 0 {
        glog.Error(model.TableName()+" page param error", form.Page, form.Rows)
        return []TRules{}
    }
    //where := " status = 1 "
    where := " 1 = 1 "
    var params []interface{}
    if form.Params != nil && form.Params["name"] != "" {
        where += " and name like ? "
        params = append(params, "%"+form.Params["name"]+"%")
    }

    num, err := model.dbModel("t").Where(where, params).Count()
    form.TotalSize = num
    form.TotalPage = num / form.Rows

    // 没有数据直接返回
    if num == 0 {
        form.TotalPage = 0
        form.TotalSize = 0
        return []TRules{}
    }

    var resData []TRules
    pageNum, pageSize := (form.Page-1)*form.Rows, form.Rows
    dbModel := model.dbModel("t").Fields(model.columns()).Fields(model.columns())
    err = dbModel.Where(where, params).Limit(pageNum, pageSize).OrderBy(form.OrderBy).Structs(&resData)
    if err != nil {
        glog.Error(model.TableName()+" page list error", err)
        return []TRules{}
    }

    return resData
}

//返回数据库TRules
func (model TRules) dbModel(alias ...string) *gdb.Model {
    var tmpAlias string
    if len(alias) > 0 {
    tmpAlias = " " + alias[0]
    }
    tableModel := g.DB("center").Table(model.TableName() + tmpAlias).Safe()
    return tableModel
}
//返回主键TRules
func (model TRules) PkVal() int {
    return model.Id
}
//表名TRules
func (model TRules) TableName() string {
    return "t_rules"
}
//列名TRules
func (model TRules) columns() string {
    sqlColumns := "t.rule_expr as RuleExpr,t.note as Note,t.created_at as CreatedAt,t.id as Id,t.rule_name as RuleName,t.rule_fn as RuleFn,t.rule_interval as RuleInterval,t.rule_alert as RuleAlert,t.rule_for as RuleFor,t.state as State,t.updated_at as UpdatedAt"
    return sqlColumns
}