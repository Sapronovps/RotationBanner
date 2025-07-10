package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Sapronovps/RotationBanner/internal/app"
	"github.com/Sapronovps/RotationBanner/internal/logger"
	"github.com/Sapronovps/RotationBanner/internal/model"
	"github.com/Sapronovps/RotationBanner/internal/server/http"
	"github.com/Sapronovps/RotationBanner/internal/storage"
	"github.com/Sapronovps/RotationBanner/internal/storage/memory"
	"github.com/Sapronovps/RotationBanner/internal/storage/sql"
	_ "github.com/lib/pq" // for postgres
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "/etc/calendar/config.yaml", "Path to configuration file")
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	config := NewConfig(configFile)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DB.Host, config.DB.Username, config.DB.Password, config.DB.DBName)
	var storageApp storage.Storage
	if config.DB.InMemory {
		storageApp = memory.New()
	} else {
		newDB := sql.NewDB(dsn)
		defer newDB.Close()
		storageApp = sql.New(newDB)
	}

	logg := logger.New(config.Logger.Level, config.Logger.File)
	application := app.NewApp(logg, storageApp)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	if config.Server.IsHTTP {
		serverAddress := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
		server := http.NewServer(serverAddress, application, logg)

		go func() {
			<-ctx.Done()

			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()

			if err := server.Stop(ctx); err != nil {
				logg.Error("failed to stop http server: " + err.Error())
			}
		}()

		logg.Info("Microservice Rotation Banner is running...")

		if err := server.Start(ctx); err != nil {
			logg.Error("failed to start http server: " + err.Error())
			cancel()
			os.Exit(1)
		}
	}
}

func testFunc() {
	// Кейсы для наполнения БД
	//testCreateData(application)

	//// Регистрируем клик
	//err := application.RegisterClick(1, 1, 1)
	//err = application.RegisterClick(1, 1, 1)
	//err = application.RegisterClick(1, 2, 1)
	//if err != nil {
	//	log.Fatalf("Failed to register click: %v", err)
	//}
	//
	//// Получим статистику по баннерам
	//result, err := application.GetBannerByMultiArmBandit(1, 1)
	//if err != nil {
	//	log.Fatalf("Failed to calculate statistic banner: %v", err)
	//}
	//
	//fmt.Println(result)
}

func testCreateData(application *app.App) {
	// Кейсы для наполнения БД
	// Создаем СЛОТ
	newSlot := model.Slot{
		Description: "Test",
	}
	err := application.AddSlot(&newSlot)
	if err != nil {
		panic("Failed to add new slot")
	}

	// Создаем БАННЕР
	newBanner := model.Banner{
		Description: "First Banner",
	}
	err = application.AddBanner(&newBanner)
	if err != nil {
		panic("Failed to add new banner")
	}

	// Создаем 2 БАННЕР
	newBanner2 := model.Banner{
		Description: "Second Banner",
	}
	err = application.AddBanner(&newBanner2)
	if err != nil {
		panic("Failed to add new banner")
	}

	// Создаем ГРУППУ
	newGroup := &model.Group{
		Title:       "Старики",
		Description: "First Group",
	}
	err = application.AddGroup(newGroup)
	if err != nil {
		panic("Failed to get banner")
	}
}
