package api

import "github.com/gin-gonic/gin"

type (
	Action struct {
		Route  string
		Method string
		Action func(c *gin.Context)
	}

	service struct{}
)

func Serve() {
	s := service{}

	actions := []Action{
		{
			Route:  "/player",
			Method: "GET",
			Action: s.GetPlayer,
		},
		{
			Route:  "/click",
			Method: "POST",
			Action: s.Click,
		},
		{
			Route:  "/buy",
			Method: "POST",
			Action: s.Buy,
		},
	}

	r := gin.Default()

	for a := range actions {
		switch a.Method {
		case "GET":
			r.GET(a.Route, a.Action)
		case "POST":
			r.POST(a.Route, a.Action)
		}
	}

	r.Run()
}

func (s *service) GetPlayer(c *gin.Context) {

}

func (s *service) Click(c *gin.Context) {

}

func (s *service) Buy(c *gin.Context) {

}
