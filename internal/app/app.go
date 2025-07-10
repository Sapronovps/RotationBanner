package app

import (
	"fmt"
	"github.com/Sapronovps/RotationBanner/internal/model"
	"github.com/Sapronovps/RotationBanner/internal/service"
	"github.com/Sapronovps/RotationBanner/internal/storage"
	"go.uber.org/zap"
	"sync"
	"time"
)

type App struct {
	mu      sync.Mutex
	logger  *zap.Logger
	storage storage.Storage
}

func NewApp(logger *zap.Logger, storage storage.Storage) *App {
	return &App{logger: logger, storage: storage}
}

func (a *App) AddSlot(slot *model.Slot) error {
	return a.storage.Banner().CreateSlot(slot)
}

func (a *App) GetSlot(id int) (slot *model.Slot, err error) {
	return a.storage.Banner().GetSlot(id)
}

func (a *App) AddBanner(banner *model.Banner) error {
	return a.storage.Banner().CreateBanner(banner)
}

func (a *App) GetBanner(id int) (banner *model.Banner, err error) {
	return a.storage.Banner().GetBanner(id)
}

func (a *App) AddGroup(group *model.Group) error {
	return a.storage.Banner().CreateGroup(group)
}

func (a *App) GetGroup(id int) (group *model.Group, err error) {
	return a.storage.Banner().GetGroup(id)
}

func (a *App) AddBannerGroupStats(group *model.BannerGroupStats) error {
	stat, err := a.GetBannerGroupStats(group.SlotID, group.BannerID, group.GroupID)
	if err == nil {
		group.ID = stat.ID
		group.Shows = stat.Shows
		group.Clicks = stat.Clicks
		group.CreatedAt = stat.CreatedAt
		group.UpdatedAt = stat.UpdatedAt
		return nil
	}

	return a.storage.Banner().CreateBannerGroupStats(group)
}

func (a *App) GetBannerGroupStats(slotID, bannerID, groupID int) (*model.BannerGroupStats, error) {
	return a.storage.Banner().GetBannerGroupStats(slotID, bannerID, groupID)
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

	return a.storage.Banner().UpdateBannerGroupStats(stats)
}

func (a *App) GetBannerByMultiArmBandit(slotID, groupID int) (banner *model.Banner, err error) {
	bannersStats := a.storage.Banner().GetBannersGroupStats(slotID, groupID)
	if bannersStats == nil {
		return nil, fmt.Errorf("banner group stats not found slot id: %d and group id: %d", slotID, groupID)
	}

	bannerID := service.CalculateBannerIdByMultiArmBandit(bannersStats)

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

		return banner, nil
	}

	return nil, fmt.Errorf("banner group stats not found")
}
