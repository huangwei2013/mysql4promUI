package model

import (
    "github.com/gogf/gf/util/gconv"
    "mysql4promUI/utils/base"
    "github.com/gogf/gf/database/gdb"
    "github.com/gogf/gf/frame/g"
    "github.com/gogf/gf/os/glog"
)

import (
	"github.com/gogf/gf/os/gtime"
)

type TRuleLabels struct {
    Id         int         `orm:"id,primary"  json:"Id,omitempty" gconv:"Id,omitempty"`                  
    RuleId     int         `orm:"rule_id"     json:"RuleId,omitempty" gconv:"RuleId,omitempty"`          
    LabelKey   string      `orm:"label_key"   json:"LabelKey,omitempty" gconv:"LabelKey,omitempty"`      
    LabelValue string      `orm:"label_value" json:"LabelValue,omitempty" gconv:"LabelValue,omitempty"`  
    State      int         `orm:"state"       json:"State,omitempty" gconv:"State,omitempty"`            
    CreatedAt  *gtime.Time `orm:"created_at"  json:"CreatedAt,omitempty" gconv:"CreatedAt,omitempty"`    
    UpdatedAt  *gtime.Time `orm:"updated_at"  json:"UpdatedAt,omitempty" gconv:"UpdatedAt,omitempty"`    
}

//创建mTRuleLabels
func (model *TRuleLabels) Insert() int64 {
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
//删除TRuleLabels
func (model TRuleLabels) Delete() int64 {
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

//更新TRuleLabels
func (model *TRuleLabels) Update() int64 {
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

//根据ID查询TRuleLabels
func (model TRuleLabels) Get() TRuleLabels {
    if model.Id <= 0 {
        glog.Error(model.TableName() + " get id error")
        return TRuleLabels{}
    }
    var resData TRuleLabels
    err := model.dbModel("t").Where(" id = ? and state=1", model.Id).Fields(model.columns()).Struct(&resData)
    if err != nil {
        glog.Error(model.TableName()+" get one error", err)
        return TRuleLabels{}
    }
    return resData
}

//分页查询TRuleLabels
func (model TRuleLabels) Page(form *base.BaseForm) []TRuleLabels {
    if form.Page <= 0 || form.Rows <= 0 {
        glog.Error(model.TableName()+" page param error", form.Page, form.Rows)
        return []TRuleLabels{}
    }
    where := " state = 1 "
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
        return []TRuleLabels{}
    }
    var resData []TRuleLabels
    pageNum, pageSize := (form.Page-1)*form.Rows, form.Rows
    dbModel := model.dbModel("t").Fields(model.columns()).Fields(model.columns())
    err = dbModel.Where(where, params).Limit(pageNum, pageSize).OrderBy(form.OrderBy).Structs(&resData)
    if err != nil {
        glog.Error(model.TableName()+" page list error", err)
        return []TRuleLabels{}
    }

    return resData
}

//返回数据库TRuleLabels
func (model TRuleLabels) dbModel(alias ...string) *gdb.Model {
    var tmpAlias string
    if len(alias) > 0 {
    tmpAlias = " " + alias[0]
    }
    tableModel := g.DB("center").Table(model.TableName() + tmpAlias).Safe()
    return tableModel
}
//返回主键TRuleLabels
func (model TRuleLabels) PkVal() int {
    return model.Id
}
//表名TRuleLabels
func (model TRuleLabels) TableName() string {
    return "t_rule_labels"
}
//列名TRuleLabels
func (model TRuleLabels) columns() string {
    sqlColumns := "t.state as State,t.created_at as CreatedAt,t.updated_at as UpdatedAt,t.id as Id,t.rule_id as RuleId,t.label_key as LabelKey,t.label_value as LabelValue"
    return sqlColumns
}