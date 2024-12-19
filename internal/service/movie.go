package service

import (
	"context"

	"github.com/wannn28/TASK-MIKTI/internal/entity"
	"github.com/wannn28/TASK-MIKTI/internal/http/dto"
	"github.com/wannn28/TASK-MIKTI/internal/repository"
)

type MovieService interface {
	GetAll(ctx context.Context, req dto.GetAllMovieRequest) ([]entity.Movie, error)
	GetByID(ctx context.Context, id int64) (*entity.Movie, error)
	Create(ctx context.Context, req dto.CreateMovieRequest) error
	Update(ctx context.Context, req dto.UpdateMovieRequest) error
	Delete(ctx context.Context, movie *entity.Movie) error
}

type movieService struct {
	movieRepository repository.MovieRepository
}

func NewMovieService(movieRepository repository.MovieRepository) MovieService {
	return &movieService{movieRepository}
}

func (s *movieService) GetAll(ctx context.Context, req dto.GetAllMovieRequest) ([]entity.Movie, error) {
	return s.movieRepository.GetAll(ctx, req)
}

func (s *movieService) GetByID(ctx context.Context, id int64) (*entity.Movie, error) {
	return s.movieRepository.GetByID(ctx, id)
}

func (s *movieService) Create(ctx context.Context, req dto.CreateMovieRequest) error {
	movie := &entity.Movie{
		Title:       req.Title,
		Year:        req.Year,
		Director:    req.Director,
		Description: req.Description,
	}
	return s.movieRepository.Create(ctx, movie)
}

func (s *movieService) Update(ctx context.Context, req dto.UpdateMovieRequest) error {
	movie, err := s.movieRepository.GetByID(ctx, req.ID)
	if err != nil {
		return err
	}
	if req.Title != "" {
		movie.Title = req.Title
	}
	if req.Year != 0 {
		movie.Year = req.Year
	}
	if req.Director != "" {
		movie.Director = req.Director
	}
	if req.Description != "" {
		movie.Description = req.Description
	}
	return s.movieRepository.Update(ctx, movie)
}

func (s *movieService) Delete(ctx context.Context, movie *entity.Movie) error {
	return s.movieRepository.Delete(ctx, movie)
}
