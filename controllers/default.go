package controllers

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	//user:=c.GetCurrentUser()

	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.Data["UserId"]=user.FullName
	//c.Data["username"]=user.FullName
	//
	//c.Layout="main_layout.html"
	c.TplName = "index.html"
	//c.Ctx.WriteString("Hello")
	//  c.Redirect(beego.URLFor("ParkingController.All"),302)
}
