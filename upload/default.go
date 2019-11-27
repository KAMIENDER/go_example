package upload

import (
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego"
	Models "test/models"
)

type UploadController struct {
	beego.Controller
}

type GetUserController struct {
	beego.Controller
}

type UserInfo struct {
	Name string
	Key string
	Check string
}

var user UserInfo

func (this *GetUserController) Post() {
	data := this.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &user)
	// 消息开头要大写！！！！，否则gg
	if err != nil {
		fmt.Println("json.Unmarshal is error: ",err.Error())
	}

	if Models.Check_md5(user.Name, user.Check) == true{
		this.Ctx.WriteString(Models.MD5(user.Check))
	} else {
		this.Ctx.WriteString(Models.MD5("false"))
	}
	fmt.Println(user.Name)
	fmt.Println(user.Key)
	
}

func (this *UploadController) Post() {
	// 根据字段名获取表单文件
	r := this.Ctx.Request
	// fmt.Print(r.FormValue("filename"))
	fmt.Println(r.FormValue("file"))
	f, h, _ := this.GetFile("file") //获取上传的文件
	fmt.Println(user.Name)
	fmt.Println(user.Key)
	path := "H:\\go\\mygo\\src\\test\\static\\" + user.Name + "\\" + h.Filename
	defer f.Close() //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	err := this.SaveToFile("file", path)
	if err != nil {
			fmt.Println(err)
			return
		}
		this.Ctx.WriteString("ok")
}
