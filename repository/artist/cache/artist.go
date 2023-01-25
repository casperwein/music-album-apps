package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/casperwein/go-edspert/album-app/entity"

	"github.com/go-redis/redis/v8"
)

// Get Specific artist cache
func (repo *artistConnection) GetArtist(ctx context.Context, id int64) (*entity.Artist, error) {
	var artist entity.Artist

	key := fmt.Sprintf(artistDetailKey, id)

	artistString, err := repo.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return &artist, nil
	}
	if err != nil {
		return &artist, err
	}

	err = json.Unmarshal([]byte(artistString), &artist)
	if err != nil {
		return &artist, err
	}

	return &artist, nil
}

// GetAllArtists is function to get all artists from database
func (repo *artistConnection) GetAllArtist(ctx context.Context) ([]entity.Artist, error) {
	var artists []entity.Artist

	artistsString, err := repo.client.Get(ctx, artistDetailKey).Result()
	if err == redis.Nil {
		return artists, nil
	}
	if err != nil {
		return artists, err
	}

	err = json.Unmarshal([]byte(artistsString), &artists)
	if err != nil {
		return artists, err
	}

	return artists, nil
}

func (repo *artistConnection) SetArtist(ctx context.Context, id int64, artist entity.Artist) error {
	key := fmt.Sprintf(artistDetailKey, id)

	artistsString, err := json.Marshal(artist)
	if err != nil {
		return err
	}

	if err := repo.client.Set(ctx, key, artistsString, expiration).Err(); err != nil {
		return err
	}

	return nil
}

func (repo *artistConnection) SetAllArtist(ctx context.Context, artist []entity.Artist) error {
	artistsString, err := json.Marshal(artist)
	if err != nil {
		return err
	}

	if err := repo.client.Set(ctx, artistKey, artistsString, expiration).Err(); err != nil {
		return err
	}

	return nil
}

func (repo *artistConnection) Delete(ctx context.Context, id int64) error {
	key := fmt.Sprintf(artistDetailKey, id)

	if err := repo.client.Del(ctx, key).Err(); err != nil {
		return err
	}

	return nil
}
