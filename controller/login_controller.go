// login_controller
package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func recover(timestamp string, claim string, signed string) string {
	return "kk"
}
func Login(c *gin.Context) {
	address, timestamp, claim, signed := c.PostForm("address"), c.PostForm("timestamp"), c.PostForm("claim"), c.PostForm("signed")
	recaddress := recover(timestamp, claim, signed)
	if recaddress != address {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "saved!",
		})
		return
	}
	session := sessions.Default(c)
	session.Set("claim", claim)
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "saved!",
	})

}

//https://dev.to/maxwellhertz/tutorial-integrate-gin-with-cabsin-56m0
