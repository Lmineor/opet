package server

import (
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	//Engine *gin.Engine
	*http.Server
}

func NewServer(flag *Flags) *Server {
	router := initGinEngine(flag)
	addr := fmt.Sprintf("%s:%d", flag.IP, flag.Port)
	srv := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return &Server{Server: srv}
}

//func NewServer(flag *Flags)*Server {
//	engine := initGinEngine(flag)
//	return &Server{Engine: engine}
//}

func (s *Server) Run(flag *Flags) error {
	err := s.ListenAndServe()
	return err
}
