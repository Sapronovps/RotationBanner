package service

import (
	"fmt"
	"math"
	"sort"
	"sync"

	"github.com/Sapronovps/RotationBanner/internal/model"
)

// CalculateBannerIDByMultiArmBandit Рассчитать ID баннера по алгоритму многорукий бандит.
func CalculateBannerIDByMultiArmBandit(bannersStats []*model.BannerGroupStats) int {
	allShows := 0

	for _, stats := range bannersStats {
		shows := stats.Shows
		if shows == 0 {
			shows = 1
		}
		allShows += shows
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
				mu.Lock()
				bannerRating := calculateRating(float64(stats.Clicks), float64(stats.Shows), float64(allShows))
				weightBanners[stats.BannerID] = bannerRating
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

	fmt.Println(weightBanners)

	return calculateBannerIDByMaxWeight(weightBanners)
}

func calculateBannerIDByMaxWeight(weightBanners map[int]float64) int {
	maxWeight := 0.0
	needBannerID := 0
	sortedKeys := iterateSorted(weightBanners, func(a, b int) bool { return a < b })

	for _, bannerID := range sortedKeys {
		if needBannerID == 0 {
			needBannerID = bannerID
		}
		if weightBanners[bannerID] > maxWeight {
			maxWeight = weightBanners[bannerID]
			needBannerID = bannerID
		}
	}

	return needBannerID
}

func calculateRating(clicks, shows, allShows float64) float64 {
	if shows == 0 {
		shows = 1
	}

	return clicks/shows + math.Sqrt(2*math.Log(allShows)/shows)
}

func iterateSorted[K comparable, V any](m map[K]V, less func(a, b K) bool) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return less(keys[i], keys[j])
	})

	return keys
}
