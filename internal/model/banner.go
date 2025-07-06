package model

import "time"

type Banner struct {
	ID          int       `db:"id"`          // ID баннера
	Title       string    `db:"title"`       // Название баннера
	Description string    `db:"description"` // Описание баннера
	CreatedAt   time.Time `db:"created_at"`  // Дата создания
}
