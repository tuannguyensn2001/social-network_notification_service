package main

import (
	"go.uber.org/zap"
	"log"
	"social-work_notification_service/src/cmd"
	"social-work_notification_service/src/config"
	_const "social-work_notification_service/src/const"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}

	if cfg.GetEnvironment() == _const.DEVELOPMENT {
		logger, err := zap.NewDevelopment()
		if err != nil {
			log.Fatalln(err)
		}
		defer logger.Sync()
		zap.ReplaceGlobals(logger)
	} else if cfg.GetEnvironment() == _const.PRODUCTION {
		logger, err := zap.NewProduction()
		if err != nil {
			log.Fatalln(err)
		}
		defer logger.Sync()
		zap.ReplaceGlobals(logger)
	}

	err = cmd.GetRoot(cfg).Execute()
	if err != nil {
		zap.S().Fatalln(err)
	}

}
