package main

import (
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	//读取配置
	mhost := beego.AppConfig.String("db.host")
	muser := beego.AppConfig.String("db.user")
	mpwd  := beego.AppConfig.String("db.password")
	mport := beego.AppConfig.String("db.name")
	mpre  := beego.AppConfig.String("db.prefix")
	mtz   := beego.AppConfig.String("db.timezone")	
	//打印
	fmt.Println("app.conf:",mhost,muser,mpwd,mport,mpre,mtz,"\n")

	//读取配置
	rhost := beego.AppConfig.String("redis.host")
	rs    := beego.AppConfig.String("sessionon")
	//打印
	fmt.Println("app.conf:",rhost,rs,"\n")

	//修改配置
	beego.AppConfig.Set("db.host","localhost")
	mhost = beego.AppConfig.String("db.host")
	fmt.Println("set app.conf",mhost,"\n")

	//加载配置，类似include "app2.conf" 
	beego.LoadAppConfig("ini", "conf/app2.conf")
	fmt.Println("app2.conf:",beego.AppConfig.String("db.host2"),"\n")

	//读取系统配置
	fmt.Println("system conf:",beego.BConfig.AppName,"\n")
}