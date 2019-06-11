package httpapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var VersionList []map[string]gin.HandlerFunc

func init()  {
	v1 := make(map[string]gin.HandlerFunc)
	v2 := make(map[string]gin.HandlerFunc)

	v1["get"] = func(context *gin.Context) {
		context.String(http.StatusOK, "v1")
	}
	v2["get"] = func(context *gin.Context) {
		context.String(http.StatusOK, "v2")
	}

	VersionList = append(VersionList, v2, v1)
}


func FindController(ctx *gin.Context) gin.HandlerFunc{
	var handler gin.HandlerFunc
	for _, v := range VersionList {
		if h, find := v[ctx.Request.RequestURI]; find {
			return h
		}
	}
	return handler
}




