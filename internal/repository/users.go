package repository

import (
	"context"

	"github.com/wannn28/TASK-MIKTI/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
	GetByResetPasswordToken(ctx context.Context, token string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	GetAll(ctx context.Context) ([]entity.User, error)
	GetByID(ctx context.Context, id int64) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, user *entity.User) error
	ResetPassword(ctx context.Context, user *entity.User) error
	GetByVerifyEmailToken(ctx context.Context, token string) (*entity.User, error)
	// VerifyEmail(ctx context.Context, token string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	result := new(entity.User)
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Create(&user).Error
}

func (r *userRepository) GetAll(ctx context.Context) ([]entity.User, error) {
	result := make([]entity.User, 0)
	if err := r.db.WithContext(ctx).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *userRepository) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	result := new(entity.User)
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *userRepository) Update(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Updates(&user).Error
}

func (r *userRepository) Delete(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Delete(&user).Error
}

func (r *userRepository) ResetPassword(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Updates(&user).Error
}

func (r *userRepository) GetByResetPasswordToken(ctx context.Context, token string) (*entity.User, error) {
	result := new(entity.User)
	if err := r.db.WithContext(ctx).Where("reset_password_token = ?", token).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// func (r *userRepository) GetByVerifyEmailToken(ctx context.Context, token string) (*entity.User, error) {
// 	result := new(entity.User)
// 	if err := r.db.WithContext(ctx).Where("verify_email_token = ?", token).First(&result).Error; err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

func (r *userRepository) GetByVerifyEmailToken(ctx context.Context, token string) (*entity.User, error) {
	result := new(entity.User)
	if err := r.db.WithContext(ctx).Where("verify_email_token = ?", token).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// func (r *userRepository) VerifyEmail(ctx context.Context, token string) error {
// 	return r.db.WithContext(ctx).Where("verify_email_token = ?", token).Updates(map[string]interface{}{"is_email_verified": true}).Error
// }
