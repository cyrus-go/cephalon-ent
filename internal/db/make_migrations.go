//go:build ignore
// +build ignore

package main

import (
	"ariga.io/atlas/sql/sqltool"
	"context"
	"entgo.io/ent/dialect/sql/schema"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/sirupsen/logrus"
	"os"
	// 需要初始化数据库配置
	"cephalon-ent/configs"
	"cephalon-ent/pkg/ent/migrate"
)

func main() {
	// 创建文件夹
	err := os.Mkdir("./migrations", 0777)
	if err != nil && !os.IsExist(err) {
		logrus.Errorf("failed at mkdir migrations")
	}
	ctx := context.Background()
	// 指定文件夹
	dir, err := sqltool.NewGolangMigrateDir("./migrations")
	if err != nil {
		logrus.Errorf("failed at creating atlas migration directory: %v", err)
		return
	}
	dbConf := configs.Conf.DBConfig
	// 迁移条件
	opts := []schema.MigrateOption{
		schema.WithDir(dir),
		// 配合 golang-migrate 使用
		schema.WithFormatter(sqltool.GolangMigrateFormatter),
		// 步进模式
		schema.WithMigrationMode(schema.ModeInspect),
		// 指定数据库类型
		schema.WithDialect("postgres"),
		// 移除外键约束
		schema.WithForeignKeys(false),
		// 可删字段
		schema.WithDropColumn(false),
		// 可删索引
		schema.WithDropIndex(false),
	}
	// 需知道数据库目标
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%v/%s?sslmode=disable&TimeZone=Asia/Shanghai", dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Database)
	// 开始创建迁移文件
	err = migrate.NamedDiff(ctx, dbUrl, "update", opts...)
	if err != nil {
		logrus.Errorf("failed at generating migrations, err: %v", err)
	}
	return
}
