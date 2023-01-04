package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type healthResponse struct {
	Alive bool `json:"alive"`
}

func HealthHandler(path string) func(w http.ResponseWriter, r *http.Request) bool {
	return func(w http.ResponseWriter, r *http.Request) bool {
		return r.URL.Path == path
	}
}

func GinHealthHandler(path string) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler := HealthHandler(path)
		if handler(c.Writer, c.Request) {
			c.JSON(200, healthResponse{
				Alive: true,
			})
			c.Abort()
		}
	}
}
