package Models

import (
	"crypto/md5"
	"encoding/hex"
	// "fmt"
	"github.com/astaxie/beego/orm"
)

type MysqlUser struct {
	Id    int64  `orm:"pk;column(Id)"`
	Name  string `orm:"column(Name)"`
	Key   string `orm:"column(Key)"`
	Check string `orm:"column(Check)"`
}

type InfoBack struct {
	Check  string `json:"check"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

// 下面这个函数用于寻找对应的user信息
func Lookup(user string) (int64, MysqlUser , error){
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable("userinfo")
	var users []MysqlUser
	num, err := qs.Filter("Name", user).All(&users)
	if err != nil {
		return -1, *(new(MysqlUser)),err
	}
	return num,users[0], err
	
}

// 下面这个函数用于修改对应user的check,key表示的是加密前的东西
// 验证过程：收到前端发过来的MD5，后端自身根据原有check字段产生新的MD5，进行对比
// 相同则验证成功，然后新的MD5赋值给check，并且返回，
func Change_check(user string, key string) error {
	_, err := orm.NewOrm().QueryTable("userinfo").Filter("Name", user).Update(orm.Params{
		"Check": key,
	})
	return err
}

func Check_md5(user string, md5 string) bool {
	var userg MysqlUser
	var num int64
	var err error
	num, userg, err = Lookup(user)
	if err == nil && MD5(userg.Check) == md5 && num == 1{
		err = Change_check(user,md5)
		if err == nil{
			return true
		}
	}else {
		return false
	}
	return false
}

func (u *MysqlUser) TableName() string {
	return "userinfo"
}

func init() {
	orm.RegisterModel(new(MysqlUser))
}
