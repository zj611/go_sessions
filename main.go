package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

var (
	cookieNameForSessionID ="mycookesinahduixu"
	sess = sessions.New(sessions.Config{Cookie:cookieNameForSessionID})
)


func secret(ctx iris.Context){
	if auth, err := sess.Start(ctx).GetBoolean("authennticated"); !auth{
		fmt.Println(auth, err)
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	ctx.WriteString("the cake is a lie!")
}

func login(c iris.Context){
	session := sess.Start(c)

	session.Set("authennticated", true)//保存在服务端内存里

	c.WriteString("logging")
}

func loginout(c iris.Context){
	session := sess.Start(c)

	// 撤销用户身份验证
	session.Set("authennticated", false)
	c.WriteString("loginout")
}
func main(){
	app	:= iris.New()

	app.Get("/secret", secret)
	app.Get("/login", login)
	app.Get("/loginout", loginout)
	app.Run(iris.Addr(":8080"))

}

