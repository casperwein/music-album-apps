package cache

import (
	"context"

	"github.com/casperwein/go-edspert/album-app/entity"

	"github.com/go-redis/redis/v8"
)

type ArtistPostgres interface {
	GetArtist(ctx context.Context, id int64) (*entity.Artist, error)
	GetAllArtist(ctx context.Context) ([]entity.Artist, error)
	SetArtist(ctx context.Context, id int64, artist entity.Artist) error
	SetAllArtist(ctx context.Context, artists []entity.Artist) error
	Delete(ctx context.Context, id int64) error
}

type artistConnection struct {
	client *redis.Client
}

// The function is to initialize the artist psql repository
func NewArtistRedis(cache *redis.Client) ArtistPostgres {
	return &artistConnection{client: cache}
}
