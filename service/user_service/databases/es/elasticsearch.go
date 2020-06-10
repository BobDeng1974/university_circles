package es

import (
	"context"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"net"
	"net/http"
	"path/filepath"
	"time"
	"university_circles/service/user_service/utils/logger"
)

//var host = "http://127.0.0.1:9200"

// Host struct
type Host struct {
	Host string
	Port string
	Username string
	Password string
}

// Config struct
type Config struct {
	Env string `toml:"env"`
	Es  struct {
		Dev  Host `toml:"development"`
		Prod Host `toml:"production"`
	} `toml:"elastic_search"`
}

// Es struct
type Es struct {
	Index  string
	Client http.Client
}

var hostURL string

// get es config
func getCfg() (Host, error) {
	var host Host
	var cfgPath string
	cfgPath, err := filepath.Abs("../../config/es.toml")
	if err != nil {
		return host, err
	}

	var cfg Config
	if _, err := toml.DecodeFile(cfgPath, &cfg); err != nil {
		return host, err
	}
	if "development" == cfg.Env {
		return cfg.Es.Dev, nil
	}
	return cfg.Es.Prod, nil
}

func NewElasticSearch() (client *elastic.Client, err error) {
	host, err := getCfg()
	if err != nil {
		logger.Logger.Error("es config file read error", zap.Error(err))
		panic(err)
	}
	hostURL = fmt.Sprintf("http://%s:%s", host.Host, host.Port)

	HttpClient := &http.Client{}
	HttpClient.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout: 30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns: 100,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout: 90 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	// 带验证连接
	client, err = elastic.NewClient(
		elastic.SetHttpClient(HttpClient),
		elastic.SetSniff(false),
		elastic.SetURL(hostURL),
		elastic.SetBasicAuth(host.Username, host.Password),
	)
	if err != nil {
		logger.Logger.Warn("New ElasticSearch failed", zap.Error(err))
		return nil, err
	}

	info, code, err := client.Ping(hostURL).Do(context.Background())
	if err != nil {
		logger.Logger.Warn("ElasticSearch client Ping failed", zap.Error(err))
		return nil, err
	}
	logger.Logger.Info("ElasticSearch retured with code and version ", zap.Any("code", code), zap.Any("Version", info.Version.Number))

	return
}
