package main

import (
	"flag"
	"fmt"
	"github.com/Sapronovps/RotationBanner/internal/app"
	"github.com/Sapronovps/RotationBanner/internal/logger"
	"github.com/Sapronovps/RotationBanner/internal/model"
	"github.com/Sapronovps/RotationBanner/internal/storage"
	"github.com/Sapronovps/RotationBanner/internal/storage/memory"
	"github.com/Sapronovps/RotationBanner/internal/storage/sql"
	_ "github.com/lib/pq" // for postgres
	"log"
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
	_ = application

	// Создаем СЛОТ
	newSlot := model.Slot{
		Description: "Test",
	}
	err := application.AddSlot(&newSlot)
	if err != nil {
		logg.Fatal(err.Error())
	}

	slot, err := application.GetSlot(newSlot.ID)
	if err != nil {
		log.Fatalf("Could not get slot: %v", err)
	}
	_ = slot

	fmt.Println(slot)

	// Создаем БАННЕР
	newBanner := model.Banner{
		Description: "First Banner",
	}
	err = application.AddBanner(&newBanner)
	if err != nil {
		panic("Failed to add new banner")
	}

	banner, err := application.GetBanner(newBanner.ID)
	if err != nil {
		panic("Failed to get banner")
	}
	_ = banner

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
	err = application.CreateGroup(newGroup)
	if err != nil {
		panic("Failed to get banner")
	}

	group, err := application.GetGroup(newGroup.ID)
	if err != nil {
		panic("Failed to get group:" + err.Error())
	}
	_ = group

	fmt.Println(group)

	// Регистрируем клик
	err = application.RegisterClick(slot.ID, banner.ID, group.ID)
	err = application.RegisterClick(slot.ID, banner.ID, group.ID)
	err = application.RegisterClick(slot.ID, newBanner2.ID, group.ID)
	if err != nil {
		log.Fatalf("Failed to register click: %v", err)
	}

	// Получим статистику по баннерам
	result, err := application.GetAndUpdateBanner(slot.ID, group.ID)
	if err != nil {
		log.Fatalf("Failed to calculate statistic banner: %v", err)
	}

	fmt.Println(result)
}
