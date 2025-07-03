package model

type Slot struct {
	ID          int    `db:"id"`          // ID слота
	Description string `db:"description"` // Описание слота
}
