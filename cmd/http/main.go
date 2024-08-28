package http

import (
	"context"
	"github.com/nilroad/kateb"
	"github.com/spf13/cobra"
	"os"
	httpserver "sotoon/internal/adapter/http"
	userhandler "sotoon/internal/adapter/http/handler/user"
	"sotoon/internal/adapter/storage/mysql"
	userrepo "sotoon/internal/adapter/storage/mysql/repo/user"
	"sotoon/internal/config"
	usersvc "sotoon/internal/core/service/user"
)

type Command struct {
	logger kateb.Logger
}

func (r *Command) Register(ctx context.Context, cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "run http server",
		Run: func(_ *cobra.Command, _ []string) {
			r.run(ctx, cfg)
		},
	}
}

func (r *Command) run(ctx context.Context, cfg *config.Config) {
	db := mysql.New(cfg.MYSQLConfig, cfg.Debug)
	defer func() {
		if err := db.Close(); err != nil {
			r.logger.Error("failed to close db connection", nil)
		}
	}()

	// repos
	userRepo := userrepo.New(db)

	// service
	userService := usersvc.New(userRepo)

	// handlers
	userHandler := userhandler.New(userService)

	// setup server
	logger := kateb.New(os.Stdout, kateb.Config{
		Level:     kateb.ConvertToLevel(cfg.LogLevel),
		AddSource: false,
		Prefix:    "sotoon:server",
		Colorize:  false,
	})
	server := httpserver.New(cfg.HTTPServer, logger)

	server.SetUpAPIRoutes(userHandler)

	if err := server.Serve(ctx); err != nil {
		r.logger.Error("failed to start http server", nil)
	}
}
