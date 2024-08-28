package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sotoon/internal/config"
	"time"
)

type SQLDb struct {
	config config.MYSQLConfig
	db     *gorm.DB
	sqlDB  *sql.DB
}

func New(cfg config.MYSQLConfig, debug bool) (*SQLDb, error) {
	loc, err := time.LoadLocation(cfg.Tz)
	if err != nil {
		panic(err)
	}

	c := mysql.Config{
		User:                    cfg.Username,
		Passwd:                  cfg.Password,
		DBName:                  cfg.Name,
		Net:                     "tcp",
		Addr:                    fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		AllowNativePasswords:    true,
		AllowCleartextPasswords: true,
		ParseTime:               true,
		MultiStatements:         true,
		Loc:                     loc,
	}
	sqlDB, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		return nil, err
	}

	gormCfg := gorm.Config{}
	if debug {
		gormCfg.Logger = logger.Default.LogMode(logger.Info)
	} else {
		gormCfg.Logger = logger.Default.LogMode(logger.Silent)
	}
	gormDB, err := gorm.Open(gormmysql.New(gormmysql.Config{
		Conn: sqlDB,
	}), &gormCfg)
	if err != nil {
		return nil, err
	}

	return &SQLDb{
		db:     gormDB,
		config: cfg,
		sqlDB:  sqlDB,
	}, nil
}

func (r *SQLDb) DB(ctx context.Context) *gorm.DB {
	db, ok := ctx.Value(r.config.Name).(*SQLDb)
	if ok {
		return db.db
	}

	return r.db
}

func (r *SQLDb) Close() error {
	return r.sqlDB.Close()
}

func (r *SQLDb) Begin() (any, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &SQLDb{db: tx}, nil
}

func (r *SQLDb) Commit() error {
	return r.db.Commit().Error
}

func (r *SQLDb) Rollback() error {
	return r.db.Rollback().Error
}
