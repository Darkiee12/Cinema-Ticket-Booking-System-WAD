package middleware

import (
	"cinema/component/appctx"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func AllowCORS(_ appctx.AppContext) gin.HandlerFunc {
	cfg := cors.DefaultConfig()
	cfg.AllowAllOrigins = true
	cfg.AllowCredentials = true
	return cors.New(cfg)
}
