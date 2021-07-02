package routers

import (
    "fmt"
    "gin-skeleton/controller"
    "gin-skeleton/pkg/setting"
    "gin-skeleton/service"
    "github.com/gin-gonic/gin"
    "net/http"
)

func InitRouter() *gin.Engine {
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    gin.SetMode(setting.RunMode)
    var loginService service.LoginService = service.StaticLoginService()
    var jwtService service.JWTService = service.JWTAuthService()
    var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)
    r.GET("/test", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "test",
        })
    })
    r.POST("/login", func(ctx *gin.Context) {
        fmt.Println(ctx)
    	token := loginController.Login(ctx)
    	if token != "" {
    		ctx.JSON(http.StatusOK, gin.H{
    			"token": token,
    		})
    	} else {
    		ctx.JSON(http.StatusUnauthorized, nil)
    	}
    })

    return r
}
