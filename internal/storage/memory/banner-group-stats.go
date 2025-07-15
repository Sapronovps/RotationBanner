package memory

import (
	"fmt"

	"github.com/Sapronovps/RotationBanner/internal/model"
)

func (b *BannerRepository) CreateBannerGroupStats(stats *model.BannerGroupStats) error {
	stats.ID = len(b.BannerGroupStats) + 1
	b.mu.Lock()
	b.BannerGroupStats[stats.ID] = stats
	b.mu.Unlock()
	return nil
}

func (b *BannerRepository) GetBannerGroupStats(slotID, bannerID, groupID int) (*model.BannerGroupStats, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, stats := range b.BannerGroupStats {
		if stats.SlotID == slotID && stats.BannerID == bannerID && stats.GroupID == groupID {
			return stats, nil
		}
	}
	return nil, fmt.Errorf("banner group stats not found")
}

func (b *BannerRepository) GetBannersGroupStats(slotID, groupID int) []*model.BannerGroupStats {
	b.mu.RLock()
	defer b.mu.RUnlock()

	bannersStats := make([]*model.BannerGroupStats, 0)
	for _, stats := range b.BannerGroupStats {
		if stats.SlotID == slotID && stats.GroupID == groupID {
			bannersStats = append(bannersStats, stats)
		}
	}
	return bannersStats
}

func (b *BannerRepository) UpdateBannerGroupStats(s *model.BannerGroupStats) error {
	b.mu.RLock()
	defer b.mu.RUnlock()

	b.BannerGroupStats[s.ID] = s

	return nil
}
