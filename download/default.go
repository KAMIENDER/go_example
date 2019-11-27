package download

import (
	"encoding/json"
	"fmt"
	Models "test/models"

	"github.com/astaxie/beego"
)

type DownloadController struct {
	beego.Controller
}

type FileInfo struct {
	Name string `json:"name"`
	Date string `json:"date"`
	Size string `json:"size"`
}

type Filelist struct {
	User  string
	Files FileInfo
	Check string
}

func (this *DownloadController) Post() {
	var Files Filelist
	data := this.Ctx.Input.RequestBody

	err := json.Unmarshal(data, &Files)
	// 消息开头要大写！！！！，否则gg
	if err != nil {
		fmt.Println("json.Unmarshal is error: ", err.Error())
	}

	if Models.Check_md5(Files.User, Files.Check) == true {
		fmt.Println(Files.Files.Name)
		this.Ctx.Output.Download("H:\\go\\mygo\\src\\test\\static\\" + Files.User + "\\" + Files.Files.Name)
	} else {
		this.Ctx.WriteString("false")
	}

	// this.ServeJSON()这一句会多返回一个null
}
