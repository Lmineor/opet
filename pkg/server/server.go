package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
}

func NewServer(flag *Flags)*Server {
	engine := initGinEngine(flag)
	return &Server{Engine: engine}
}

func (s *Server) Run(flag *Flags) error{
	addr := fmt.Sprintf("%s:%s", flag.IP, flag.Port)
	err:= s.Engine.Run(addr)
	return err
}