package controllers

import "C"
import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

type InputData struct {
	Name string
	Pass string
}


// Get Get方式
func (c *LoginController) Get() {
	c.TplName = "login.tpl"
}

//// Post Post方式
//func (c *LoginController) Post(){
//
//	var Input InputData
//
//	name := c.GetString("name")
//	pass := c.GetString("pass")
//
//	if name == ""  && pass == "" {
//		data := c.Ctx.Input.RequestBody
//		err := json.Unmarshal(data, &Input)
//		if err != nil {
//			b, _ := json.MarshalIndent(err, "", "     ")
//			c.Ctx.WriteString(string(b))
//		}
//		name = Input.Name
//		pass = Input.Pass
//	}
//
//	//user, error := models.FindUsersByUsername(name)
//	//
//	//if error != nil {
//	//	fmt.Println("查询用户失败!错误详情\n", error)
//	//}else {
//	//	fmt.Println(user)
//	//	c.Ctx.WriteString("User Login!")
//	//	//c.Data["json"] = map[string]interface{}{
//	//	//	"code": 200,
//	//	//	"data": user,
//	//	//}
//	//	//c.ServeJSON()
//	//}
//
//	//fmt.Println("name:", name,"Pass: ", pass)
//	//
//	//c.Ctx.WriteString(name + pass)
//	//c.Ctx.WriteString("aaa")
//}
