package config

import "flag"

const (
	SRV_NAME = "university_circles.srv.api"

	SRV_HOME_NAME    = "university_circles.srv.home"
	CLIENT_HOME_NAME = "university_circles.client.home"

	SRV_USER_NAME    = "university_circles.srv.user"
	CLIENT_USER_NAME = "university_circles.client.user"

	SRV_COMMON_NAME    = "university_circles.srv.common"
	CLIENT_COMMON_NAME = "university_circles.client.common"

	SRV_IM_NAME    = "university_circles.srv.im"
	CLIENT_IM_NAME = "university_circles.client.im"
)

var configFile string

type (
	Config struct {
		Version string
		Etcd    struct {
			Addr     []string
			UserName string
			Password string
		}
	}
)

func (c *Config) GetConfigFile() string {
	if configFile != "" {
		return configFile
	}
	var cFile string
	flag.StringVar(&cFile, "f", "./config/config.json", "please use config.json")
	flag.Parse()
	configFile = cFile
	return configFile
}
