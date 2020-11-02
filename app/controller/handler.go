package controller

import (
	"net/http"

	"github.com/YoriDigitalent/KonsepMVC_GO-23/app/model"
	"github.com/gin-gonic/gin"
)

//AddAntrianHandler for add queue
func AddAntrianHandler(c *gin.Context) {
	flag, err := model.AddAntrian()

	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error":  err,
		})
	}
}

//GetAntrianHandler for get all queue
func GetAntrianHandler(c *gin.Context) {
	flag, data, err := model.GetAntrian()

	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
			"data":   data,
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error":  err,
		})
	}
}

//UpdateAntrianHandler for update queue by id
func UpdateAntrianHandler(c *gin.Context) {
	idAntrian := c.Param("idAntrian")
	flag, err := model.UpdateAntrian(idAntrian)

	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error":  err,
		})
	}
}

//DeleteAntrianHandler for delete a queue
func DeleteAntrianHandler(c *gin.Context) {
	idAntrian := c.Param("idAntrian")
	flag, err := model.DeleteAntrian(idAntrian)

	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error":  err,
		})
	}
}

//PageAntrianHandler for html queue
func PageAntrianHandler(c *gin.Context) {
	flag, result, err := model.GetAntrian()
	var currentAntrian map[string]interface{}

	for _, item := range result {
		if item != nil {
			currentAntrian = item
			break
		}
	}

	if flag && len(result) > 0 {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"antrian": currentAntrian["id"],
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error":  err,
		})
	}
}
