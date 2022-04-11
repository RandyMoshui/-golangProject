package main

import (
	"github.com/astaxie/beego"
	_ "mybeego/routers"
)

func main() {
	// 设置静态目录
	//beego.SetStaticPath("download1", "download2")
	beego.Run(":8888") // 默认8080端口
}

