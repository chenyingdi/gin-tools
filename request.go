package ginTools

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type JsonMap map[string]interface{}

func (j JsonMap) Get(key string) interface{} {
	return j[key]
}

func (j JsonMap) GetString(key string) string {
	val, ok := j[key].(string)
	if !ok {
		return ""
	}
	return val
}

func (j JsonMap) GetInt(key string) int {
	return int(j.GetInt64(key))
}

func (j JsonMap) GetInt8(key string) int8 {
	return int8(j.GetInt64(key))
}

func (j JsonMap) GetInt16(key string) int16 {
	return int16(j.GetInt64(key))
}

func (j JsonMap) GetInt32(key string) int32 {
	return int32(j.GetInt64(key))
}

func (j JsonMap) GetInt64(key string) int64 {
	switch v := j[key].(type) {
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case uint:
		return int64(v)
	case uint8:
		return int64(v)
	case uint16:
		return int64(v)
	case uint32:
		return int64(v)
	case uint64:
		return int64(v)
	case float32:
		return int64(v)
	case float64:
		return int64(v)
	case bool:
		if v {
			return 1
		} else {
			return 0
		}
	case string:
		val, _ := strconv.Atoi(v)
		return int64(val)

	default:
		return 0
	}
}

func (j JsonMap) GetUint(key string) uint {
	return uint(j.GetUint64(key))
}

func (j JsonMap) GetUint8(key string) uint8 {
	return uint8(j.GetUint64(key))
}

func (j JsonMap) GetUint16(key string) uint16 {
	return uint16(j.GetUint64(key))
}

func (j JsonMap) GetUint32(key string) uint32 {
	return uint32(j.GetUint64(key))
}

func (j JsonMap) GetUint64(key string) uint64 {
	switch v := j[key].(type) {
	case int:
		return uint64(v)
	case int8:
		return uint64(v)
	case int16:
		return uint64(v)
	case int32:
		return uint64(v)
	case int64:
		return uint64(v)
	case uint:
		return uint64(v)
	case uint8:
		return uint64(v)
	case uint16:
		return uint64(v)
	case uint32:
		return uint64(v)
	case uint64:
		return v
	case float32:
		return uint64(v)
	case float64:
		return uint64(v)
	case bool:
		if v {
			return 1
		} else {
			return 0
		}
	case string:
		val, _ := strconv.Atoi(v)
		return uint64(val)

	default:
		return 0
	}
}

func (j JsonMap) GetBool(key string) bool {
	val, ok := j[key].(bool)
	if !ok {
		return false
	}
	return val
}

func (j JsonMap) GetFloat32(key string) float32 {
	switch v := j[key].(type) {
	case string:
		val, _ := strconv.ParseFloat(v, 32)
		return float32(val)
	default:
		return float32(j.GetFloat64(key))
	}
}

func (j JsonMap) GetFloat64(key string) float64 {
	switch v := j[key].(type) {
	case string:
		val, _ := strconv.ParseFloat(v, 64)
		return val
	case float64:
		return v
	case float32:
		return float64(v)
	case int:
		return float64(v)
	case int8:
		return float64(v)
	case int16:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case uint:
		return float64(v)
	case uint8:
		return float64(v)
	case uint16:
		return float64(v)
	case uint32:
		return float64(v)
	case uint64:
		return float64(v)
	default:
		return 0
	}
}

func (j JsonMap) ScanReq(ctx *gin.Context) error {
	return ctx.ShouldBindJSON(&j)
}
