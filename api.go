package ginTools

import "github.com/gin-gonic/gin"

type Restful struct {
	group *gin.RouterGroup
}

func NewRestful(group *gin.RouterGroup) *Restful {
	return &Restful{group: group}
}

func (r *Restful) REST(funcMap map[string]gin.HandlerFunc) {
	for k, v := range funcMap {
		switch k {
		case "GET":
			r.group.GET("", v)
		case "POST":
			r.group.POST("", v)
		case "PUT":
			r.group.PUT("", v)
		case "DELETE":
			r.group.DELETE("", v)
		case "OPTIONS":
			r.group.OPTIONS("", v)
		case "PATCH":
			r.group.PATCH("", v)
		}
	}
}
