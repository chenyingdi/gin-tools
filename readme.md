## 1. 快速发送json响应

```go
package response

import (
    "github.com/chenyingdi/ginTools"
    "github.com/gin-gonic/gin"
)

func example() {
    resp := ginTools.NewJsonResp(ctx *gin.Context) 

    resp.JSONP(
        http.StatusOK,  // code  int
        "success!",     // msg   string
        nil,            // data  interface{}
    )
    
    resp.Success(data interface{})   // code  200
                                     // msg   操作成功！   
                                     // data  为所传参数
    
    resp.BadRequest(data interface{}) // code 400！
                                      // msg  错误的请
    
    resp.UnAuthorized(data interface{}) // code 401
                                        // msg  未授权！！
    
    resp.Forbidden(data interface{})    // code 403
                                        // msg 禁止访问！
    
    resp.TooManyRequest(data interface{}) // code 429
                                          // msg 请求次数过多！
}
```

## 2. 快速从request提取json数据

```go
package request

import (
    "github.com/chenyingdi/gin-tools"
    "github.com/gin-gonic/gin"
)

func example(ctx *gin.Context) {
    req := make(ginTools.JsonMap)
    
    _ := req.ScanReq(ctx)

    req.GetString("username") 
}
```

## 3. 快速构建Restful路由
**Restful结构：**
```go
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

```
**使用：**
```go
package router

import (
    "github.com/chenyingdi/gin-tools"
    "github.com/gin-gonic/gin"
)

func Router(eng *gin.Engine){
    NewRestful(eng.Group("/user")).
        REST(map[string]gin.HandlerFunc{
            "GET": func(ctx *gin.Context) {
                // do something...
            },    
    })
}





```