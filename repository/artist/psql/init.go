package psql

import (
	"context"
	"database/sql"

	"github.com/casperwein/go-edspert/album-app/entity"
)

type ArtistPostgres interface {
	Get(ctx context.Context, id int64) (*entity.Artist, error)
	Create(ctx context.Context, artist *entity.Artist) (int64, error)
	GetAllArtist(ctx context.Context) ([]entity.Artist, error)
	BatchCreate(ctx context.Context, artists []entity.Artist) ([]int64, error)
	Update(ctx context.Context, artist entity.Artist) error
	Delete(ctx context.Context, id int64) error
}

type artistConnection struct {
	db *sql.DB
}

func NewArtistPostgres(db *sql.DB) ArtistPostgres {
	return &artistConnection{db: db}
}
