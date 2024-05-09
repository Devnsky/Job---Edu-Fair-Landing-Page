package domain

import "context"

type Pengunjung struct {
	ID         int64  `db:"id"`
	Nama       string `db:"nama"`
	Nohp       string `db:"no_hp"`
	TahunLulus string `db:"tahun_lulus"`
	Asal       string `db:"asal"`
	Jurusan    string `db:"jurusan"`
	Kategori   string `db:"kategori"`
}

type PengunjungRepository interface {
	GetAll(ctx context.Context) ([]Pengunjung, error)
	Register(ctx context.Context, data *Pengunjung) error
	NameExist(ctx context.Context, nama string) (bool, error)
}

type PengunjungService interface {
	GetAll(ctx context.Context) ApiResponse
	Register(ctx context.Context, data PengunjungData) ApiResponse
}

type PengunjungData struct {
	Nama       string `json:"nama"`
	Nohp       string `json:"no_hp"`
	TahunLulus string `json:"tahun_lulus"`
	Asal       string `json:"asal"`
	Jurusan    string `json:"jurusan"`
	Kategori   string `json:"kategori"`
}
