package data

import (
	"context"
	"kra/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type StudentRepo struct {
	data *Data
	log  *log.Helper
}

func NewStudentRepo(data *Data, logger log.Logger) biz.StudentRepo {
	return &StudentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *StudentRepo) Save(ctx context.Context, g *biz.Student) (*biz.Student, error) {
	return g, nil
}
