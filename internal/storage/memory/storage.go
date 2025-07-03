package memory

import (
	"github.com/Sapronovps/RotationBanner/internal/model"
	"github.com/Sapronovps/RotationBanner/internal/storage"
)

type Storage struct {
	bannerRepository storage.BannerRepository
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Banner() storage.BannerRepository {
	if s.bannerRepository != nil {
		return s.bannerRepository
	}
	s.bannerRepository = &BannerRepository{
		Slots:            make(map[int]*model.Slot),
		Banners:          make(map[int]*model.Banner),
		Groups:           make(map[int]*model.Group),
		BannerGroupStats: make(map[int]*model.BannerGroupStats),
	}
	return s.bannerRepository
}
