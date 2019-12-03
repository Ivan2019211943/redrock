func Register(c *gin.Context) {
//获取用户名、密码
name := c.Request.FormValue("Name")
passwd := c.Request.FormValue("Passwd")
//判断用户是否存在
//不存在创建用户，保存密码与用户名
Bool := Func.IsExist(name)
if Bool {
//注册状态
State["state"] = 1
State["text"] = "此用户已存在！"
} else {
//用户不存在即添加用户
AddStruct(name, passwd)
State["state"] = 1
State["text"] = "注册成功！"
}
//登录
func Login(c *gin.Context) {
name := c.Request.FormValue("Name")
passwd := c.Request.FormValue("Passwd")
if Bool_Pwd {
State["state"] = 1
State["text"] = "登录成功！"
} else {
State["state"] = 0
State["text"] = "密码错误！"
}
