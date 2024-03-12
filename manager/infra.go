package manager

import (
	"database/sql"
	"fmt"
	"github.com/EkoEdyPurwanto/goforit/config"
	_ "github.com/lib/pq"
)

type (
	InfraManager interface {
		Conn() *sql.DB
	}

	infraManager struct {
		DB  *sql.DB
		Cfg *config.Config
	}
)

func (i *infraManager) initDb() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		i.Cfg.DbConfig.Host,
		i.Cfg.DbConfig.Port,
		i.Cfg.DbConfig.User,
		i.Cfg.DbConfig.Password,
		i.Cfg.DbConfig.Name,
		i.Cfg.DbConfig.SSLMode,
		i.Cfg.DbConfig.TimeZone,
	)
	db, err := sql.Open(i.Cfg.DbConfig.Driver, dsn)
	if err != nil {
		return err
	}
	i.DB = db
	fmt.Printf("successfully connected to database: 🚀%s🚀\n", i.Cfg.DbConfig.Name)

	return nil
}

func (i *infraManager) Conn() *sql.DB {
	return i.DB
}

// NewInfraManager constructor
func NewInfraManager(cfg *config.Config) (InfraManager, error) {
	conn := &infraManager{
		Cfg: cfg,
	}
	err := conn.initDb()
	if err != nil {
		return nil, err
	}
	return conn, nil
}
