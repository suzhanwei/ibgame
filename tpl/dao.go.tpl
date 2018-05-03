package dbdao

{{$ilen := len .Imports}}
import (
    {{range .Imports}}"{{.}}"{{end}}
    "reflect"
    "github.com/go-xorm/xorm"
)

{{range .Tables}}
{{$tb := Mapper .Name}}
{{$table := .}}
{{$dao := printf "%sDao" $tb}}

func init(){
    DbStructs["{{.Name}}"]=reflect.TypeOf({{$tb}}{})
    DbDaos["{{.Name}}"]=reflect.TypeOf({{$dao}}{})
}

type {{$tb}} struct {
{{range .ColumnsSeq}}{{$col := $table.GetColumn .}} {{Mapper $col.Name}}    {{Type $col}} {{Tag $table $col}}
{{end}}
}

type {{$dao}} struct {
    DbBaseDao
}

func New{{$dao}}(v ...interface{}) *{{$dao}} {
    this := new({{$dao}})
    this.UpdateEngine(v...)
    return this
}

{{$pl := len .PrimaryKeys}}
{{if gt $pl 0}}
func (this *{{$dao}})Get({{genParams $table .PrimaryKeys true}}) (ret []{{$tb}}, err error) {
    ret = make([]{{$tb}},0)
    err = this.getByPrimaryKey({{genParams $table .PrimaryKeys false}}, 0, 0).Find(&ret)
    return
}
func (this *{{$dao}})GetLimit({{genParams $table .PrimaryKeys true}}, pn, rn int) (ret []{{$tb}}, err error) {
    ret = make([]{{$tb}},0)
    err = this.getByPrimaryKey({{genParams $table .PrimaryKeys false}}, pn, rn).Find(&ret)
    return
}
func (this *{{$dao}})GetCount({{genParams $table .PrimaryKeys true}}) (ret int64, err error) {
    ret, err = this.getByPrimaryKey({{genParams $table .PrimaryKeys false}}, 0, 0).Count(new({{$tb}}))
    return
}
func (this *{{$dao}})getByPrimaryKey({{genParams $table .PrimaryKeys true}}, pn, rn int) (session *xorm.Session) {
    session = this.GetSession()
    {{range .PrimaryKeys}}
        {{$p := Mapper .}}
        this.buildQuery(session, m{{$p}}, "{{.}}")
    {{end}}
    if rn > 0 {
        session = session.Limit(rn, pn)
    }
    return
}

{{end}}

{{range .Indexes}}
func (this *{{$dao}})Get{{getMethodName .Name}}({{genParams $table .Cols true}}) (ret []{{$tb}}, err error) {
    ret = make([]{{$tb}},0)
    err = this.get{{getMethodName .Name}}WithParams({{genParams $table .Cols false}}, 0, 0).Find(&ret)
    return
}
func (this *{{$dao}})Get{{getMethodName .Name}}Count({{genParams $table .Cols true}}) (ret int64, err error) {
    ret, err = this.get{{getMethodName .Name}}WithParams({{genParams $table .Cols false}}, 0, 0).Count(new({{$tb}}))
    return
}
func (this *{{$dao}})Get{{getMethodName .Name}}Limit({{genParams $table .Cols true}}, pn,rn int) (ret []{{$tb}}, err error) {
    ret = make([]{{$tb}},0)
    err = this.get{{getMethodName .Name}}WithParams({{genParams $table .Cols false}}, pn, rn).Find(&ret)
    return
}
func (this *{{$dao}})get{{getMethodName .Name}}WithParams({{genParams $table .Cols true}},pn,rn int) (session *xorm.Session) {
    session = this.GetSession()
    {{range .Cols}}
        {{$p := Mapper .}}
        this.buildQuery(session, m{{$p}}, "{{.}}")
    {{end}}
    if rn > 0 {
        session = session.Limit(rn, pn)
    }
    return
}
{{end}}

{{end}}
