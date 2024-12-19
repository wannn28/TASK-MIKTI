package repository

import (
	"context"
	"strings"

	"github.com/wannn28/TASK-MIKTI/internal/entity"
	"github.com/wannn28/TASK-MIKTI/internal/http/dto"
	"gorm.io/gorm"
)

type MovieRepository interface {
	GetAll(ctx context.Context, req dto.GetAllMovieRequest) ([]entity.Movie, error)
	GetByID(ctx context.Context, id int64) (*entity.Movie, error)
	Create(ctx context.Context, movie *entity.Movie) error
	Update(ctx context.Context, movie *entity.Movie) error
	Delete(ctx context.Context, movie *entity.Movie) error
}

type movieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return &movieRepository{db}
}

func (r *movieRepository) GetAll(ctx context.Context, req dto.GetAllMovieRequest) ([]entity.Movie, error) {
	result := make([]entity.Movie, 0)
	query := r.db.WithContext(ctx)
	if req.Search != "" {
		search := strings.ToLower(req.Search)
		query = query.Where("LOWER(title) LIKE ?", "%"+search+"%").
			Or("LOWER(director) LIKE ?", "%"+search+"%").
			Or("LOWER(year) = ?", search)
	}

	if req.Sort != "" && req.Order != "" {
		query = query.Order(req.Sort + " " + req.Order)
	}
	if req.Page != 0 && req.Limit != 0 {
		query = query.Offset((req.Page - 1) * req.Limit).Limit(req.Limit)
	}
	if err := query.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *movieRepository) GetByID(ctx context.Context, id int64) (*entity.Movie, error) {
	result := new(entity.Movie)
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *movieRepository) Create(ctx context.Context, movie *entity.Movie) error {
	return r.db.WithContext(ctx).Create(&movie).Error
}

func (r *movieRepository) Update(ctx context.Context, movie *entity.Movie) error {
	return r.db.WithContext(ctx).Updates(movie).Error
}

func (r *movieRepository) Delete(ctx context.Context, movie *entity.Movie) error {
	return r.db.WithContext(ctx).Delete(movie).Error
}
