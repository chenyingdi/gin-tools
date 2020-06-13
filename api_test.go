package ginTools

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestRestful_REST(t *testing.T) {
	eng := gin.Default()

	NewRestful(eng.Group("/test")).
		REST(map[string]gin.HandlerFunc{
			"GET": func(ctx *gin.Context) {
				ctx.Writer.WriteString("Hello")
			},
		})

	eng.Run()
}
