package main

import (
	"github.com/YoriDigitalent/KonsepMVC_GO-23/app/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("views/*")

	//router.GET("/", getSomething)

	r.POST("/api/v1/antrian", controller.AddAntrianHandler)
	r.GET("/api/v1/antrian/status", controller.GetAntrianHandler)
	r.PUT("/api/v1/antrian/id/:idAntrian", controller.UpdateAntrianHandler)
	r.DELETE("/api/v1/antrian/id/:idAntrian/delete", controller.DeleteAntrianHandler)
	r.GET("/antrian", controller.PageAntrianHandler)
	r.Run(":5050")
}

/*func getSomehing(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"body1": "Something Success",
	})

	return
}*/
