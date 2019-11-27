package main

import (
	_ "test/routers"
	// "fmt"
	// Models "test/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 这个函数很重要取决于
func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/user?charset=utf8")
}

func main() {
	// orm.Debug = true
	// o := orm.NewOrm()
	// qs := o.QueryTable("userinfo")
	// fmt.Println(qs)
	beego.Run()
}
