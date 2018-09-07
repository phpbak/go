package main

import (
	"fmt"
	"gotest/models"
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

}

func main() {
	//读取配置
	dbhost := beego.AppConfig.String("db.host")
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")
	dbport := beego.AppConfig.String("db.port")
	mpre := beego.AppConfig.String("db.prefix")
	mtz := beego.AppConfig.String("db.timezone")

	//打印
	fmt.Println("app.conf:", dbhost, dbuser, dbpassword, dbport, mpre, mtz, "\n")

	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	if mtz != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(mtz)
	}

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dsn)

	orm.RegisterModel(new(models.User))

	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	//new
	user := new(models.User)
	//查询
	user1 := models.User{Id: 1}
	fmt.Println(o.Read(&user1))
	//添加
	user.LoginName = "get12347"

	fmt.Println(o.Insert(user))

	//修改
	user.LoginName = "get123456"
	user.Id = 9
	fmt.Println(o.Update(user))
	//删除
	user.Id = 11
	fmt.Println(o.Delete(user))
}
