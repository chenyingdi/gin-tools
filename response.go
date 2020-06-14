package ginTools

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JsonResp struct {
	ctx  *gin.Context
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewJsonResp(ctx *gin.Context) *JsonResp {
	return &JsonResp{ctx: ctx}
}

func (j *JsonResp) JSONP(code int, msg string, data interface{}) {
	j.Code = code
	j.Msg = msg
	j.Data = data
	j.ctx.JSONP(http.StatusOK, j)
	j.ctx.Abort()
}

func (j *JsonResp) Success(data interface{}) {
	j.JSONP(http.StatusOK, "操作成功！", data)
}

func (j *JsonResp) ServerError(data interface{}) {
	j.JSONP(http.StatusInternalServerError, "数据错误！", data)
}

func (j *JsonResp) BadRequest(data interface{}) {
	j.JSONP(http.StatusBadRequest, "错误的请求！", data)
}

func (j *JsonResp) UnAuthorized(data interface{}) {
	j.JSONP(http.StatusUnauthorized, "未授权！", data)
}

func (j *JsonResp) Forbidden(data interface{}) {
	j.JSONP(http.StatusForbidden, "禁止访问！", data)
}

func (j *JsonResp) TooManyRequest(data interface{}) {
	j.JSONP(http.StatusTooManyRequests, "请求次数过多！", data)
}
