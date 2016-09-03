package main

import (
	_"routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/wrx", "static")  
	beego.Run()
}
