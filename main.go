package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/caarlos0/env/v9"
	"github.com/ghostrepo00/go-dashboard-api/api"
	appconstant "github.com/ghostrepo00/go-dashboard-api/domain/app_constant"
	"github.com/ghostrepo00/go-dashboard-api/domain/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetLogFileName(appSettings *model.AppSettings) string {
	currentDate := time.Now()
	result := fmt.Sprintf(appSettings.Log.FileNameFormat, currentDate.Format(appconstant.TimestampFormat))
	return result
}

func LoadAppSettings(file string) *model.AppSettings {
	var appSettings *model.AppSettings
	appSettingsFile, err := os.Open(file)
	defer appSettingsFile.Close()
	if err != nil {
		panic(err)
	}
	jsonParser := json.NewDecoder(appSettingsFile)
	jsonParser.Decode(&appSettings)
	return appSettings
}

func SetupApiRouters(appSettings *model.AppSettings, dbClient *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	api.SetupRouters(router, appSettings, dbClient)
	api.SetupSwagger(router, appSettings)

	return router
}

func main() {
	// Load environment settings
	envSettings := model.EnvironmentSettings{}
	if err := env.Parse(&envSettings); err != nil {
		panic(err)
	}

	// Load app settings based on environment settings
	configFileExtension := ""
	if envSettings.EnvironmentName == "dev" {
		configFileExtension = ".dev"
	}
	appSettings := LoadAppSettings(fmt.Sprintf("%s%s", ".app_settings", configFileExtension))

	// configure slog
	logFile, err := os.OpenFile(GetLogFileName(appSettings), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	slogHandler := slog.NewJSONHandler(
		io.MultiWriter(os.Stdout, logFile),
		&slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		})
	logger := slog.New(slogHandler)
	slog.SetDefault(logger)

	// start app
	slog.Info("Starting app")

	dbClient, err := gorm.Open(postgres.Open(envSettings.ConnectionString), &gorm.Config{})
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
	dbCon, _ := dbClient.DB()
	defer dbCon.Close()

	slog.Info("Database initialized")

	router := SetupApiRouters(appSettings, dbClient)

	slog.Info("Routers initialized")

	router.Run(appSettings.Api.Host)

	slog.Info("App terminated")
}
