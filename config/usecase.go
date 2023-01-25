package config

import (
	albumRepository "github.com/casperwein/go-edspert/album-app/repository/album"
	artistRepository "github.com/casperwein/go-edspert/album-app/repository/artist"
	albumUsecase "github.com/casperwein/go-edspert/album-app/usecase/album"
	artistUsecase "github.com/casperwein/go-edspert/album-app/usecase/artist"
)

type Usecase struct {
	AlbumUsecase albumUsecase.AlbumUsecase
}

type Artist_Usecase struct {
	ArtistUsecase artistUsecase.ArtistUsecase
}

// Function to initialize usecase
func InitAlbumUsecase(albumRepository albumRepository.AlbumRepository) Usecase {
	return Usecase{
		AlbumUsecase: albumUsecase.NewAlbumUsecase(albumRepository),
	}
}

func InitArtistUsecase(artistRepository artistRepository.ArtistRepository) Artist_Usecase {
	return Artist_Usecase{
		ArtistUsecase: artistUsecase.NewArtistUsecase(artistRepository),
	}
}
