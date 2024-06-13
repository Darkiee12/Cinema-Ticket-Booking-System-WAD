package middleware

import (
	"cinema/component/appctx"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func AllowCORS(ac appctx.AppContext) gin.HandlerFunc {
	cfg := cors.DefaultConfig()
	cfg.AllowAllOrigins = true
	return cors.New(cfg)
}
