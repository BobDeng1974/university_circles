package config

const SRV_NAME = "university_circles.srv.im"

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
