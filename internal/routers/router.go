package routers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"go_gin_blog/internal/routers/api"
	_ "go_gin_blog/docs"
	"go_gin_blog/global"
	"go_gin_blog/internal/middleware" 
	"go_gin_blog/internal/routers/api/v1"
	"go_gin_blog/pkg/limiter"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})


func NewRouter() *gin.Engine {
	r := gin.New()

	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}

	r.Use(middleware.Translations()) //註冊中間件
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))

	article := v1.NewArticle()
  	tag := v1.NewTag()
	upload := api.NewUpload()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/auth", api.GetAuth)
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath)) //提供靜態資源訪問
	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.JWT())
	{
	  apiV1.POST("/tags", tag.Create)
      apiV1.DELETE("/tags/:id", tag.Delete)
      apiV1.PUT("/tags/:id", tag.Update)
      apiV1.PATCH("/tags/:id/state", tag.Update)
      apiV1.GET("/tags", tag.List)

      apiV1.POST("/articles", article.Create)
      apiV1.DELETE("/articles/:id", article.Delete)
      apiV1.PUT("/articles/:id", article.Update)
      apiV1.PATCH("/articles/:id/state", article.Update)
      apiV1.GET("/articles/:id", article.Get)
      apiV1.GET("/articles", article.List)
	}

	return r
}