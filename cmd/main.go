package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	c "github.com/ogonna-anaekwe/matrixserver/config"
	h "github.com/ogonna-anaekwe/matrixserver/internal/handler"
	l "github.com/ogonna-anaekwe/matrixserver/internal/logger"
	u "github.com/ogonna-anaekwe/matrixserver/internal/utils"

	"github.com/sirupsen/logrus"
)

const configFileLocation string = "./config.yml"

type Server struct {
	handler h.Handler
	server  *http.Server
	wg      *sync.WaitGroup
	ctx     context.Context
}

// Server constructor
func NewServer() Server {
	cfg := c.Config{}
	cfg.ParseConfig(configFileLocation)

	var w http.ResponseWriter
	rows, err := u.ReadFile(cfg.CSVFileLocation, w)
	if err != nil {
		logrus.Fatal(err)
	}

	logger := l.NewLogger()

	hh := h.Handler{Cfg: cfg, Log: logger, Rows: rows}
	s := &http.Server{Addr: cfg.Port}
	wg := &sync.WaitGroup{}
	ctx := context.Background()

	srv := Server{handler: hh, server: s, wg: wg, ctx: ctx}
	return srv
}

// Registers request handlers and starts API server.
func (s *Server) start() {
	log := s.handler.Log.WithFields(logrus.Fields{
		"module": "main",
		"method": "start",
	})

	defer s.wg.Done()
	defer s.ctx.Done()

	http.HandleFunc(h.EchoPath, s.handler.HandlerEcho)
	http.HandleFunc(h.SumPath, s.handler.HandlerSum)
	http.HandleFunc(h.InvertPath, s.handler.HandlerInvert)
	http.HandleFunc(h.FlattenPath, s.handler.HandlerFlatten)
	http.HandleFunc(h.MultiplyPath, s.handler.HandlerMultiply)

	log.Infof("✅ Started service on port %v", s.handler.Cfg.Port)

	err := s.server.ListenAndServe()
	if err != nil {
		log.Errorf("Server not listening on port %v", s.handler.Cfg.Port)
	}
}

// Gracefully stops API server.
func (s *Server) stop() {
	log := s.handler.Log.WithFields(logrus.Fields{
		"module": "main",
		"method": "stop",
	})

	shutdown := make(chan os.Signal, 1)
	signals := []os.Signal{syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM}
	signal.Notify(shutdown, signals...)
	<-shutdown

	log.Warnf("Received signal to shutdown server")

	err := s.server.Shutdown(s.ctx)
	if err != nil {
		log.Errorf("Could not shut down server %v", err)
	}

	s.wg.Wait()

	log.Infof("✅ Gracefully shutdown server")

	os.Exit(0)
}

// Entrypoint to server.
func main() {
	s := NewServer()

	s.wg.Add(1)
	go s.start() // non-blocking so we can await shut down signals below

	s.stop()
}
