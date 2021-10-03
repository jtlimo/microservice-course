package domain

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
)

type ServiceInterface interface {
	ListMatrix(context.Context, map[string]interface{}) ([]Matrix, error)
	CreateMatrix(context.Context, *Matrix) (Matrix, error)
	FindMatrix(context.Context, string) (Matrix, error)
	UpdateMatrix(context.Context, *Matrix) (Matrix, error)
	DeleteMatrix(context.Context, string) error
}

type Service struct {
	repo   Repository
	logger log.Logger
}

func NewService(repo Repository, logger log.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}

func (s *Service) ListMatrix(_ context.Context, filters map[string]interface{}) ([]Matrix, error) {
	ms, err := s.repo.List(filters)
	if err != nil {
		return []Matrix{}, fmt.Errorf("Service didn't found any matrix: %w", err)
	}
	return ms, nil
}

func (s *Service) CreateMatrix(_ context.Context, matrix *Matrix) (Matrix, error) {
	m, err := s.repo.Create(matrix)
	if err != nil {
		return Matrix{}, fmt.Errorf("Service can't create matrix: %w", err)
	}
	return m, nil
}

func (s *Service) FindMatrix(_ context.Context, id string) (Matrix, error) {
	m, err := s.repo.Find(id)
	if err != nil {
		return Matrix{}, fmt.Errorf("Service can't find matrix: %w", err)
	}
	return m, nil
}

func (s *Service) UpdateMatrix(_ context.Context, matrix *Matrix) (Matrix, error) {
	m, err := s.repo.Update(matrix)
	if err != nil {
		return Matrix{}, fmt.Errorf("Service can't update matrix: %w", err)
	}
	return m, nil
}

func (s *Service) DeleteMatrix(_ context.Context, id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("Service can't delete matrix: %w", err)
	}
	return nil
}