package main

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/zboyco/notify/notify/internal/config"
	"github.com/zboyco/notify/notify/internal/handler"
	"github.com/zboyco/notify/notify/internal/logic"
	"github.com/zboyco/notify/notify/internal/notify"
	"github.com/zboyco/notify/notify/internal/svc"
	"github.com/zboyco/notify/notify/model"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/notify-api.yaml", "the config file")
var c config.Config

func init() {
	conf.MustLoad(*configFile, &c)

	// 注册sender
	notify.RegisterSender(notify.NewWxPusher(c.Wxpusher.AppToken))
}

func main() {

	flag.Parse()
	// migrate database
	if strings.ToLower(flag.Arg(0)) == "migrate" {
		if err := model.Migrate(svc.ConnectDB(c)); err != nil {
			panic(err)
		}
		return
	}

	// 初始化
	server := rest.MustNewServer(c.RestConf, rest.WithCors("http://localhost:3000"))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 加载定时任务
	_ = logic.NewResetCronLogic(context.Background(), ctx).ResetCron(nil)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
