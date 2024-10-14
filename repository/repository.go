package repository

import (
	"context"

	"github.com/mariasilva795/go-api-rest/models"
)

type Repository interface {
	InsertUser(ctx context.Context, user *models.User) error
	Close() error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	InsertEmotionalDailyLog(ctx context.Context, emotion *models.EmotionalDailyLog) error
	GetEmotionalDailyLogById(ctx context.Context, id string) (*models.EmotionalDailyLog, error)
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implementation.GetUserByEmail(ctx, email)
}

func Close() error {
	return implementation.Close()
}

// Bank Emotional

func InsertEmotionalDailyLog(ctx context.Context, emotionalDailyLog *models.EmotionalDailyLog) error {
	return implementation.InsertEmotionalDailyLog(ctx, emotionalDailyLog)
}

func GetEmotionalDailyLogById(ctx context.Context, id string) (*models.EmotionalDailyLog, error) {
	return implementation.GetEmotionalDailyLogById(ctx, id)
}
