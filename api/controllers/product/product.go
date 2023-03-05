package product

import "github.com/gin-gonic/gin"

type Controller struct{}

func (p Controller) Status(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Pong",
	})
}
