package controllers

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/casbin/casbin"
	"github.com/casbin/beego-orm-adapter"
)

type RoleController struct{
	beego.Controller
}

var e *casbin.Enforcer

func init()  {
	a := beegoormadapter.NewAdapter("mysql", "root:111111@tcp(127.0.0.1:3306)/beego-archetype",true)
	e=casbin.NewEnforcer("conf/rbac_model.conf",a)
	e.LoadPolicy()
}

func (c *RoleController) Test()  {
	beego.Info(e.Enforce("alice","data1","write"))
	beego.Info("ddddddddddd")
//	c.TplName = "index.html"
}