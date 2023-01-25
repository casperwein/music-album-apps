package usecase

import (
	"context"

	"github.com/casperwein/go-edspert/album-app/entity"
	artistRepository "github.com/casperwein/go-edspert/album-app/repository/artist"
)

type ArtistUsecase interface {
	Get(ctx context.Context, id int64) (*entity.Artist, error)
	Create(ctx context.Context, artist *entity.Artist) (*entity.Artist, error)
	GetAllArtist(ctx context.Context) ([]entity.Artist, error)
	BatchCreate(ctx context.Context, artists []entity.Artist) ([]entity.Artist, error)
	Update(ctx context.Context, artist entity.Artist) (entity.Artist, error)
	Delete(ctx context.Context, id int64) error
}

type artistUsecase struct {
	artistRepository artistRepository.ArtistRepository
}

// The function is to initialize the album usecase
func NewArtistUsecase(artistRepository artistRepository.ArtistRepository) ArtistUsecase {
	return &artistUsecase{
		artistRepository: artistRepository,
	}
}
