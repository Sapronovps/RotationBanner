package app

import (
	"fmt"
	"github.com/Sapronovps/RotationBanner/internal/model"
	"github.com/Sapronovps/RotationBanner/internal/service"
	"github.com/Sapronovps/RotationBanner/internal/storage"
	"go.uber.org/zap"
	"sync"
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

func (a *App) CreateGroup(group *model.Group) error {
	return a.storage.Banner().CreateGroup(group)
}

func (a *App) GetGroup(id int) (group *model.Group, err error) {
	return a.storage.Banner().GetGroup(id)
}

func (a *App) CreateBannerGroupStats(group *model.BannerGroupStats) error {
	return a.storage.Banner().CreateBannerGroupStats(group)
}

func (a *App) GetBannerGroupStats(slotID, bannerID, groupID int) (*model.BannerGroupStats, error) {
	return a.storage.Banner().GetBannerGroupStats(slotID, bannerID, groupID)
}

func (a *App) RegisterClick(slotID, bannerID, groupID int) error {
	stats, err := a.storage.Banner().GetBannerGroupStats(slotID, bannerID, groupID)
	if err != nil {
		return err
	}
	stats.Clicks++

	return nil
}

func (a *App) GetAndUpdateBanner(slotID, groupID int) (banner *model.Banner, err error) {
	bannersStats := a.storage.Banner().GetBannersGroupStats(slotID, groupID)
	if bannersStats == nil {
		return nil, fmt.Errorf("banner group stats not found")
	}

	bannerID := service.CalculateBannerIdByOneArmBandit(bannersStats)

	if bannerID > 0 {
		return a.storage.Banner().GetBanner(bannerID)
	}

	return nil, fmt.Errorf("banner group stats not found")
}

//func (a *App) RemoveBanner(id int) error {
//	return a.storage.Banner().DeleteBanner(id)
//}

//func (a *App) AttachBannerToSlot(slotID, bannerID int) error {
//	slot, err := a.GetSlot(slotID)
//	if err != nil {
//		return fmt.Errorf("could not find slot with id %d: %w", slotID, err)
//	}
//
//	banner, err := a.GetBanner(bannerID)
//	if err != nil {
//		return fmt.Errorf("could not find banner with id %d: %w", bannerID, err)
//	}
//	if slot.Banners == nil {
//		slot.Banners = make(map[int]*model.Banner)
//	}
//
//	slot.Banners[bannerID] = banner
//	return nil
//}

//func (a *App) DetachBannerFromSlot(slotID, bannerID int) error {
//	slot, err := a.GetSlot(slotID)
//	if err != nil {
//		return fmt.Errorf("could not find slot with id %d: %w", slotID, err)
//	}
//	_, ok := slot.Banners[bannerID]
//	if !ok {
//		return fmt.Errorf("no banner found with id %d", bannerID)
//	}
//	delete(slot.Banners, bannerID)
//	return nil
//}

//func (a *App) RegisterClick(bannerID int) error {
//	a.mu.Lock()
//	defer a.mu.Unlock()
//
//	banner, err := a.GetBanner(bannerID)
//	if err != nil {
//		return fmt.Errorf("could not find banner with id %d: %w", bannerID, err)
//	}
//
//	banner.Clicks++
//	banner.Weight = float64(banner.Clicks) / float64(banner.Shows) // Обновляем CTR
//	return nil
//}
