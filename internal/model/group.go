package model

type Group struct {
	ID          int    `db:"id"`
	Title       string `db:"name"`
	Description string `db:"description"`
}
