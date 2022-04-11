package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type HelloController struct {
	beego.Controller
}

// Get Get方式
func (c *HelloController) Get() {
	c.Ctx.WriteString("hello Luke!")
}

// Post Post方式
func (c *HelloController) Post(){
	c.Ctx.WriteString("hello Luke for post method!")

	name := c.GetString("name")
	pass := c.GetString("pass")

	fmt.Println("name:", name, "pass:", pass)
	c.Ctx.WriteString(name+" "+pass)
	c.Ctx.WriteString("index post")
}
