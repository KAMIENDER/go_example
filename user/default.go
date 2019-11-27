package User

import (
	"encoding/json"
	"fmt"
	Models "test/models"

	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
	// "io"
)

type UserController struct {
	beego.Controller
}

type UserInfo struct {
	Name string
	Key  string
}

func (this *UserController) Post() {
	var temp UserInfo
	var check Models.InfoBack
	var user Models.MysqlUser

	data := this.Ctx.Input.RequestBody
	// var pass string
	// pass = MD5("19990519")

	err := json.Unmarshal(data, &temp)
	// 消息开头要大写！！！！，否则gg
	if err != nil {
		fmt.Println("json.Unmarshal is error: ", err.Error())
	}

	var num int64
	num, user, err = Models.Lookup(temp.Name)

	// 每次的check都是上次的MD5，用于加密验证
	if err == nil && num == 1 && user.Key == temp.Key {
		check.Status = "ok"
		check.Check = Models.MD5(temp.Key)
		check.Name = temp.Name

		err = Models.Change_check(user.Name, Models.MD5(user.Key))
		if err != nil {
			check.Status = "false"
		}
	} else {
		check.Status = "false"
	}
	this.Data["json"] = &check
	this.ServeJSON()

	// this.ServeJSON()这一句会多返回一个null
}
