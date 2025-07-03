package main

import (
	"flag"
	"fmt"
	"github.com/Sapronovps/RotationBanner/internal/app"
	"github.com/Sapronovps/RotationBanner/internal/logger"
	"github.com/Sapronovps/RotationBanner/internal/model"
	"github.com/Sapronovps/RotationBanner/internal/storage/memory"
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
	logg := logger.New(config.Logger.Level, config.Logger.File)
	storage := memory.New()

	application := app.NewApp(logg, storage)

	// Создаем СЛОТ
	newSlot := model.Slot{
		ID:          1,
		Description: "First Slot",
	}
	err := application.AddSlot(&newSlot)
	if err != nil {
		panic("Failed to add new slot")
	}

	slot, err := application.GetSlot(1)
	if err != nil {
		panic("Failed to get slot")
	}
	_ = slot

	// Создаем БАННЕР
	newBanner := model.Banner{
		Description: "First Banner",
	}
	err = application.AddBanner(&newBanner)
	if err != nil {
		panic("Failed to add new banner")
	}

	banner, err := application.GetBanner(1)
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

	group, err := application.GetGroup(1)
	if err != nil {
		panic("Failed to get group:" + err.Error())
	}
	_ = group

	// Создаем Привязку Слот -> Баннер -> Группа
	bannerGroupStats := &model.BannerGroupStats{
		SlotID:   slot.ID,
		BannerID: banner.ID,
		GroupID:  group.ID,
	}
	err = application.CreateBannerGroupStats(bannerGroupStats)
	if err != nil {
		panic("Failed to get banner")
	}

	secondBannerGroupStats := &model.BannerGroupStats{
		SlotID:   slot.ID,
		BannerID: newBanner2.ID,
		GroupID:  group.ID,
	}
	err = application.CreateBannerGroupStats(secondBannerGroupStats)
	if err != nil {
		panic("Failed to get banner")
	}

	fmt.Println(bannerGroupStats)

	// Регистрируем клик
	err = application.RegisterClick(slot.ID, banner.ID, group.ID)
	err = application.RegisterClick(slot.ID, banner.ID, group.ID)
	err = application.RegisterClick(slot.ID, newBanner2.ID, group.ID)
	if err != nil {
		panic("Failed to register click")
	}

	bannerGroupStats.Shows = 10
	secondBannerGroupStats.Shows = 5

	// Получим статистику по баннерам
	result, err := application.GetAndUpdateBanner(slot.ID, group.ID)
	if err != nil {
		panic("Failed to get and update banner")
	}

	fmt.Println(result)
}
