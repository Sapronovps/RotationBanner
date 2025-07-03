package model

type Banner struct {
	ID          int    `db:"id"`          // ID баннера
	Title       string `db:"title"`       // Название баннера
	Description string `db:"description"` // Описание баннера
}
