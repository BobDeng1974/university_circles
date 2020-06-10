package redis

import (
	"fmt"
	"sync"
	"time"
	"university_circles/service/home_service/utils/zaplog"

	"github.com/BurntSushi/toml"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

var RedisObj = &MyRedis{}

var Once sync.Once

type MyRedis struct {
	RedisPoolInstance *redis.Pool
}

// RedisConfig redis config
type RedisConfig struct {
	Addr         string
	Password     string
	MaxIdleConns int
	MaxOpenConns int
	CTimeout     time.Duration
	RTimeout     time.Duration
	WTimeout     time.Duration
	ITimeout     time.Duration
}

const (
	DefaultRedisTomlConfigPath = "../../config/redis.toml"
	DefaultMaxIdle             = 10
	DefaultMaxOpen             = 1000
	DefaultConnectTimeout      = 1 * time.Second
	DefaultReadTimeout         = 3 * time.Second
	DefaultWriteTimeout        = 3 * time.Second
	DefaultIdleTimeout         = 300 * time.Second
)

// get json redis config
func (r *MyRedis) getRedisConfig(index string) (*RedisConfig, error) {
	var config map[string]interface{}
	_, err := toml.DecodeFile(DefaultRedisTomlConfigPath, &config)
	if err != nil {
		zaplog.Info("decode redis toml file failed", zaplog.String("error", err.Error()))
		return nil, err
	}
	env := config["env"].(string)
	if instanceConf, ok := config[index].(map[string]interface{}); ok {
		if envConf, ok := instanceConf[env].(map[string]interface{}); ok {
			c := &RedisConfig{
				Addr:         fmt.Sprintf("%s:%s", envConf["host"].(string), envConf["port"].(string)),
				Password:     envConf["password"].(string),
				MaxIdleConns: DefaultMaxIdle,
				MaxOpenConns: DefaultMaxOpen,
				CTimeout:     DefaultConnectTimeout,
				RTimeout:     DefaultReadTimeout,
				WTimeout:     DefaultWriteTimeout,
				ITimeout:     DefaultIdleTimeout,
			}
			if maxIdle, ok := envConf["maxIdle"].(int64); ok {
				c.MaxIdleConns = int(maxIdle)
			}
			if maxOpen, ok := envConf["maxOpen"].(int64); ok {
				c.MaxOpenConns = int(maxOpen)
			}
			if cTimeout, ok := envConf["connect_timeout"].(int64); ok {
				c.CTimeout = time.Duration(cTimeout) * time.Second
			}
			if rTimeout, ok := envConf["read_timeout"].(int64); ok {
				c.RTimeout = time.Duration(rTimeout) * time.Second
			}
			if wTimeout, ok := envConf["write_timeout"].(int64); ok {
				c.WTimeout = time.Duration(wTimeout) * time.Second
			}
			if iTimeout, ok := envConf["idle_timeout"].(int64); ok {
				c.ITimeout = time.Duration(iTimeout) * time.Second
			}
			return c, nil
		} else {
			return nil, errors.New("invalid database instance " + index)
		}
	} else {
		zaplog.Warn("invalid database instance " + index)
		return nil, errors.New("invalid database instance " + index)
	}
}

// NewRedis new redis
func (r *MyRedis) NewRedis(index string) (*redis.Pool, error) {
	conf, err := r.getRedisConfig(index)
	if err != nil {
		zaplog.Warn("get redis config failed", zaplog.String("error", err.Error()))
		return nil, err
	}
	pool := &redis.Pool{
		MaxIdle:   conf.MaxIdleConns,
		MaxActive: conf.MaxOpenConns,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", conf.Addr,
				redis.DialPassword(conf.Password),
				redis.DialConnectTimeout(conf.CTimeout),
				redis.DialReadTimeout(conf.RTimeout),
				redis.DialWriteTimeout(conf.WTimeout))
			if err != nil {
				zaplog.Warn("redis connect failed", zaplog.String("error", err.Error()))
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		IdleTimeout: conf.ITimeout,
		Wait:        false,
	}
	zaplog.Info("new redis pool success")
	return pool, nil
}

// GetInstance method
func (r *MyRedis) GetInstance(index string) *MyRedis {
	Once.Do(func() {
		pool, err := r.NewRedis(index)
		fmt.Println("GetInstance", pool, err)
		if err != nil {
			zaplog.Warn(" new redis instance failed", zaplog.String("error", err.Error()))
			panic(err)
		}
		RedisObj.RedisPoolInstance = pool
	})
	return RedisObj
}

var mRedis MyRedis

// DefaultRedisPool 默认redis连接池
var DefaultRedisPool = mRedis.GetInstance("default_redis").RedisPoolInstance
