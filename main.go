package main

import (
    "fmt"
    "gin-skeleton/pkg/setting"
    "gin-skeleton/routers"
    "net/http"
)

func main() {
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	//var loginService service.LoginService = service.StaticLoginService()
	//var jwtService service.JWTService = service.JWTAuthService()
	//var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)
    //fmt.Print(string(setting.HTTPPort))
	//server := gin.New()
	//server.POST("/login", func(ctx *gin.Context) {
	//	token := loginController.Login(ctx)
	//	if token != "" {
	//		ctx.JSON(http.StatusOK, gin.H{
	//			"token": token,
	//		})
	//	} else {
	//		ctx.JSON(http.StatusUnauthorized, nil)
	//	}
	//})
	//server.Run(":" + strconv.Itoa(setting.HTTPPort))

    router := routers.InitRouter()

    s := &http.Server{
        Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
        Handler:        router,
        ReadTimeout:    setting.ReadTimeout,
        WriteTimeout:   setting.WriteTimeout,
        MaxHeaderBytes: 1 << 20,
    }

    s.ListenAndServe()

}
