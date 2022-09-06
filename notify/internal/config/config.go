package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	Postgres struct {
		Endpoint string
	}

	Auth struct {
		Token string
	}

	Wxpusher struct {
		AppToken string
	}
}
