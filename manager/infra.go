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
		db  *sql.DB
		cfg *config.Config
	}
)

func (i *infraManager) initDb() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		i.cfg.DbConfig.Host,
		i.cfg.DbConfig.Port,
		i.cfg.DbConfig.User,
		i.cfg.DbConfig.Password,
		i.cfg.DbConfig.Name,
		i.cfg.DbConfig.SSLMode,
		i.cfg.DbConfig.TimeZone,
	)
	db, err := sql.Open(i.cfg.DbConfig.Driver, dsn)
	if err != nil {
		return err
	}
	i.db = db
	fmt.Printf("successfully connected to database: 🚀%s🚀\n", i.cfg.DbConfig.Name)

	return nil
}

func (i *infraManager) Conn() *sql.DB {
	return i.db
}

// NewInfraManager constructor
func NewInfraManager(cfg *config.Config) (InfraManager, error) {
	conn := &infraManager{
		cfg: cfg,
	}
	err := conn.initDb()
	if err != nil {
		return nil, err
	}
	return conn, nil
}
