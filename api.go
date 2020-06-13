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
		case "GET", "Get", "get":
			r.group.GET("", v)
		case "POST", "Post", "post":
			r.group.POST("", v)
		case "PUT", "Put", "put":
			r.group.PUT("", v)
		case "DELETE", "Delete", "delete":
			r.group.DELETE("", v)
		case "OPTIONS", "Options", "options":
			r.group.OPTIONS("", v)
		case "PATCH", "Patch", "patch":
			r.group.PATCH("", v)
		}
	}
}
