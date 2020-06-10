package mysql

import (
	"fmt"
	"university_circles/service/common_service/utils/zaplog"

	"database/sql"

	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

// default database toml config path
const (
	DefaultDatabaseTomlConfigPath = "../../config/database.toml"
	DefaultDriver                 = "mysql"
	DefaultTimeout                = 1
	DefaultMaxIdle                = 10
	DefaultMaxOpen                = 1000
	DefaultMaxLifetime            = 300
)

// database config
type DBConfig struct {
	Driver      string
	Host        string
	Port        string
	Database    string
	Username    string
	Password    string
	Timeout     int64
	MaxIdle     int64
	MaxOpen     int64
	MaxLifetime int64
}

// get toml config
func getDatabaseConfig(instance string) (*DBConfig, error) {
	// get real instance
	var config map[string]interface{}
	_, err := toml.DecodeFile(DefaultDatabaseTomlConfigPath, &config)
	if err != nil {
		fmt.Println("打开配置文件", err)
		zaplog.Info("decode database toml file failed", zaplog.String("error", err.Error()))
		return nil, err
	}
	env := config["env"].(string)
	fmt.Println("打开配置文件变量", env)
	if instanceConf, ok := config[instance].(map[string]interface{}); ok {
		if envConf, ok := instanceConf[env].(map[string]interface{}); ok {
			c := &DBConfig{
				Driver:      DefaultDriver,
				Host:        envConf["host"].(string),
				Port:        envConf["port"].(string),
				Database:    envConf["database"].(string),
				Username:    envConf["username"].(string),
				Password:    envConf["password"].(string),
				Timeout:     DefaultTimeout,
				MaxIdle:     DefaultMaxIdle,
				MaxOpen:     DefaultMaxOpen,
				MaxLifetime: DefaultMaxLifetime,
			}
			if driver, ok := envConf["driver"].(string); ok {
				c.Driver = driver
			}
			if timeout, ok := envConf["timeout"].(int64); ok {
				c.Timeout = timeout
			}
			if maxIdle, ok := envConf["maxIdle"].(int64); ok {
				c.Timeout = maxIdle
			}
			if maxOpen, ok := envConf["maxOpen"].(int64); ok {
				c.Timeout = maxOpen
			}
			if maxLifetime, ok := envConf["maxLifetime"].(int64); ok {
				c.Timeout = maxLifetime
			}
			return c, nil
		} else {
			return nil, errors.New("invalid database instance " + instance)
		}
	} else {
		zaplog.Warn("invalid database instance " + instance)
		return nil, errors.New("invalid database instance " + instance)
	}
}

// get dsn
func getDSN(c *DBConfig) string {
	//temp := "%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=%s&parseTime=true"
	temp := "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true"
	dsn := fmt.Sprintf(temp, c.Username, c.Password,
		c.Host, c.Port, c.Database)
	fmt.Println(dsn)
	//c.Host, c.Port, c.Database, url.QueryEscape("Asia/Shanghai"))
	return dsn
}

// new mysql
func NewMySQL(instance string) (*sql.DB, error) {
	var dbInstance *sql.DB
	conf, err := getDatabaseConfig(instance)
	fmt.Println("数据库配置", conf)
	if err != nil {
		zaplog.Warn("get config error", zaplog.String("error", err.Error()))
		return nil, err
	}
	dsn := getDSN(conf)
	fmt.Println(dsn)
	dbInstance, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("dsn failed", err)
		zaplog.Warn("gorm open failed", zaplog.String("error", err.Error()))
		return nil, err
	}
	return dbInstance, nil
}
