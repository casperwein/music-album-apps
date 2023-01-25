package config

import (
	albumHandler "github.com/casperwein/go-edspert/album-app/handler/album"

	artistHandler "github.com/casperwein/go-edspert/album-app/handler/artist"
	albumUsecase "github.com/casperwein/go-edspert/album-app/usecase/album"
	artistUseCase "github.com/casperwein/go-edspert/album-app/usecase/artist"
)

type Handler struct {
	AlbumHandler albumHandler.AlbumHandler
}

type ArtistsHandler struct {
	ArtistHandler artistHandler.ArtistHandler
}

func InitAlbumHandler(albumUsecase albumUsecase.AlbumUsecase) Handler {
	return Handler{
		AlbumHandler: albumHandler.NewAlbumHandler(albumUsecase),
	}
}

func InitArtistHandler(artistUseCase artistUseCase.ArtistUsecase) ArtistsHandler {
	return ArtistsHandler{
		ArtistHandler: artistHandler.NewArtistHandler(artistUseCase),
	}
}
