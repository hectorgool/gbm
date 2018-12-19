package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hectorgool/gbm/common"
	"github.com/hectorgool/gbm/elasticsearch"
	"net/http"
)

func setupRouter() *gin.Engine {

	r := gin.Default()
	r.Use(common.CORSMiddleware())

	r.GET("/", func(c *gin.Context) {
		
		result, err := elasticsearch.Ping()
		common.CheckError(err)
		c.JSON(200, gin.H{"data": result})

	})

	// /geolocation/666-xxx
	r.GET("/geolocation/:vehicleid", func(c *gin.Context) {

		vehicleid := c.Param("vehicleid")

		latitude, longitud, err := common.JSONToStruct()
		common.CheckError(err)

		if err := elasticsearch.CreateDocument(vehicleid, latitude, longitud); err != nil {
			common.CheckError(err)
		}

		c.JSON(http.StatusOK, gin.H{ 
			"latitude": latitude,
			"longitud": longitud,
		})

	})

	r.GET("/record", func(c *gin.Context) {

		q := c.DefaultQuery("q", "")
		result, err := elasticsearch.Search(q)
		common.CheckError(err)
		c.JSON(200, gin.H{"data": result})

	})

	return r

}

func main() {
	r := setupRouter()
	r.Run(common.PORT) // listen and serve on 0.0.0.0:8080
}