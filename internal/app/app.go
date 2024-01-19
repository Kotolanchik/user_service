package app

import (
	yamlConfig "user-service/internal/config"
	"user-service/internal/db"
	loggingCodes "user-service/internal/types/logging"
	logging "user-service/pkg/logger"
)

func Run() {
	logger := logging.New(logging.GetDefaultConfig())

	config, err := yamlConfig.GetConfigFromYAML()
	if err != nil {
		logger.Error(loggingCodes.CodeErrorReadConfigFailed, "error read config", err)
	}

	postgresConn, err := db.CreateConnectin(config.PostgresConnStr)
	if err != nil {
		logger.Error(loggingCodes.CodeErrorDBConnectionFailed, "error connect to db", err)
	}

	CreateAppHttpServer(":" + port)
}
