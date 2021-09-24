package domain

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
)

type Service interface {
	ListMatrix(context.Context) ([]Matrix, error)
	CreateMatrix(context.Context, *Matrix) (Matrix, error)
	FindMatrix(context.Context, string) (Matrix, error)
	UpdateMatrix(context.Context, *Matrix) (Matrix, error)
	DeleteMatrix(context.Context, string) error
}

type service struct {
	repo   Repository
	logger log.Logger
}

func NewService(repo Repository, logger log.Logger) *service { // nolint: revive
	return &service{
		repo:   repo,
		logger: logger,
	}
}

func (s *service) ListMatrix(_ context.Context) ([]Matrix, error) {
	ms, err := s.repo.List()
	if err != nil {
		return []Matrix{}, fmt.Errorf("service didn't found any matrix: %w", err)
	}
	return ms, nil
}

func (s *service) CreateMatrix(_ context.Context, matrix *Matrix) (Matrix, error) {
	m, err := s.repo.Create(matrix)
	if err != nil {
		return Matrix{}, fmt.Errorf("service can't create matrix: %w", err)
	}
	return m, nil
}

func (s *service) FindMatrix(_ context.Context, id string) (Matrix, error) {
	m, err := s.repo.Find(id)
	if err != nil {
		return Matrix{}, fmt.Errorf("service can't find matrix: %w", err)
	}
	return m, nil
}

func (s *service) UpdateMatrix(_ context.Context, matrix *Matrix) (Matrix, error) {
	m, err := s.repo.Update(matrix)
	if err != nil {
		return Matrix{}, fmt.Errorf("service can't update matrix: %w", err)
	}
	return m, nil
}

func (s *service) DeleteMatrix(_ context.Context, id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("service can't delete matrix: %w", err)
	}
	return nil
}
