package pengunjung

import (
	"context"
	"database/sql"
	"eduFair/domain"

	"github.com/doug-martin/goqu/v9"
)

type repository struct {
	db *goqu.Database
}

func NewRepository(con *sql.DB) domain.PengunjungRepository {
	return &repository{
		db: goqu.New("default", con),
	}
}

func (r repository) NameExist(ctx context.Context, nama string) (bool, error) {
	query := r.db.From("pengunjung").Select(goqu.COUNT("*")).Where(goqu.Ex{"nama": nama})
	var count int
	_, err := query.ScanValContext(ctx, &count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r repository) Register(ctx context.Context, data *domain.Pengunjung) error {
	executor := r.db.Insert("pengunjung").Rows(goqu.Record{
		"nama":        data.Nama,
		"no_hp":       data.Nohp,
		"tahun_lulus": data.TahunLulus,
		"asal":        data.Asal,
		"jurusan":     data.Jurusan,
		"kategori":    data.Kategori,
	}).Returning("id").Executor()
	_, err := executor.ScanStructContext(ctx, data)
	return err
}

func (r *repository) GetAll(ctx context.Context) (pengunjung []domain.Pengunjung, err error) {
	dataset := r.db.From("pengunjung").Order(goqu.I("id").Asc())

	if err := dataset.ScanStructsContext(ctx, &pengunjung); err != nil {
		return nil, err
	}
	return
}
