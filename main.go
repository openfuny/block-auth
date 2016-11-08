package main

import (
	_ "block-auth/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
