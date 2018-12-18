package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hectorgool/mvp_gbm/common"
	"github.com/hectorgool/mvp_gbm/elasticsearch"
)

func main() {

	r := gin.Default()
	r.Use(common.CORSMiddleware())

	r.GET("/", func(c *gin.Context) {
		
		result, err := elasticsearch.Ping()
		common.CheckError(err)
		c.JSON(200, gin.H{"data": result})

	})

	r.GET("/geolocation", func(c *gin.Context) {
		
		latitude, longitud, err := common.JSONToStruct()
		common.CheckError(err)

		if err := elasticsearch.CreateDocument(latitude, longitud); err != nil {
			common.CheckError(err)
		}

		c.JSON(200, gin.H{ 
			"latitude": latitude,
			"longitud": longitud,
		})

	})

	r.GET("/record", func(c *gin.Context) {
		
		result, err := elasticsearch.Search()
		common.CheckError(err)
		c.JSON(200, gin.H{"data": result})

	})

	r.Run(common.PORT) // listen and serve on 0.0.0.0:8080

}