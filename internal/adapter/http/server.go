package httpserver

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nilroad/kateb"
	"net/http"
	"sotoon/internal/config"
	"time"
)

const headerReadTimeout = time.Second * 10

type Server struct {
	engine *gin.Engine
	logger *kateb.Logger
	cfg    config.HTTPServer
}

func New(cfg config.HTTPServer, logger *kateb.Logger) *Server {
	return &Server{
		engine: gin.New(),
		cfg:    cfg,
		logger: logger,
	}
}

func (r *Server) Serve(ctx context.Context) error {
	srv := &http.Server{
		Handler:           r.engine,
		Addr:              fmt.Sprintf("%s:%d", r.cfg.Host, r.cfg.Port),
		ReadHeaderTimeout: headerReadTimeout,
	}

	r.logger.Info("starting http server", nil)
	srvErr := make(chan error)
	go func() {
		srvErr <- srv.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		r.logger.Info("http server is shutting down", nil)

		return srv.Shutdown(ctx)
	case err := <-srvErr:
		return err
	}
}
