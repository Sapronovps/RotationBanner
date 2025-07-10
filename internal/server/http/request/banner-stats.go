package request

type BannerStatsRequest struct {
	SlotID   int
	GroupID  int
	BannerID int
}

type GetBannerRequest struct {
	SlotID  int
	GroupID int
}
