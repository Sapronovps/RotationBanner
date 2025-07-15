package app

import (
	"fmt"
	"github.com/Sapronovps/RotationBanner/internal/broker/kafka"
	"github.com/Sapronovps/RotationBanner/internal/model"
	"github.com/Sapronovps/RotationBanner/internal/service"
	"github.com/Sapronovps/RotationBanner/internal/storage"
	"time"

	"go.uber.org/zap"
)

type App struct {
	logger        *zap.Logger
	storage       storage.Storage
	kafkaProducer *kafka.Producer
}

func NewApp(logger *zap.Logger, storage storage.Storage, kafkaProducer *kafka.Producer) *App {
	return &App{logger: logger, storage: storage, kafkaProducer: kafkaProducer}
}

func (a *App) AddSlot(slot *model.Slot) error {
	err := a.storage.Banner().CreateSlot(slot)
	successMessage := fmt.Sprintf("Слот успешно создан с ID: %d", slot.ID)
	a.kafkaProducer.SendCustomMessage(err, successMessage, "addSlot", a.logger)

	return err
}

func (a *App) GetSlot(id int) (slot *model.Slot, err error) {
	slot, err = a.storage.Banner().GetSlot(id)
	successMessage := fmt.Sprintf("Слот успешно получен с ID: %d", id)
	a.kafkaProducer.SendCustomMessage(err, successMessage, "getSlot", a.logger)

	return slot, err
}

func (a *App) AddBanner(banner *model.Banner) error {
	err := a.storage.Banner().CreateBanner(banner)

	successMessage := fmt.Sprintf("Баннер успешно создан с ID: %d", banner.ID)
	a.kafkaProducer.SendCustomMessage(err, successMessage, "addBanner", a.logger)

	return err
}

func (a *App) GetBanner(id int) (banner *model.Banner, err error) {
	banner, err = a.storage.Banner().GetBanner(id)

	successMessage := fmt.Sprintf("Баннер успешно получен с ID: %d", id)
	a.kafkaProducer.SendCustomMessage(err, successMessage, "getBanner", a.logger)

	return banner, err
}

func (a *App) AddGroup(group *model.Group) error {
	err := a.storage.Banner().CreateGroup(group)

	successMessage := fmt.Sprintf("Группа успешно создана с ID: %d", group.ID)
	a.kafkaProducer.SendCustomMessage(err, successMessage, "addGroup", a.logger)

	return err
}

func (a *App) GetGroup(id int) (group *model.Group, err error) {
	group, err = a.storage.Banner().GetGroup(id)

	successMessage := fmt.Sprintf("Группа успешно получена с ID: %d", id)
	a.kafkaProducer.SendCustomMessage(err, successMessage, "getGroup", a.logger)

	return group, err
}

func (a *App) AddBannerGroupStats(stats *model.BannerGroupStats) error {
	stat, err := a.GetBannerGroupStats(stats.SlotID, stats.BannerID, stats.GroupID)
	if err == nil {
		stats.ID = stat.ID
		stats.Shows = stat.Shows
		stats.Clicks = stat.Clicks
		stats.CreatedAt = stat.CreatedAt
		stats.UpdatedAt = stat.UpdatedAt
		return nil
	}

	successMessage := fmt.Sprintf("Статистика успешно создана с ID: %d", stats.ID)
	a.kafkaProducer.SendCustomMessage(err, successMessage, "addStats", a.logger)

	return a.storage.Banner().CreateBannerGroupStats(stats)
}

func (a *App) GetBannerGroupStats(slotID, bannerID, groupID int) (*model.BannerGroupStats, error) {
	stats, err := a.storage.Banner().GetBannerGroupStats(slotID, bannerID, groupID)

	successMessage := fmt.Sprintf("Статистика успешно получена с banner ID: %d", bannerID)
	a.kafkaProducer.SendCustomMessage(err, successMessage, "getStats", a.logger)

	return stats, err
}

func (a *App) RegisterClick(slotID, bannerID, groupID int) error {
	stats, err := a.storage.Banner().GetBannerGroupStats(slotID, bannerID, groupID)
	if err != nil {
		stats = &model.BannerGroupStats{
			SlotID:   slotID,
			BannerID: bannerID,
			GroupID:  groupID,
		}
		err = a.AddBannerGroupStats(stats)
		if err != nil {
			return err
		}
	}
	stats.Clicks++
	stats.Shows++
	stats.UpdatedAt = time.Now()

	successMessage := fmt.Sprintf("Регистрация клика прошла успешна для баннера с ID  %d", bannerID)
	a.kafkaProducer.SendCustomMessage(err, successMessage, "registerClick", a.logger)

	return a.storage.Banner().UpdateBannerGroupStats(stats)
}

func (a *App) GetBannerByMultiArmBandit(slotID, groupID int) (banner *model.Banner, err error) {
	bannersStats := a.storage.Banner().GetBannersGroupStats(slotID, groupID)
	if bannersStats == nil {
		return nil, fmt.Errorf("banner group stats not found slot id: %d and group id: %d", slotID, groupID)
	}

	bannerID := service.CalculateBannerIDByMultiArmBandit(bannersStats)

	if bannerID > 0 {
		banner, err := a.storage.Banner().GetBanner(bannerID)
		if err != nil {
			return nil, err
		}
		stats, err := a.storage.Banner().GetBannerGroupStats(slotID, bannerID, groupID)
		if err != nil {
			return nil, err
		}
		stats.Shows++
		stats.UpdatedAt = time.Now()

		err = a.storage.Banner().UpdateBannerGroupStats(stats)
		if err != nil {
			return nil, err
		}

		successMessage := fmt.Sprintf("Баннер успешно рассчитан по алгоритму многорукого бандита: ID  %d", bannerID)
		a.kafkaProducer.SendCustomMessage(err, successMessage, "getBannerByMultiArmBandit", a.logger)

		return banner, nil
	}

	return nil, fmt.Errorf("banner group stats not found")
}
