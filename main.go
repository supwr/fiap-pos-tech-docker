package main

import (
	"github.com/gin-gonic/gin"
	"github.com/supwr/fiap-pos-tech-docker/src/application"
	"github.com/supwr/fiap-pos-tech-docker/src/domain/service"
	"github.com/supwr/fiap-pos-tech-docker/src/infra"
	"github.com/supwr/fiap-pos-tech-docker/src/infra/database"
	"github.com/supwr/fiap-pos-tech-docker/src/infra/repository"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	var err error
	var logger = loadLogger()

	r := gin.Default()

	cfg, err := loadConfig()

	if err != nil {
		logger.Error("error loading config", err)
		panic(err)
	}

	db, err := loadDatabase(cfg)

	if err != nil {
		logger.Error("error connecting to database", err)
		panic(err)
	}

	migrate(db, cfg)

	repo := repository.NewLanguageRepository(db)
	s := service.NewLanguageService(repo)
	app := application.NewApp(s)

	r.GET("/languages", func(c *gin.Context) {
		l, err := app.ListLanguages()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error listing languages",
			})

			return
		}

		c.JSON(http.StatusOK, l)
	})

	r.Run()
}

func migrate(db *gorm.DB, cfg infra.Config) {
	database.CreateSchema(db, cfg)
	database.MigrateDB(db)
	database.InsertLanguages(db, cfg)
}

func loadDatabase(cfg infra.Config) (*gorm.DB, error) {
	return database.NewConnection(
		cfg,
	)
}

func loadConfig() (infra.Config, error) {
	return infra.NewConfig()
}

func loadLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stderr, nil))
}
