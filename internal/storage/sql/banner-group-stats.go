package sql

import "github.com/Sapronovps/RotationBanner/internal/model"

func (r *BannerRepository) CreateBannerGroupStats(s *model.BannerGroupStats) error {
	query := `INSERT INTO banner_group_stats (
                                slot_id, 
                                banner_id, 
                                group_id,
                                clicks,
                                shows
                                ) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`

	return r.storage.db.QueryRow(query, s.SlotID, s.BannerID, s.GroupID, s.Clicks, s.Shows).Scan(&s.ID, &s.CreatedAt)
}

func (r *BannerRepository) GetBannerGroupStats(slotID, bannerID, groupID int) (*model.BannerGroupStats, error) {
	bannerGroupStats := &model.BannerGroupStats{}
	err := r.storage.db.Get(bannerGroupStats, "SELECT * FROM banner_group_stats"+
		" WHERE slot_id = $1 AND banner_id = $2 AND group_id = $3", slotID, bannerID, groupID)
	if err != nil {
		return nil, err
	}
	return bannerGroupStats, nil
}

func (r *BannerRepository) UpdateBannerGroupStats(s *model.BannerGroupStats) error {
	query := `UPDATE banner_group_stats SET 
                              slot_id=:slot_id, 
                              banner_id=:banner_id, 
                              group_id=:group_id, 
                              clicks=:clicks, 
                              shows=:shows,
                              updated_at=:updated_at
                          WHERE id = :id`
	if _, err := r.storage.db.NamedExec(query, s); err != nil {
		return err
	}
	return nil
}

func (r *BannerRepository) GetBannersGroupStats(slotId, groupId int) (bannersStats []*model.BannerGroupStats) {
	var bannersGroupStats []*model.BannerGroupStats
	_ = r.storage.db.Select(&bannersGroupStats,
		"SELECT * FROM banner_group_stats WHERE slot_id = $1 AND group_id = $2", slotId, groupId)

	return bannersGroupStats
}
