package storage

import "github.com/Sapronovps/RotationBanner/internal/model"

type BannerRepository interface {
	CreateSlot(s *model.Slot) error                                                     // Создание слота
	GetSlot(id int) (*model.Slot, error)                                                // Получение слота
	CreateBanner(b *model.Banner) error                                                 // Создание баннера
	GetBanner(id int) (*model.Banner, error)                                            // Получение баннера
	UpdateBanner(b *model.Banner) error                                                 // Обновление баннера
	DeleteBanner(id int) error                                                          // Удаление баннера
	CreateGroup(g *model.Group) error                                                   // Создание группы
	GetGroup(id int) (*model.Group, error)                                              // Получение группы
	CreateBannerGroupStats(stats *model.BannerGroupStats) error                         // Создание связки Slot -> Banner -> Group
	GetBannerGroupStats(slotID, bannerID, groupID int) (*model.BannerGroupStats, error) // Получение связки
	UpdateBannerGroupStats(stats *model.BannerGroupStats) error                         // Обновление статистики
	GetBannersGroupStats(slotId, groupId int) (bannersStats []*model.BannerGroupStats)  // Получение связок по-заданному ID слота и ID группы
}
