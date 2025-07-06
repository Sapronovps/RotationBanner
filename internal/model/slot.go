package model

import "time"

type Slot struct {
	ID          int       `db:"id"`          // ID слота
	Description string    `db:"description"` // Описание слота
	CreatedAt   time.Time `db:"created_at"`  // Дата создания слота
}
