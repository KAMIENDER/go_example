package wangpan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	Models "test/models"
	"github.com/astaxie/beego"
)

type WangpanController struct {
	beego.Controller
}

type Wangpaninfo struct {
	User string
	Check string
}

// 注意下面这个大小写标签并不是用单引号
type FileInfo struct {
	Name string `json:"name"`
	Date string `json:"date"`
	Size string `json:"size"`
}

type Files struct {
	Data []FileInfo
	Check string `json:"check"`
}

func (this *WangpanController) Post() {
	var temp Wangpaninfo
	var jsonResponse Files
	data := this.Ctx.Input.RequestBody

	err := json.Unmarshal(data, &temp)
	// 消息开头要大写！！！！，否则gg
	if err != nil {
		fmt.Println("json.Unmarshal is error: ", err.Error())
	}
	fmt.Println(temp)
	if Models.Check_md5(temp.User, temp.Check) == true{
		path := "H:\\go\\mygo\\src\\test\\static\\"
		path += temp.User
		fmt.Println(path)
		jsonResponse.Data, err = getDirList(path)
		jsonResponse.Check = temp.Check
		if err != nil {
			fmt.Println(err.Error())
		}
		this.Data["json"] = &jsonResponse
		
	} else {
		jsonResponse.Check = "false"
	}
	fmt.Println(jsonResponse)
	this.ServeJSON()
	// this.ServeJSON()这一句会多返回一个null
}

func getDirList(dirpath string) ([]FileInfo, error) {
	var dirlist []FileInfo
	rd, err := ioutil.ReadDir(dirpath)
	for _, fi := range rd {
		var file FileInfo
		fmt.Println(fi.Name())
		file.Date = fi.ModTime().Format("2006-01-02 03:04:05 PM")
		file.Name = fi.Name()
		file.Size = strconv.FormatInt(fi.Size(), 10) + "B"
		dirlist = append(dirlist, file)
	}
	return dirlist, err
}
