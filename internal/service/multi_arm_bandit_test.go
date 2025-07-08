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
			name: "Когда у всех баннеров 0 просмотров и кликов",
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
		{
			name: "Когда у всех баннеров разное количество просмотров и нет кликов",
			bannerGroupStats: []*model.BannerGroupStats{
				{SlotID: 1, BannerID: 1, GroupID: 1, Clicks: 0, Shows: 3},
				{SlotID: 1, BannerID: 2, GroupID: 1, Clicks: 0, Shows: 4},
				{SlotID: 1, BannerID: 3, GroupID: 1, Clicks: 0, Shows: 2},
			},
			expectedBannerId: 3,
		},
		{
			name: "Когда у одного баннера есть клик",
			bannerGroupStats: []*model.BannerGroupStats{
				{SlotID: 1, BannerID: 1, GroupID: 1, Clicks: 1, Shows: 5},
				{SlotID: 1, BannerID: 2, GroupID: 1, Clicks: 0, Shows: 4},
				{SlotID: 1, BannerID: 3, GroupID: 1, Clicks: 0, Shows: 4},
			},
			expectedBannerId: 1,
		},
		{
			name: "Когда у одного баннера есть клик, но у него много просмотров",
			bannerGroupStats: []*model.BannerGroupStats{
				{SlotID: 1, BannerID: 1, GroupID: 1, Clicks: 1, Shows: 6},
				{SlotID: 1, BannerID: 2, GroupID: 1, Clicks: 0, Shows: 4},
				{SlotID: 1, BannerID: 3, GroupID: 1, Clicks: 0, Shows: 4},
			},
			expectedBannerId: 2,
		},
		{
			name: "Когда у двоих баннеров есть клик",
			bannerGroupStats: []*model.BannerGroupStats{
				{SlotID: 1, BannerID: 1, GroupID: 1, Clicks: 2, Shows: 7},
				{SlotID: 1, BannerID: 2, GroupID: 1, Clicks: 1, Shows: 4},
				{SlotID: 1, BannerID: 3, GroupID: 1, Clicks: 0, Shows: 4},
			},
			expectedBannerId: 2,
		},
		{
			name: "Когда у всех баннеров есть клик",
			bannerGroupStats: []*model.BannerGroupStats{
				{SlotID: 1, BannerID: 1, GroupID: 1, Clicks: 2, Shows: 7},
				{SlotID: 1, BannerID: 2, GroupID: 1, Clicks: 1, Shows: 5},
				{SlotID: 1, BannerID: 3, GroupID: 1, Clicks: 1, Shows: 4},
			},
			expectedBannerId: 3,
		},
		{
			name: "Когда у одного баннера много просмотров и кликов",
			bannerGroupStats: []*model.BannerGroupStats{
				{SlotID: 1, BannerID: 1, GroupID: 1, Clicks: 799, Shows: 16000},
				{SlotID: 1, BannerID: 2, GroupID: 1, Clicks: 59, Shows: 9000},
				{SlotID: 1, BannerID: 3, GroupID: 1, Clicks: 9, Shows: 3000},
			},
			expectedBannerId: 1,
		},
		{
			name: "Когда у всех баннеров одинаковое количество просмотров и кликов",
			bannerGroupStats: []*model.BannerGroupStats{
				{SlotID: 1, BannerID: 1, GroupID: 1, Clicks: 10, Shows: 50},
				{SlotID: 1, BannerID: 2, GroupID: 1, Clicks: 10, Shows: 50},
				{SlotID: 1, BannerID: 3, GroupID: 1, Clicks: 10, Shows: 50},
			},
			expectedBannerId: 1,
		},
		{
			name: "Когда у одного баннера больше всех просмотров и мало кликов",
			bannerGroupStats: []*model.BannerGroupStats{
				{SlotID: 1, BannerID: 1, GroupID: 1, Clicks: 50, Shows: 1000},
				{SlotID: 1, BannerID: 2, GroupID: 1, Clicks: 60, Shows: 1500},
				{SlotID: 1, BannerID: 3, GroupID: 1, Clicks: 100, Shows: 800},
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
