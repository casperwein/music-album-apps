package repository

import (
	"context"
	"database/sql"

	"github.com/casperwein/go-edspert/album-app/entity"
	"github.com/casperwein/go-edspert/album-app/repository/artist/cache"
	"github.com/casperwein/go-edspert/album-app/repository/artist/psql"

	"github.com/go-redis/redis/v8"
)

type ArtistRepository interface {
	Get(ctx context.Context, id int64) (*entity.Artist, error)
	Create(ctx context.Context, artist *entity.Artist) (int64, error)
	GetAllArtist(ctx context.Context) ([]entity.Artist, error)
	BatchCreate(ctx context.Context, artists []entity.Artist) ([]int64, error)
	Update(ctx context.Context, artist entity.Artist) error
	Delete(ctx context.Context, id int64) error

	GetArtistCache(ctx context.Context, id int64) (*entity.Artist, error)
	GetAllArtistCache(ctx context.Context) ([]entity.Artist, error)
	SetArtistCache(ctx context.Context, id int64, artist entity.Artist) error
	SetAllArtistCache(ctx context.Context, artist []entity.Artist) error
	DeleteArtistCache(ctx context.Context, id int64) error
}

type artistRepository struct {
	postgres psql.ArtistPostgres
	cache    cache.ArtistPostgres
}

// The function is to initialize the album repository
func NewArtistRepository(db *sql.DB, client *redis.Client) ArtistRepository {
	return &artistRepository{
		postgres: psql.NewArtistPostgres(db),
		cache:    cache.NewArtistRedis(client),
	}
}
