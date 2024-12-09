package dao

import (
	"context"
	"github.com/zhuguangfeng/go-chat/model"
)

type ActivitySignUp interface {
	InsertActivitySignUp(ctx context.Context, signUp model.ActivitySignUp) error
	UpdateActivitySignUp(ctx context.Context, signUp model.ActivitySignUp) error
	ListActivitySignUp(ctx context.Context, pageNum, pageSize int, status uint) ([]model.ActivitySignUp, error)
}

type GormActivitySignUp struct {
}

func (dao GormActivitySignUp) InsertActivitySignUp(ctx context.Context, signUp model.ActivitySignUp) error {
	//TODO implement me
	panic("implement me")
}

func (dao GormActivitySignUp) UpdateActivitySignUp(ctx context.Context, signUp model.ActivitySignUp) error {
	//TODO implement me
	panic("implement me")
}

func (dao GormActivitySignUp) ListActivitySignUp(ctx context.Context, pageNum, pageSize int, status uint) ([]model.ActivitySignUp, error) {
	//TODO implement me
	panic("implement me")
}

func NewActivitySignUp() ActivitySignUp {
	return &GormActivitySignUp{}
}
