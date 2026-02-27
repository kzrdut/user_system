package mysql

import (
	"context"
	"fmt"
	"user_system/internal/domain"

	"gorm.io/gorm"
)

// MySQLRepository Repository 层实现
type MySQLRepository struct {
	db *gorm.DB
}

// NewMySQLRepository 初始化用户数据仓库
func NewMySQLRepository(db *gorm.DB) *MySQLRepository {
	return &MySQLRepository{db: db}
}

// CreateUser 新增用户
func (r *MySQLRepository) CreateUser(ctx context.Context, user *domain.User) (domain.UserID, error) {
	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

// GetUserByID 根据 id 获取用户
func (r *MySQLRepository) GetUserByID(ctx context.Context, userID domain.UserID) (*domain.User, error) {
	user := domain.User{}
	if err := r.db.WithContext(ctx).First(&user, userID).Error; err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}
	return &user, nil
}

// GetUserByUsername 根据用户名获取用户
func (r *MySQLRepository) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	user := domain.User{}
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, fmt.Errorf("根据用户名查找用户失败: %w", err)
	}
	return &user, nil
}

// DeleteUser 删除用户
func (r *MySQLRepository) DeleteUser(ctx context.Context, userID domain.UserID) error {
	if err := r.db.WithContext(ctx).Delete(&domain.User{}, userID).Error; err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
