package main

import (
	"demo/Global"
	"demo/MiddleWare/Cross"
	"demo/MiddleWare/JWT"
	"demo/Connect/MySQL"
	response "demo/Response"
	"demo/Structs"
	"demo/Utils"
	"demo/Viper"
	"fmt"

	"demo/Email"
	"demo/Connect/Redis"

	"github.com/gin-gonic/gin"
)

func init() {
	Viper.Config()
	MySQL.Mysql()
	Redis.Redis()
}

func main() {
	r := gin.Default()
	WOToken := r.Group("/WO")
	WOToken.Use(Cross.Cors())
	{
		WOToken.POST("/GetToken", func(c *gin.Context) {

			var user Structs.User

			Username := c.PostForm("Username")
			Password := c.PostForm("Password")

			result := Global.DB.Where("username = ?", Username).First(&user)

			TokenString, err := Utils.ReleaseToken(user)

			if result.Error != nil || user.Password != Password {
				response.FailWithInfor(response.NoAuth, "密码错误或者查询接口失效", c)
			} else if err != nil {
				response.FailWithInfor(response.ServerErr, "Token生成失败请联系管理员", c)
			} else {
				response.Result(response.Success, fmt.Sprintf("Token:%s", TokenString), fmt.Sprintf("登录成功:%s", Username), c)
			}

		})

		//注册函数
		WOToken.POST("/user/regist", Email.Regist)
		//邮件发送
		WOToken.POST("/sendEmail", Email.SendEmailCode)
	}

	WToken := r.Group("/W")
	WToken.Use(JWT.Auth())
	{
		WToken.POST("/ParaseToken", func(c *gin.Context) {

			UID := c.MustGet("UID").(uint)
			Password := c.MustGet("Password").(string)
			Username := c.MustGet("Username").(string)

			//临时封装
			data := map[string]interface{}{
				"UID":      UID,
				"Password": Password,
				"Username": Username,
			}
			response.SucceedWithDetailed(data, "解析后的用户信息为", c)

		})

	}

	r.Run(":8000")

}
