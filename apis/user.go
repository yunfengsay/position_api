package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"position_api/models"
)

func WXLogin(c *gin.Context) {
	openId := c.Request.FormValue("open_id")
	if user, err := models.WXLogin(openId); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(user)
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"user": user,
		})
	}
}

func CreateWXUser(c *gin.Context) {}
