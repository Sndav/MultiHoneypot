package website

import (
	"Multi-Honeypot/internal/app/website/controllers/victims"
	"Multi-Honeypot/internal/app/website/middlewares"
	"Multi-Honeypot/internal/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

const (
	DebugMode   = "debug"
	ReleaseMode = "release"
)

type Server struct {
	_gin  *gin.Engine
	_gorm *gorm.DB
	_conf *config.Config
}

func (s *Server) setRouter() {
	victims.RegisterRouter(s._gin)
}

func (s *Server) initGin() {
	s._gin = gin.New()
	gin.SetMode(getMode())
	s._gin.Use(gin.Logger())
	s._gin.Use(middlewares.Recovery())
	s._gin.Use(middlewares.SetDB(s._gorm))
	s._gin.Use(middlewares.SetConfig(s._conf))
	s._gin.Use(middlewares.Auth())
	s.setRouter()
}

func (s *Server) initGorm() {
	var err error
	s._gorm, err = gorm.Open(s._conf.Get("DB", "type"), s._conf.Get("DB", "connect_url"))
	if err != nil {
		panic(err)
	}
}

func (s *Server) run() {
	var err error
	err = s._gin.Run(s._conf.Get("panel", "port"))
	if err != nil {
		panic(err)
	}
}

func (s *Server) Start() {
	s.initGorm()
	s.initGin()
	s.run()
}

func getMode() string {
	if os.Getenv("GYM_SERVER_MODE") != DebugMode {
		return ReleaseMode
	}
	return DebugMode
}

func NewServer(configFile string) *Server {
	return &Server{
		_conf: config.NewConfig(configFile),
	}
}
