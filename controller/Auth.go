package controller

import (
	"fmt"
	"gin-skeleton/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)
type UserInfo struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
func Login(c *gin.Context)  {
	// 用户发送用户名和密码过来
	email ,ok := c.GetPostForm("email")
	password ,ok := c.GetPostForm("password")
	if !ok  {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}
	// 校验用户名和密码是否正确
	fmt.Println(email+password)
	if email == "noah@zixi.net" && password == "381479" {
		// 生成Token
		tokenString, _ := middleware.GenToken(111)
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": gin.H{"token": tokenString},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 2002,
		"msg":  "鉴权失败",
	})
	return

}