package server

import (
	"pdd/server/controller"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	v1 := router.Group("/v1", gin.BasicAuth(gin.Accounts{
		"pdd": "xxx",
	}))

	v1.GET("/config", controller.ConfigController)
	v1.GET("/coupon", controller.CouponController)
	v1.GET("/pdd", controller.PddController)

	router.Run(":3000")
}
