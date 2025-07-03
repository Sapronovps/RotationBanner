package memory

import (
	"github.com/Sapronovps/RotationBanner/internal/model"
	"sync"
)

type BannerRepository struct {
	mu               sync.RWMutex
	Slots            map[int]*model.Slot
	Banners          map[int]*model.Banner
	Groups           map[int]*model.Group
	BannerGroupStats map[int]*model.BannerGroupStats
}
