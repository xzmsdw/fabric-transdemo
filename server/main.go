package main

import (
	"application/api/v1"
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个gin对象：r
	r := gin.Default()

	//r.Use(ginsession.New()) //使用session插件

	//  添加静态文件处理
	//r.StaticFile("/", "dist/index.html")
	//r.StaticFile("/index", "dist/index.html")
	//r.Static("/static", "dist/static")

	// 设置路由规则： 当浏览器 请求 http://localhost:8282/ping GET 使用func(c *gin.Context)处理
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("CreateCrops", v1.CreateCrops)
	r.POST("RecordCropsGrow", v1.RecordCropsGrow)
	r.GET("QueryCropsById", v1.QueryCropsById)
	r.GET("QueryCropsProcessByCropsId", v1.QueryCropsProcessByCropsId)
	r.POST("CreateTransport", v1.CreateTransport)
	r.GET("QueryTransportById", v1.QueryTransportById)
	r.GET("QueryTransportByCropsId", v1.QueryTransportByCropsId)
	r.POST("CreateMachining", v1.CreateMachining)
	r.GET("QueryMachiningById", v1.QueryMachiningById)
	r.GET("QueryMachiningByCropsId", v1.QueryMachiningByCropsId)
	r.POST("/CreateSelling", v1.CreateSelling)
	r.GET("/QueryBySellID", v1.QueryBySellID)
	r.GET("/QueryBySellerID", v1.QueryBySellerID)
	r.GET("/QueryByBuyerID", v1.QueryByBuyerID)
	r.POST("/CreateCommodity", v1.CreateCommodity)
	r.POST("/ChangeCommodity", v1.ChangeCommodity)
	r.GET("/DelCommodity", v1.DelCommodity)
	r.GET("/QueryAllBySellerID", v1.QueryAllBySellerID)
	r.POST("/CreateBalance", v1.CreateBalance)
	r.GET("/QueryBalance", v1.QueryBalance)
	r.POST("/ChangeBalance", v1.ChangeBalance)
	r.POST("/CreateAccount",v1.CreateAccount)
	r.POST("/ChangePass",v1.ChangePass)

	_ = r.Run(":8282") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
