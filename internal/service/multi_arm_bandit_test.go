package service

import (
	"github.com/Sapronovps/RotationBanner/internal/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculateBanner(t *testing.T) {
	tests := []struct {
		name             string
		bannerGroupStats []*model.BannerGroupStats
		expectedBannerId int
	}{
		{
			name: "У всех баннеров 0 просмотров и кликов",
			bannerGroupStats: []*model.BannerGroupStats{
				{SlotID: 1, BannerID: 1, GroupID: 1, Clicks: 0, Shows: 0},
				{SlotID: 1, BannerID: 2, GroupID: 1, Clicks: 0, Shows: 0},
				{SlotID: 1, BannerID: 3, GroupID: 1, Clicks: 0, Shows: 0},
			},
			expectedBannerId: 1,
		},
		{
			name: "Когда у одного баннера есть просмотры, но нет кликов",
			bannerGroupStats: []*model.BannerGroupStats{
				{SlotID: 1, BannerID: 1, GroupID: 1, Clicks: 0, Shows: 2},
				{SlotID: 1, BannerID: 2, GroupID: 1, Clicks: 0, Shows: 0},
				{SlotID: 1, BannerID: 3, GroupID: 1, Clicks: 0, Shows: 0},
			},
			expectedBannerId: 2,
		},
		{
			name: "Когда у двоих баннеров есть просмотры, но нет кликов",
			bannerGroupStats: []*model.BannerGroupStats{
				{SlotID: 1, BannerID: 1, GroupID: 1, Clicks: 0, Shows: 2},
				{SlotID: 1, BannerID: 2, GroupID: 1, Clicks: 0, Shows: 2},
				{SlotID: 1, BannerID: 3, GroupID: 1, Clicks: 0, Shows: 0},
			},
			expectedBannerId: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bannerID := CalculateBannerIdByMultiArmBandit(test.bannerGroupStats)
			require.Equal(t, test.expectedBannerId, bannerID)
		})
	}
}
