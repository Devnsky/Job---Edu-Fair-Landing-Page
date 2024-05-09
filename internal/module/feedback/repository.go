package feedback

import (
	"context"
	"database/sql"
	"eduFair/domain"

	"github.com/doug-martin/goqu/v9"
)

type repository struct {
	db *goqu.Database
}

func NewRepository(con *sql.DB) domain.FeedbackRepository {
	return &repository{
		db: goqu.New("default", con),
	}
}

func (r repository) Register(ctx context.Context, data *domain.Feedback) error {
	executor := r.db.Insert("feedback").Rows(goqu.Record{
		"feedback": data.Feedback,
	}).Returning("id").Executor()
	_, err := executor.ScanStructContext(ctx, data)
	return err
}

func (r repository) GetAll(ctx context.Context) (feed []domain.Feedback, err error) {
	dataset := r.db.From("feedback").Order(goqu.I("id").Asc())

	if err := dataset.ScanStructsContext(ctx, &feed); err != nil {
		return nil, err
	}
	return
}
