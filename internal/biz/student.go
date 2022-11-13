package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Student struct {
	Name string
	Age  int
}

type StudentRepo interface {
	Save(context.Context, *Student) (*Student, error)
}
type StudentUsecase struct {
	repo StudentRepo
	log  *log.Helper
}

func NewStudentUsecase(repo StudentRepo, logger log.Logger) *StudentUsecase {
	return &StudentUsecase{repo: repo, log: log.NewHelper(logger)}
}
