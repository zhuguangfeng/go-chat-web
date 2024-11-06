package jwt

import "github.com/gin-gonic/gin"

type JwtHandler interface {
	//清除token
	ClearToken(ctx *gin.Context) error
	ExtractToken(ctx *gin.Context) string
	SetLoginToken(ctx *gin.Context, uid uint) error
	SetJwtToken(ctx *gin.Context, uid uint, ssid string) error
	CheckSession(ctx *gin.Context, ssid string) error
}
