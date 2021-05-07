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

type TConfigs struct {
    Id          int         `orm:"id,primary"   json:"Id,omitempty" gconv:"Id,omitempty"`                    
    ConfigKey   string      `orm:"config_key"   json:"ConfigKey,omitempty" gconv:"ConfigKey,omitempty"`      
    ConfigValue string      `orm:"config_value" json:"ConfigValue,omitempty" gconv:"ConfigValue,omitempty"`  
    State       int         `orm:"state"        json:"State,omitempty" gconv:"State,omitempty"`              
    CreatedAt   *gtime.Time `orm:"created_at"   json:"CreatedAt,omitempty" gconv:"CreatedAt,omitempty"`      
    UpdatedAt   *gtime.Time `orm:"updated_at"   json:"UpdatedAt,omitempty" gconv:"UpdatedAt,omitempty"`      
}

//创建mTConfigs
func (model *TConfigs) Insert() int64 {
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
//删除TConfigs
func (model TConfigs) Delete() int64 {
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

//更新TConfigs
func (model *TConfigs) Update() int64 {
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

//根据ID查询TConfigs
func (model TConfigs) Get() TConfigs {
    if model.Id <= 0 {
        glog.Error(model.TableName() + " get id error")
        return TConfigs{}
    }
    var resData TConfigs
    err := model.dbModel("t").Where(" id = ? and status=1", model.Id).Fields(model.columns()).Struct(&resData)
    if err != nil {
        glog.Error(model.TableName()+" get one error", err)
        return TConfigs{}
    }
    return resData
}

//分页查询TConfigs
func (model TConfigs) Page(form *base.BaseForm) []TConfigs {
    if form.Page <= 0 || form.Rows <= 0 {
        glog.Error(model.TableName()+" page param error", form.Page, form.Rows)
        return []TConfigs{}
    }
    where := " status = 1 "
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
        return []TConfigs{}
    }
    var resData []TConfigs
    pageNum, pageSize := (form.Page-1)*form.Rows, form.Rows
    dbModel := model.dbModel("t").Fields(model.columns()).Fields(model.columns())
    err = dbModel.Where(where, params).Limit(pageNum, pageSize).OrderBy(form.OrderBy).Structs(&resData)
    if err != nil {
        glog.Error(model.TableName()+" page list error", err)
        return []TConfigs{}
    }

    return resData
}

//返回数据库TConfigs
func (model TConfigs) dbModel(alias ...string) *gdb.Model {
    var tmpAlias string
    if len(alias) > 0 {
    tmpAlias = " " + alias[0]
    }
    tableModel := g.DB("center").Table(model.TableName() + tmpAlias).Safe()
    return tableModel
}
//返回主键TConfigs
func (model TConfigs) PkVal() int {
    return model.Id
}
//表名TConfigs
func (model TConfigs) TableName() string {
    return "t_configs"
}
//列名TConfigs
func (model TConfigs) columns() string {
    sqlColumns := "t.id as Id,t.config_key as ConfigKey,t.config_value as ConfigValue,t.state as State,t.created_at as CreatedAt,t.updated_at as UpdatedAt"
    return sqlColumns
}