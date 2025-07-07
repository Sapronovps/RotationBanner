package model

import "time"

type BannerGroupStats struct {
	ID        int       `db:"id"`
	SlotID    int       `db:"slot_id"`
	BannerID  int       `db:"banner_id"`
	GroupID   int       `db:"group_id"`
	Shows     int       `db:"shows"`
	Clicks    int       `db:"clicks"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
