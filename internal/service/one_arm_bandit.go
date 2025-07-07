package service

import (
	"github.com/Sapronovps/RotationBanner/internal/model"
	"math"
	"sync"
)

// CalculateBannerIdByOneArmBandit Рассчитать ID баннера по алгоритму Однорукий бандит.
func CalculateBannerIdByOneArmBandit(bannersStats []*model.BannerGroupStats) int {
	allShows := 0

	for _, stats := range bannersStats {
		allShows += stats.Shows
	}

	// Расчеты выполним параллельно

	weightBanners := make(map[int]float64)
	workers := 5
	jobs := make(chan *model.BannerGroupStats, len(bannersStats))
	wg := sync.WaitGroup{}
	var mu sync.Mutex

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for stats := range jobs {
				numerator := float64(stats.Clicks) * math.Log(float64(allShows))
				fraction := numerator / float64(stats.Shows)
				sqrtVal := math.Sqrt(fraction)
				result := float64(stats.Clicks) + sqrtVal
				mu.Lock()
				weightBanners[stats.BannerID] = result
				mu.Unlock()
			}
		}()
	}

	// Отправляем баннеры в канал
	for _, bannerStat := range bannersStats {
		jobs <- bannerStat
	}
	close(jobs)

	wg.Wait()

	return calculateBannerIdByMaxWeight(weightBanners)
}

func calculateBannerIdByMaxWeight(weightBanners map[int]float64) int {
	maxWeight := 0.0
	needBannerID := 0

	for bannerID, weight := range weightBanners {
		if weight > maxWeight {
			maxWeight = weight
			needBannerID = bannerID
		}
	}

	return needBannerID
}
