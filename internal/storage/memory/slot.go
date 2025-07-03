package memory

import (
	"fmt"
	"github.com/Sapronovps/RotationBanner/internal/model"
)

func (b *BannerRepository) CreateSlot(slot *model.Slot) error {
	slot.ID = len(b.Slots) + 1
	b.mu.Lock()
	b.Slots[slot.ID] = slot
	b.mu.Unlock()
	return nil
}

func (b *BannerRepository) GetSlot(id int) (*model.Slot, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	slot, ok := b.Slots[id]
	if !ok {
		return nil, fmt.Errorf("no slot with id %d", id)
	}
	return slot, nil
}
