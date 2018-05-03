package dbdao

import (
    "git.n.xiaomi.com/miot_shop/go_libs/dao"
    "github.com/go-xorm/xorm"
)

type DbBaseDao struct {
    Engine  *xorm.Engine
    Session *xorm.Session
}

type Param interface{}
type ParamNil struct{}
type ParamDesc bool
type ParamIn []interface{}
type ParamRange struct {
    Min interface{}
    Max interface{}
}
type ParamInDesc ParamIn
type ParamRangeDesc ParamRange

func CastToParamIn(input interface{}) ParamIn {
    params := make(ParamIn, 0)
    switch v := input.(type) {
    case []interface{}:
        for _, param := range v {
            params = append(params, param)
        }
    case []int64:
        for _, param := range v {
            params = append(params, param)
        }
    case []int:
        for _, param := range v {
            params = append(params, param)
        }
    case []int32:
        for _, param := range v {
            params = append(params, param)
        }
    case []int8:
        for _, param := range v {
            params = append(params, param)
        }
        return params
    case []string:
        for _, param := range v {
            params = append(params, param)
        }
    default:
        params = append(params, 0)
    }
    return params
}

func CastToParamInDesc(input interface{}) ParamInDesc {
    return ParamInDesc(CastToParamIn(input))
}

func (this *DbBaseDao) buildQuery(session *xorm.Session, input Param, name string) {
    name = session.Engine.Quote(name)
    switch val := input.(type) {
    case ParamDesc:
        if val {
            session = session.Desc(name)
        }
    case ParamIn:
        if len(val) == 1 {
            session = session.And(name+"=?", val[0])
        } else {
            session = session.In(name, val)
        }
    case ParamInDesc:
        if len(val) == 1 {
            session = session.And(name+"=?", val[0])
        } else {
            session = session.In(name, val)
        }
        session = session.Desc(name)
    case ParamRange:
        if val.Min != nil {
            session = session.And(name+">=?", val.Min)
        }
        if val.Max != nil {
            session = session.And(name+"<?", val.Max)
        }
    case ParamRangeDesc:
        if val.Min != nil {
            session = session.And(name+">=?", val.Min)
        }
        if val.Max != nil {
            session = session.And(name+"<?", val.Max)
        }
        session = session.Desc(name)
    case ParamNil:
    case nil:
    default:
        session = session.And(name+"=?", val)
    }
}

func (this *DbBaseDao) UpdateEngine(v ...interface{}) {
    if len(v) == 0 {
        this.Engine = dao.GetDefault("reader").Engine
        this.Session = nil
    } else if len(v) == 1 {
        param := v[0]
        if engine, ok := param.(*xorm.Engine); ok {
            this.Engine = engine
            this.Session = nil
        } else if session, ok := param.(*xorm.Session); ok {
            this.Session = session
            this.Engine = nil
        } else if tpe, ok := param.(bool); ok {
            cluster := "reader"
            if tpe == true {
                cluster = "writer"
            }
            this.Engine = dao.GetDefault(cluster).Engine
            this.Session = nil
        }
    }
}

func (this *DbBaseDao) Create(bean interface{}) (int64, error) {
    return this.GetSession(true).Insert(bean)
}

func (this *DbBaseDao) Update(bean interface{}) (int64, error) {
    session := this.GetSession(true)
    return session.Id(session.Engine.IdOf(bean)).AllCols().Update(bean)
}

func (this *DbBaseDao) UpdateCols(bean interface{}, cols ...string) (int64, error) {
    session := this.GetSession(true)
    return session.Id(session.Engine.IdOf(bean)).Cols(cols...).Update(bean)
}

func (this *DbBaseDao) UpdateWhere(where string, bean interface{}, cols ...string) (int64, error) {
    session := this.GetSession(true)
    session = session.Where(where)
    return session.Cols(cols...).Update(bean)
}

func (this *DbBaseDao) Delete(bean interface{}) (int64, error) {
    session := this.GetSession(true)
    return session.Id(session.Engine.IdOf(bean)).Delete(bean)
}

func (this *DbBaseDao) GetSession(v ...interface{}) *xorm.Session {
    if this.Session != nil {
        return this.Session
    }
    if this.Engine != nil {
        session := this.Engine.NewSession()
        session.IsAutoClose = true
        return session
    }
    return nil
}

