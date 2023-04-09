package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//多文件上传
	r := gin.Default()

	// 处理跨域请求
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.POST("/upload_audio", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			log.Fatal(err)
		}
		//通过字段名映射
		f := form.File["file"]
		//for range遍历文件
		for _, file := range f {
			fmt.Println(file.Filename)
			c.SaveUploadedFile(file, "./"+file.Filename)
			c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s"+file.Filename))
			c.File("./" + file.Filename)
		}
	})
	r.Run(":8081")
}
