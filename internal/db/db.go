package db

import (
	"cephalon-ent/configs"
	"cephalon-ent/pkg/cep_ent"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"time"
)

var url = ""
var DB *cep_ent.Client

func init() {
	dbConf := configs.Conf.DBConfig
	url = fmt.Sprintf("postgres://%s:%s@%s:%v/%s?sslmode=disable&TimeZone=Asia/Shanghai", dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Database)
	fmt.Printf("db_url is :%s", url)
	DB = NewDBClient(dbConf)
	err := migrateWithMigrationFiles()
	if err != nil {
		panic(fmt.Sprintf("db init failed: %v", err))
	}
}

func migrateWithMigrationFiles() (err error) {
	m, err := migrate.New("file://internal/db/migrations", url)
	if err != nil {
		logrus.Errorf("failed at new migrate, err: %v", err)
		return err
	}
	// 没有变更不算错误
	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logrus.Errorf("faied at migrating: %v", err)
		return err
	}
	return nil
}

func NewDBClient(dbConf configs.DBConfig) *cep_ent.Client {
	dataSourceName := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", dbConf.Host, dbConf.Port, dbConf.Username, dbConf.Password, dbConf.Database)
	logrus.Debugf("dsn: %s\n", dataSourceName)
	db, err := sql.Open("pgx", dataSourceName)
	if err != nil {
		panic(fmt.Sprintf("new db client failed: %v", err))
	}
	db.SetConnMaxLifetime(time.Hour)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(100)
	if err != nil {
		logrus.Errorf("failed at creating ent client with db %s, err: %v", dataSourceName, err)
		panic(fmt.Sprintf("new db client failed: %v", err))
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	return cep_ent.NewClient(cep_ent.Driver(drv))
}
