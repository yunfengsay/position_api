package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"position_api/db"
	"position_api/models"
	"time"
)

func AddLike(c *gin.Context) {
	var like = models.Like{1, 1, 1, time.Now()}
	db.Conn.Create(&like)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
}

func TestLike() {
	var like = models.Like{1, 1, 1, time.Now()}
	db.Conn.Create(&like)
	fmt.Println("ok")
}
