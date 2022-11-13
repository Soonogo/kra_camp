package service

import (
	"context"
	"fmt"

	pb "kra/api/student/v1"
	"kra/internal/biz"
)

type StudentService struct {
	pb.UnimplementedStudentServer

	uc *biz.StudentUsecase
}

func NewStudentService(uc *biz.StudentUsecase) *StudentService {
	return &StudentService{uc: uc}
}

func (s *StudentService) CallStudent(ctx context.Context, req *pb.StudentRequest) (*pb.StudentReply, error) {
	msg := fmt.Sprintf("姓名：%s ,年龄：%d", req.Name, req.Age)
	return &pb.StudentReply{Message: msg}, nil
}
