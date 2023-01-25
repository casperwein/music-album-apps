package cache

import "time"

const (
	artistKey       = "artists"
	artistDetailKey = "artists:%d"
	expiration      = time.Hour * 1
)
