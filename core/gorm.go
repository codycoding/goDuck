package core

import (
	"github.com/codycoding/goDuck/core/internal"
	"github.com/codycoding/goDuck/global"
	"github.com/codycoding/goDuck/global/config"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

//
// GormMysql
//  @Description: mysql连接
//  @param m
//  @return *gorm.DB
//
func GormMysql(m config.Mysql) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig(m.LogMode, m.LogZap)); err != nil {
		global.Log.Error("MySQL启动异常", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		global.Log.Info("Mysql数据库已连接")
		return db
	}
}

//
// GormPostgres
//  @Description: Postgre连接实例
//  @param m
//  @return *gorm.DB
//
func GormPostgres(m config.Postgres) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	dsn := m.Dsn()
	postgresConfig := postgres.Config{
		DSN:                  dsn,  // DSN data source name
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}
	if db, err := gorm.Open(postgres.New(postgresConfig), gormConfig(m.LogMode, m.LogZap)); err != nil {
		global.Log.Error("Postgres数据库启动异常", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		global.Log.Info("Postgres数据库已连接")
		return db
	}
}

func gormConfig(logMod bool, logZap string) *gorm.Config {
	gConfig := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	if logMod {
		switch logZap {
		case "silent", "Silent":
			gConfig.Logger = internal.Default.LogMode(logger.Silent)
		case "error", "Error":
			gConfig.Logger = internal.Default.LogMode(logger.Error)
		case "warn", "Warn":
			gConfig.Logger = internal.Default.LogMode(logger.Warn)
		case "info", "Info":
			gConfig.Logger = internal.Default.LogMode(logger.Info)
		default:
			gConfig.Logger = internal.Default.LogMode(logger.Info)
		}
	} else {
		gConfig.Logger = internal.Default.LogMode(logger.Silent)
	}
	return gConfig
}
