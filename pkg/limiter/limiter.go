package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type LimiterIface interface {
	Key(c *gin.Context) string //取得對應的限流器鍵值對名稱
	GetBucket(key string) (*ratelimit.Bucket, bool) //取得令牌桶
	AddBuckets(rules ...LimiterBucketRule) LimiterIface //新增多個令牌桶
}

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

type LimiterBucketRule struct {
	Key          string //自定義鍵值對名稱
	FillInterval time.Duration //間隔多久時間放N個令牌
	Capacity     int64 //令牌桶的容量
	Quantum      int64 //每次到達間隔時間後所放的具體令牌數量
}