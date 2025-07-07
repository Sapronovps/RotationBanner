package sql

import "github.com/Sapronovps/RotationBanner/internal/model"

func (r *BannerRepository) CreateSlot(s *model.Slot) error {
	query := `INSERT INTO slots (description) VALUES ($1) RETURNING id, created_at`

	err := r.storage.db.QueryRow(query, s.Description).Scan(&s.ID, &s.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *BannerRepository) GetSlot(id int) (*model.Slot, error) {
	slot := &model.Slot{}
	err := r.storage.db.Get(slot, "SELECT * FROM slots WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return slot, nil
}
