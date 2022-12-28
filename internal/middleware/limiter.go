package middleware

import (
	"github.com/gin-gonic/gin"
	"go_gin_blog/pkg/app"
	"go_gin_blog/pkg/errorcode"
	"go_gin_blog/pkg/limiter"
)

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errorcode.TooManyRequests)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}