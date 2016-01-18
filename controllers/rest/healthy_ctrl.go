package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthyController struct{}

func (self *HealthyController) Init(r *gin.Engine) {
	r.GET("/api/healthy", self.healthy)
}

func (self *HealthyController) healthy(c *gin.Context) {
	c.String(http.StatusOK, "ok")
	return;
}