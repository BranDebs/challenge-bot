package postgres

import (
	"context"

	"github.com/BranDebs/challenge-bot/challenge"
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

	return &Client{
		db: db,
	}
}

func (c *Client) CreateChallenge(ctx context.Context, challenge *challenge.Challenge) error {
	return nil
}

func (c *Client) FindChallenge(ctx context.Context, id uint64) (*challenge.Challenge, error) {
	return nil, nil
}

func (c *Client) ListChallenges(ctx context.Context, filters repository.Filters, offset, limit uint64) ([]*challenge.Challenge, error) {
	return nil, nil
}

func (c *Client) CreateUser(ctx context.Context, user *challenge.User) error {
	return nil
}

func (c *Client) FindUser(ctx context.Context, id uint64) (*challenge.User, error) {
	return nil, nil
}

func (c *Client) ListUsers(ctx context.Context, filter repository.Filters, offset, limit uint64) ([]*challenge.User, error) {
	return nil, nil
}

func (c *Client) CreateGoal(ctx context.Context, goal *challenge.Goal) error {
	return nil
}

func (c *Client) FindGoal(ctx context.Context, id uint64) (*challenge.Goal, error) {
	return nil, nil
}

func (c *Client) ListGoals(ctx context.Context, filter repository.Filters, offset, limit uint64) ([]*challenge.Goal, error) {
	return nil, nil
}

func (c *Client) CreateProgress(ctx context.Context, progress *challenge.Progress) error {
	return nil
}
func (c *Client) FindProgress(ctx context.Context, id uint64) (*challenge.Progress, error) {
	return nil, nil
}

func (c *Client) ListProgress(ctx context.Context, filter repository.Filters, offset, limit uint64) ([]*challenge.Progress, error) {
	return nil, nil
}
