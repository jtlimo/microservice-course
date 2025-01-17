package domain

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) Course(_ context.Context, id uuid.UUID) (Course, error) {
	c, err := s.courses.Course(id)
	if err != nil {
		return Course{}, fmt.Errorf("service can't find course: %w", err)
	}
	return c, nil
}

func (s *Service) Courses(_ context.Context) ([]Course, error) {
	cc, err := s.courses.Courses()
	if err != nil {
		return []Course{}, fmt.Errorf("service didn't found any course: %w", err)
	}
	return cc, nil
}

func (s *Service) CreateCourse(_ context.Context, c *Course) error {
	if err := s.courses.CreateCourse(c); err != nil {
		return fmt.Errorf("service can't create course: %w", err)
	}
	return nil
}

func (s *Service) UpdateCourse(_ context.Context, c *Course) error {
	if err := s.courses.UpdateCourse(c); err != nil {
		return fmt.Errorf("service can't update course: %w", err)
	}
	return nil
}

func (s *Service) DeleteCourse(_ context.Context, id uuid.UUID) error {
	if err := s.courses.DeleteCourse(id); err != nil {
		return fmt.Errorf("service can't delete course: %w", err)
	}
	return nil
}
