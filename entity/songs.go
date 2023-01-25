package entity

type Song struct {
	Id      int64  `json:"id"`
	AlbumId int64  `json:"album_id"`
	Title   string `json:"title"`
	Lirycs  string `json:"lirycs"`
}
