package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

func AddTrace(tracer trace.Tracer) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "tracer", tracer))
		c.Next()
	}
}
