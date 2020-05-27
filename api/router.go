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
		Player *player.Player
	}

	BuyPayload struct {
		Name  string `json:"name" binding:"required"`
		Count int    `json:"count"`
	}
)

func Serve(p *player.Player) {
	s := service{
		Player: p,
	}

	r := gin.Default()

	v1 := r.Group("/api/v1")

	v1.GET("/player", s.GetPlayer)
	v1.POST("/buy", s.Buy)

	r.Run(":4333")
}

func (s *service) GetPlayer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"player": s.Player,
	})
}

func (s *service) Buy(c *gin.Context) {
	var payload BuyPayload

	// default count to 1
	payload.Count = 1

	c.BindJSON(&payload)

	err := s.Player.Add(payload.Name, payload.Count)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("%s does not exist", payload.Name),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"player": s.Player,
		})
	}
}
