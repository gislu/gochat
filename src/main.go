package main

import (
	_"routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/gochat", "static")
	beego.Run()
}
