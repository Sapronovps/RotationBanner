package sql

import (
	"github.com/Sapronovps/RotationBanner/internal/storage"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq" // for postgres
)

type Storage struct {
	db               *sqlx.DB
	bannerRepository *storage.BannerRepository
}

func New(db *sqlx.DB) *Storage {
	return &Storage{db: db}
}

func NewDB(dsn string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err)
	}
	return db
}

func (s *Storage) Banner() storage.BannerRepository {
	if s.bannerRepository != nil {
		return *s.bannerRepository
	}
	return &BannerRepository{storage: s}
}
