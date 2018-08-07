package main

import (
	_ "beego-archetype/routers"
	"github.com/astaxie/beego"

)

func main() {
	beego.Debug("This is the entry of beego archetype!")
	
	beego.Run()
}

