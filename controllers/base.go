package controllers

import (
	"github.com/astaxie/beego"
)
/**
** 这是控制器的基类，建议所有控制器扩展这个基类以实现一些常用功能的复用。
**/
type BaseController struct{
	beego.Controller
}

