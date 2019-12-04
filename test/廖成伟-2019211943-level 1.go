package routers

import (
	"fmt"
	"net/http"
)

//注册
func Register(c *gin.Context) {
	//获取用户名、密码
	name := c.Request.FormValue("Name")
	password := c.Request.FormValue("Password")
	//判断用户是否存在
	Bool := Func.IsExist(name)
	if Bool {
		//注册状态
		State["state"] = 1
		State["text"] = "此用户已存在！"
	} else {
		//用户不存在即添加用户
		AddStruct(name, password)
		State["state"] = 1
		State["text"] = "注册成功！"
	}
	//登录

	func SendMsg(c *gin.Context){

		username,err:=c.Cookie("username")

		fmt.Println("username"+username)

		if err != nil{

			c.JSON(500,gin.H{"status": http.StatusForbidden,"message":"cookie读取失败"})

			return

		}

		message:=c.PostForm("message")

		if model.SendMesage(username,message){

			c.JSON(200, gin.H{"内容":message,"用户名":username})

		}else {

			c.JSON(403, gin.H{"status": http.StatusForbidden, "message": "发送失败"})

		}

	}



	func GetMsg(c *gin.Context){



		err,msg:=model.GetMesage(10)

		if err !=nil{

			c.JSON(500,gin.H{"status": http.StatusInternalServerError,"message":"数据库读取失败"})

		}

		c.JSON(200,gin.H{"Data":msg})

