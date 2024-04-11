package monza_middleware_gin

import (
	"github.com/chinmayrelkar/monza"
	"github.com/gin-gonic/gin"
	"time"
)

const GinService = "gin-middleware"

func GetGinMiddleware(client monza.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		client.Record(
			monza.Event{
				Event:      "request",
				ClientTime: time.Now(),
				Data: map[string]interface{}{
					"method": ctx.Request.Method,
					"path":   ctx.Request.RequestURI,
				},
				ServiceID: GinService,
			},
		)

		start := time.Now()
		ctx.Next()
		end := time.Now()

		client.Record(
			monza.Event{
				Event:      "response",
				ClientTime: time.Now(),
				Data: map[string]interface{}{
					"method":  ctx.Request.Method,
					"path":    ctx.Request.RequestURI,
					"latency": end.Sub(start).Milliseconds(),
					"status":  ctx.Writer.Status(),
				},
				ServiceID: GinService,
			},
		)
	}
}
