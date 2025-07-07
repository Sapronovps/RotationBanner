package sql

import "github.com/Sapronovps/RotationBanner/internal/model"

func (r *BannerRepository) CreateGroup(g *model.Group) error {
	query := `INSERT INTO groups (title, description) VALUES ($1, $2) RETURNING id, created_at`

	return r.storage.db.QueryRow(query, g.Title, g.Description).Scan(&g.ID, &g.CreatedAt)
}

func (r *BannerRepository) GetGroup(id int) (*model.Group, error) {
	group := &model.Group{}
	err := r.storage.db.Get(group, "SELECT * FROM groups WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return group, nil
}
