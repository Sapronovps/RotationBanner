package memory

import (
	"fmt"
	"github.com/Sapronovps/RotationBanner/internal/model"
)

func (b *BannerRepository) CreateBanner(banner *model.Banner) error {
	banner.ID = len(b.Banners) + 1
	b.mu.Lock()
	b.Banners[banner.ID] = banner
	b.mu.Unlock()
	return nil
}

func (b *BannerRepository) GetBanner(id int) (*model.Banner, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	banner, ok := b.Banners[id]
	if !ok {
		return nil, fmt.Errorf("no banner with id %d", id)
	}
	return banner, nil
}

func (b *BannerRepository) UpdateBanner(banner *model.Banner) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.Banners[banner.ID] = banner
	return nil
}

func (b *BannerRepository) DeleteBanner(id int) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	delete(b.Banners, id)
	return nil
}
