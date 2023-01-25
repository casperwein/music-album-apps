package config

import (
	"database/sql"

	albumRepository "github.com/casperwein/go-edspert/album-app/repository/album"
	artistRepository "github.com/casperwein/go-edspert/album-app/repository/artist"

	"github.com/go-redis/redis/v8"
)

type Repository struct {
	AlbumRepository albumRepository.AlbumRepository
}

type ArtistRepository struct {
	ArtistRepository artistRepository.ArtistRepository
}

// Function to initialize repository
func InitAlbumRepository(db *sql.DB, cache *redis.Client) Repository {
	return Repository{
		AlbumRepository: albumRepository.NewAlbumRepository(db, cache),
	}
}

func InitArtistRepository(db *sql.DB, cache *redis.Client) ArtistRepository {
	return ArtistRepository{
		ArtistRepository: artistRepository.NewArtistRepository(db, cache),
	}
}
