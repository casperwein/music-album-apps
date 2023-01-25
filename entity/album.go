package entity

type Album struct {
	ID       int64   `json:"id"`
	ArtistID int64   `json:"artist_id,string"`
	Title    string  `json:"title"`
	Price    float32 `json:"price"`
}

// type AlbumDetail struct {
// 	ID     int64 `json:"id"`
// 	Artist Artist
// 	Title  string  `json:"title"`
// 	Price  float32 `json:"price"`
// }
