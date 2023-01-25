package repository

import (
	"context"

	"github.com/casperwein/go-edspert/album-app/entity"
)

// ! ganti dulu nama albumnya

// It will call the function Create in psql/artist
func (repo *artistRepository) Create(ctx context.Context, artist *entity.Artist) (int64, error) {
	return repo.postgres.Create(ctx, artist)
}

// It will call the function Get in psql/artist
func (repo *artistRepository) Get(ctx context.Context, id int64) (*entity.Artist, error) {
	return repo.postgres.Get(ctx, id)
}

// It will call the function GetAllartist in psql/artist
func (repo *artistRepository) GetAllArtist(ctx context.Context) ([]entity.Artist, error) {
	return repo.postgres.GetAllArtist(ctx)
}

// It will call the function BatchCreate in psql/artist
func (repo *artistRepository) BatchCreate(ctx context.Context, artists []entity.Artist) ([]int64, error) {
	return repo.postgres.BatchCreate(ctx, artists)
}

// It will call the function Update in psql/artist
func (repo *artistRepository) Update(ctx context.Context, artist entity.Artist) error {
	return repo.postgres.Update(ctx, artist)
}

// It will call the function Delete in psql/artist
func (repo *artistRepository) Delete(ctx context.Context, id int64) error {
	return repo.postgres.Delete(ctx, id)
}

func (repo *artistRepository) GetArtistCache(ctx context.Context, id int64) (*entity.Artist, error) {
	return repo.cache.GetArtist(ctx, id)
}

func (repo *artistRepository) GetAllArtistCache(ctx context.Context) ([]entity.Artist, error) {
	return repo.cache.GetAllArtist(ctx)
}

func (repo *artistRepository) SetArtistCache(ctx context.Context, id int64, artist entity.Artist) error {
	return repo.cache.SetArtist(ctx, id, artist)
}

func (repo *artistRepository) SetAllArtistCache(ctx context.Context, artists []entity.Artist) error {
	return repo.cache.SetAllArtist(ctx, artists)
}

func (repo *artistRepository) DeleteArtistCache(ctx context.Context, id int64) error {
	return repo.cache.Delete(ctx, id)
}
