package postgres

import (
	"context"
	"log"

	"github.com/BranDebs/challenge-bot/model"
	"github.com/BranDebs/challenge-bot/repository"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Client struct {
	db *gorm.DB
}

func New(settings *Settings) repository.Repository {
	db, err := gorm.Open(postgres.Open(settings.DSN()), &gorm.Config{})
	if err != nil {
		return nil
	}

	db.AutoMigrate(&challengeEntity{})

	return &Client{
		db: db,
	}
}

func (c *Client) CreateChallenge(ctx context.Context, challenge *model.Challenge) error {
	var e challengeEntity
	e.fromModel(challenge)

	log.Printf("Challenge entity: %v", e)

	result := c.db.Create(&e)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c *Client) FindChallenge(ctx context.Context, id uint64) (*model.Challenge, error) {
	return nil, nil
}

func (c *Client) ListChallenges(ctx context.Context, filters repository.Filters, offset, limit uint) ([]*model.Challenge, error) {
	return nil, nil
}

func (c *Client) CreateUser(ctx context.Context, user *model.User) error {
	return nil
}

func (c *Client) FindUser(ctx context.Context, id uint64) (*model.User, error) {
	return nil, nil
}

func (c *Client) ListUsers(ctx context.Context, filter repository.Filters, offset, limit uint) ([]*model.User, error) {
	return nil, nil
}

func (c *Client) CreateGoal(ctx context.Context, goal *model.Goal) error {
	return nil
}

func (c *Client) FindGoal(ctx context.Context, id uint64) (*model.Goal, error) {
	return nil, nil
}

func (c *Client) ListGoals(ctx context.Context, filter repository.Filters, offset, limit uint) ([]*model.Goal, error) {
	return nil, nil
}

func (c *Client) CreateProgress(ctx context.Context, progress *model.Progress) error {
	return nil
}
func (c *Client) FindProgress(ctx context.Context, id uint64) (*model.Progress, error) {
	return nil, nil
}

func (c *Client) ListProgress(ctx context.Context, filter repository.Filters, offset, limit uint) ([]*model.Progress, error) {
	return nil, nil
}
