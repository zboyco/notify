package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	// 数据库配置
	Postgres struct {
		Endpoint string
	}

	// 鉴权配置
	Auth struct {
		Token string
	}

	// 微信推送配置
	Wxpusher struct {
		AppToken string
	}

	// 跨域配置
	Cors struct {
		AllowOrigins []string
	}
}
