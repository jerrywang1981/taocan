package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jerrywang1981/taocan/services"
)

func handlePredictTaoCan(c *gin.Context) {
	input := &services.TaoCanInput{}
	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := services.Send(input)
	c.JSON(http.StatusOK, res)
}

func LoadWATaoCan(r *gin.RouterGroup) {
	g := r.Group("/taocan")
	{
		g.GET("/", handlePredictTaoCan)
	}
}
