package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	// 对应表中users字段
	Id int
	Username string
	Password string
	Sex int
	// 外键需要对应
	Role *UserRole `orm:"rel(fk)"`
}

type UserRole struct {
	// 对应User_role表中字段
	Id int
	Names string `orm:"column(role_name)"`
	Descs string `orm:"column(role_desc)"`
	Users []*Users `orm:"reverse(many)"` // 反向一对多 确定关联关系
}

func init()  {
	orm.Debug = true
	orm.RegisterDataBase(
		"default",
		"mysql",
		"root:ubuntu@tcp(192.168.254.128:3306)/laravel?charset=utf8")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterModel(new(Users), new(UserRole))
}

func FindUsersByUsername(Username string)  (*Users, error)  {
	o := orm.NewOrm() // 创建新的orm
	user := &Users{}
	o.QueryTable("Users").Filter("username", Username).One(user)

	//if err == orm.ErrNoRows{
	//	fmt.Println("查询不到该账户")
	//}else if err == orm.ErrMissPK{
	//	fmt.Println("查询不到主键")
	//}else {
	//	fmt.Println(user.Name, user.Username)
	//}

	return user,nil
}

func AllUsers() ([]*Users, error) {
	var u []*Users
	_, err := orm.NewOrm().QueryTable("Users").RelatedSel().All(&u)
	return u, err
}

func FindRoleById(id int) *UserRole {
	r := &UserRole{}

	o := orm.NewOrm()
	o.QueryTable("UserRole").Filter("id", id).One(r)
	o.LoadRelated(r, "Users")
	return r
}