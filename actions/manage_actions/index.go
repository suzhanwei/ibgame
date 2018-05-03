package manage_actions

import "github.com/astaxie/beego"

type Index struct {
	beego.Controller
}

func (this *Index) Get() {
	this.Data["player"] = models.GetAll()
	this.Layout = "layout.tpl"
	this.TplName = "index.tpl"
}
