package server

import (
	"eos-layout/internal/config"
	"eos-layout/internal/handler"
	"eos-layout/internal/middleware"
	v1 "eos-layout/internal/server/v1"
	"eos-layout/pkg/http"
	"eos-layout/pkg/log"

	"github.com/gin-gonic/gin"
)

func NewHTTPServer(
	cfg *config.Config,
	log *log.Logger,
	areaHandler handler.AreaHandler,
) Server {
	return &httpServer{
		cfg:         cfg,
		log:         log,
		areaHandler: areaHandler,
	}
}

// 服务接口，仅提供Run方法启动服务
type Server interface {
	Run()
}

type httpServer struct {
	cfg         *config.Config
	log         *log.Logger
	areaHandler handler.AreaHandler
}

func (s *httpServer) Run() {
	defer func() {
		if rec := recover(); rec != nil {
			s.log.Sugar().Errorln("server run error: ", rec)
		}
	}()
	g := s.router()
	http.Run(g, s.cfg.Host)
}

/**
 * 统一路由管理
 */
func (s *httpServer) router() *gin.Engine {
	if s.cfg.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	e := gin.Default()
	// 中间件配置
	e.Use(
		middleware.CORS(),
		middleware.ResponseLogger(s.log),
		middleware.RequestLogger(s.log),
	)
	// 路由配置
	api := e.Group("/api")
	v1.Register(api).AreaRouter(s.areaHandler)
	return e
}
