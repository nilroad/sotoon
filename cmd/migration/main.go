package migration

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/nilroad/kateb"
	"github.com/spf13/cobra"
	"os"
	"sotoon/internal/config"
	"time"

	"github.com/go-sql-driver/mysql"
	migratemysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file" // for adding init of file migrations
)

const migrationSourceFilesAddr = "file://internal/adapter/storage/mysql/migration"

type Command struct {
	up      bool
	down    bool
	version bool
	step    int

	logger *kateb.Logger
}

func (r *Command) Register(ctx context.Context, cfg *config.Config) *cobra.Command {
	r.logger = kateb.New(os.Stdout, kateb.Config{
		Level:     kateb.ConvertToLevel(cfg.LogLevel),
		AddSource: false,
		Prefix:    "sotoon:cmd:migrate",
		Colorize:  false,
	})

	c := &cobra.Command{
		Use:   "migrate-mysql",
		Short: "migrate database",
		Run: func(_ *cobra.Command, _ []string) {
			r.run(ctx, cfg)
		},
	}

	c.Flags().BoolVar(&r.up, "up", false, "migrate up")
	c.Flags().BoolVar(&r.down, "down", false, "migrate down")
	c.Flags().BoolVar(&r.version, "version", false, "get current version")
	c.Flags().IntVar(&r.step, "step", 0, "migrate step, negative number for step down and positive for step up")

	return c
}

func (r *Command) run(_ context.Context, cfg *config.Config) {
	loc, err := time.LoadLocation(cfg.Tz)
	if err != nil {
		r.logger.Panic("failed to load location", map[string]interface{}{
			"error": err.Error(),
		})
	}

	c := mysql.Config{
		User:                    cfg.MYSQLConfig.Username,
		Passwd:                  cfg.MYSQLConfig.Password,
		DBName:                  cfg.MYSQLConfig.Name,
		Net:                     "tcp",
		Addr:                    fmt.Sprintf("%s:%d", cfg.MYSQLConfig.Host, cfg.MYSQLConfig.Port),
		AllowNativePasswords:    true,
		AllowCleartextPasswords: true,
		ParseTime:               true,
		MultiStatements:         true,
		Loc:                     loc,
		Params: map[string]string{
			"time_zone": "'" + cfg.Tz + "'",
		},
	}

	sqlDB, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		r.logger.Panic("failed to connect database", map[string]interface{}{
			"error": err.Error(),
		})
	}
	defer func() {
		if err := sqlDB.Close(); err != nil {
			r.logger.Error("failed to close database connection", map[string]interface{}{
				"error": err.Error(),
			})
		}
	}()

	driver, err := migratemysql.WithInstance(sqlDB, &migratemysql.Config{})
	if err != nil {
		r.logger.Panic("failed to connect database", map[string]interface{}{
			"error": err.Error(),
		})
	}

	m, err := migrate.NewWithDatabaseInstance(migrationSourceFilesAddr, "mysql", driver)
	if err != nil {
		r.logger.Panic("failed to connect database", map[string]interface{}{
			"error": err.Error(),
		})
	}

	if r.up {
		r.Up(m)
	}
	if r.down {
		r.Down(m)
	}
	if r.version {
		r.Version(m)
	}
	if r.step != 0 {
		r.Step(m, r.step)
	}
}

func (r *Command) Up(m *migrate.Migrate) {
	t := time.Now()
	if err := m.Up(); err != nil {
		r.logger.Error("failed to run migration", map[string]interface{}{
			"error": err.Error(),
		})

		return
	}

	r.logger.Info("migration added successfully", map[string]interface{}{
		"duration": time.Since(t),
	})
}

func (r *Command) Down(m *migrate.Migrate) {
	t := time.Now()
	if err := m.Down(); err != nil {
		r.logger.Error("failed to run migration", map[string]interface{}{
			"error": err.Error(),
		})

		return
	}

	r.logger.Info("migration downed successfully", map[string]interface{}{
		"duration": time.Since(t),
	})
}

func (r *Command) Version(m *migrate.Migrate) {
	version, dirty, err := m.Version()
	if err != nil {
		r.logger.Error("failed to get migration version", map[string]interface{}{
			"error": err.Error(),
		})

		return
	}

	r.logger.Info("migration version", map[string]interface{}{
		"version": version,
		"dirty":   dirty,
	})
}

func (r *Command) Step(m *migrate.Migrate, step int) {
	t := time.Now()
	if err := m.Steps(step); err != nil {
		r.logger.Error("failed to run step on migration", map[string]interface{}{
			"error": err.Error(),
			"steps": step,
		})

		return
	}

	r.logger.Info("migration step add successfully", map[string]interface{}{
		"duration": time.Since(t),
		"steps":    step,
	})
}
