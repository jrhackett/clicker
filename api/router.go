package api

import (
	"clicker/internal/player"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	Action struct {
		Route  string
		Method string
		Action func(c *gin.Context)
	}

	service struct {
		player *player.Player
	}

	BuyPayload struct {
		Name  string `json:"name" form:"name" binding:"required"`
		Count int    `json:"count" form:"count"`
	}
)

func Serve(p *player.Player) {
	s := service{
		player: p,
	}

	r := gin.Default()

	v1 := r.Group("/api/v1")

	v1.GET("/player", s.Player)
	v1.GET("/cost", s.Cost)
	v1.POST("/buy", s.Buy)

	r.Run(":4333")
}

func (s *service) Player(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"player": s.player,
	})
}

func (s *service) Cost(c *gin.Context) {
	var payload BuyPayload

	// default count to 1
	payload.Count = 1

	c.Bind(&payload)

	cost, err := s.player.Cost(payload.Name, payload.Count)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("%s does not exist", payload.Name),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"cost": cost,
		})
	}
}

func (s *service) Buy(c *gin.Context) {
	var payload BuyPayload

	// default count to 1
	payload.Count = 1

	c.BindJSON(&payload)

	err := s.player.Add(payload.Name, payload.Count)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("%s does not exist", payload.Name),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"player": s.player,
		})
	}
}
