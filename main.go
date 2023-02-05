package main

import (
	"fmt"

	"github.com/blogxp/goapi/configs"
	"github.com/blogxp/goapi/pkg/env"
	"github.com/blogxp/goapi/pkg/logger"
	"github.com/blogxp/goapi/pkg/timeutil"
	"go.uber.org/zap"
)

func main() {

	fmt.Println("Hello World!")
	test := configs.Get()
	fmt.Println(test.Mail.Host)
	// 初始化 access logger
	accessLogger, err := logger.NewJSONLogger(
		logger.WithDisableConsole(),
		logger.WithField("domain", fmt.Sprintf("%s[%s]", configs.ProjectName, env.Active().Value())),
		logger.WithTimeLayout(timeutil.CSTLayout),
		logger.WithFileP(configs.ProjectAccessLogFile),
	)
	if err != nil {
		panic(err)
	}

	// 初始化 cron logger
	cronLogger, err := logger.NewJSONLogger(
		logger.WithDisableConsole(),
		logger.WithField("domain", fmt.Sprintf("%s[%s]", configs.ProjectName, env.Active().Value())),
		logger.WithTimeLayout(timeutil.CSTLayout),
		logger.WithFileP(configs.ProjectCronLogFile),
	)

	if err != nil {
		panic(err)
	}

	defer func() {
		_ = accessLogger.Sync()
		_ = cronLogger.Sync()
	}()
	accessLogger.Fatal("http server startup err", zap.Error(err))
}
