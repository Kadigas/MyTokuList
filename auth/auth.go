package auth

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func BasicAuth(c *gin.Context) {
	user, password, hasAuth := c.Request.BasicAuth()
	if hasAuth && ((user == "admin" && password == "password") || (user == "editor" && password == "secret")) {
		log.WithFields(log.Fields{
			"user": user,
		}).Info("User authenticated")
	} else {
		c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		c.AbortWithStatus(401)
		return
	}
}
