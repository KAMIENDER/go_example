package routers

import (
	animate "test/animate"
	"test/controllers"
	download "test/download"
	upload "test/upload"
	User "test/user"
	"test/wangpan"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &User.UserController{})
	beego.Router("/index", &wangpan.WangpanController{})
	beego.Router("/download", &download.DownloadController{})
	beego.Router("/upload", &upload.UploadController{})
	beego.Router("/upload/user", &upload.GetUserController{})
	beego.Router("/animate", &animate.AnimateController{})
}
