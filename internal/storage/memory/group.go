package memory

import (
	"fmt"

	"github.com/Sapronovps/RotationBanner/internal/model"
)

func (b *BannerRepository) CreateGroup(group *model.Group) error {
	group.ID = len(b.Groups) + 1
	b.mu.Lock()
	b.Groups[group.ID] = group
	b.mu.Unlock()
	return nil
}

func (b *BannerRepository) GetGroup(id int) (*model.Group, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	group, ok := b.Groups[id]
	if !ok {
		return nil, fmt.Errorf("no group with id %d", id)
	}
	return group, nil
}
