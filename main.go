package main

import (
	"fmt"
	"log"
	"os"

	config "github.com/casperwein/go-edspert/album-app/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Get the environment
	env := os.Getenv("ENV")
	if env == "production" || env == "staging" {
		// Set to release mode
		gin.SetMode(gin.ReleaseMode)
	} else {
		// Get the config from .env file
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	// Initialize gin
	r := gin.Default()

	port := os.Getenv("PORT")

	// Load db config
	db, err := config.OpenDB(os.Getenv("POSTGRES_URL"), true)
	if err != nil {
		log.Fatalln(err)
	}
	defer config.CloseDB(db)

	// Load redis
	cache, err := config.OpenCache(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatalln(err)
	}

	// album
	albumRepository := config.InitAlbumRepository(db, cache)
	albumUsecase := config.InitAlbumUsecase(albumRepository.AlbumRepository)
	albumHandler := config.InitAlbumHandler(albumUsecase.AlbumUsecase)

	// album API
	albumRoutes := r.Group("/api/v1/albums")
	{
		albumRoutes.GET("/", albumHandler.AlbumHandler.GetAllAlbum)
		albumRoutes.POST("/", albumHandler.AlbumHandler.Create)
		albumRoutes.POST("/batch", albumHandler.AlbumHandler.BatchCreate)
		albumRoutes.GET("/:id", albumHandler.AlbumHandler.Get)
		albumRoutes.PUT("/:id", albumHandler.AlbumHandler.Update)
		albumRoutes.DELETE("/:id", albumHandler.AlbumHandler.Delete)
	}

	// artist
	artistRepository := config.InitArtistRepository(db, cache)
	artistUsecase := config.InitArtistUsecase(artistRepository.ArtistRepository)
	artistHandler := config.InitArtistHandler(artistUsecase.ArtistUsecase)

	artistRoutes := r.Group("/api/v1/artists")
	{
		artistRoutes.GET("/", artistHandler.ArtistHandler.GetAllArtist)
		artistRoutes.POST("/", artistHandler.ArtistHandler.Create)
		artistRoutes.POST("/batch", artistHandler.ArtistHandler.BatchCreate)
		artistRoutes.GET("/:id", artistHandler.ArtistHandler.Get)
		artistRoutes.PUT("/:id", artistHandler.ArtistHandler.Update)
		artistRoutes.DELETE("/:id", artistHandler.ArtistHandler.Delete)
	}

	// Run the gin gonic in port:
	runWithPort := fmt.Sprintf("127.0.0.1:%s", port)
	r.Run(runWithPort)
}
