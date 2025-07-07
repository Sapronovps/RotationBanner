package sql

import "github.com/Sapronovps/RotationBanner/internal/model"

func (r *BannerRepository) CreateBanner(b *model.Banner) error {
	query := `INSERT INTO banners (title, description) VALUES ($1, $2) RETURNING id, created_at`

	err := r.storage.db.QueryRow(query, b.Title, b.Description).Scan(&b.ID, &b.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *BannerRepository) GetBanner(id int) (*model.Banner, error) {
	banner := &model.Banner{}
	err := r.storage.db.Get(banner, "SELECT * FROM banners WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return banner, nil
}

func (r *BannerRepository) UpdateBanner(b *model.Banner) error {
	query := `UPDATE banners SET title=:title, description=:description WHERE id = :id`
	if _, err := r.storage.db.NamedExec(query, b); err != nil {
		return err
	}
	return nil
}

func (r *BannerRepository) DeleteBanner(id int) error {
	query := `DELETE FROM banners WHERE id = $1`
	if _, err := r.storage.db.Exec(query, id); err != nil {
		return err
	}
	return nil
}
