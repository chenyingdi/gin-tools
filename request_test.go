package ginTools

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestReq(t *testing.T) {
	eng := gin.Default()

	eng.GET("/", func(ctx *gin.Context) {
		resp := NewJsonResp(ctx)
		j := make(JsonMap)

		err := j.ScanReq(ctx)
		if err != nil {
			resp.ServerError(err.Error())
			return
		}

		resp.Success("Hello " + j.GetString("name"))
	})

	eng.Run()
}
