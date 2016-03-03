package main

import (
	"github.com/astaxie/beego"
	"github.com/kubego/beego-blog/g"
	_ "github.com/kubego/beego-blog/routers"
)

func main() {
	g.InitEnv()
	beego.Run()
}
