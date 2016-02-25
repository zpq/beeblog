package controllers

import (
	//	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	//	this.Data["isLogin"] = checkAccout(this.Ctx)
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	post := this.Input()
	uname := post.Get("uname")
	pwd := post.Get("pwd")
	autologin := post.Get("autologin") == "on"

	if beego.AppConfig.String("uname") == uname && beego.AppConfig.String("pwd") == pwd {
		maxAge := 0
		if autologin {
			maxAge = 1<<31 - 1
		} else {
			maxAge = 36000
		}
		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
		//		this.Ctx.WriteString(fmt.Sprint(this.Input()))
	}
	this.Redirect("/", 301)

	return
}

func checkAccout(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	uname, pwd := ck.Value, ck.Value
	return beego.AppConfig.String("uname") == uname && beego.AppConfig.String("pwd") == pwd
}
