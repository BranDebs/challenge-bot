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

	db.AutoMigrate(&challengeEntity{}, &goalEntity{}, &progressEntity{})

	return &Client{
		db: db,
	}
}

func (c *Client) CreateChallenge(ctx context.Context, challenge *model.Challenge) error {
	var e challengeEntity
	e.fromModel(challenge)

	log.Printf("Challenge entity: %+v", e)

	res := c.db.Create(&e)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (c *Client) FindChallenge(ctx context.Context, id uint64) (*model.Challenge, error) {
	var e challengeEntity

	res := c.db.First(&e, id)
	if res.Error != nil {
		return nil, res.Error
	}

	return e.toModel(), nil
}

func (c *Client) ListChallenges(ctx context.Context, filters repository.Filters, offset, limit uint) ([]*model.Challenge, error) {
	var entities []*challengeEntity

	var res *gorm.DB
	if len(filters) > 0 {
		res = c.db.Where(filters).Find(&entities).Offset(int(offset)).Limit(int(limit))
	} else {
		res = c.db.Find(&entities).Offset(int(offset)).Limit(int(limit))
	}

	if res.Error != nil {
		return nil, res.Error
	}

	challenges := make([]*model.Challenge, len(entities))

	for i, e := range entities {
		challenges[i] = e.toModel()
	}

	return challenges, nil
}

func (c *Client) UpdateChallenge(ctx context.Context, challenge *model.Challenge) error {
	var e challengeEntity
	e.fromModel(challenge)

	res := c.db.Model(&e).Updates(e)

	if res.Error != nil {
		return res.Error
	}

	return nil
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
	var e goalEntity
	e.fromModel(goal)

	log.Printf("Goal entity: %+v", e)

	res := c.db.Create(&e)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (c *Client) FindGoal(ctx context.Context, id uint64) (*model.Goal, error) {
	var e goalEntity

	res := c.db.First(&e, id)
	if res.Error != nil {
		return nil, res.Error
	}

	return e.toModel(), nil
}

func (c *Client) ListGoals(ctx context.Context, filters repository.Filters, offset, limit uint) ([]*model.Goal, error) {
	var entities []*goalEntity

	var res *gorm.DB
	if len(filters) > 0 {
		res = c.db.Where(map[string]interface{}(filters)).Find(&entities).Offset(int(offset)).Limit(int(limit))
	} else {
		res = c.db.Find(&entities).Offset(int(offset)).Limit(int(limit))
	}

	if res.Error != nil {
		return nil, res.Error
	}

	goals := make([]*model.Goal, len(entities))

	for i, e := range entities {
		goals[i] = e.toModel()
	}

	return goals, nil
}

func (c *Client) CreateProgress(ctx context.Context, progress *model.Progress) error {
	var e progressEntity

	e.fromModel(progress)

	log.Printf("Progress entity: %+v\n", e)

	res := c.db.Create(&e)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
func (c *Client) FindProgress(ctx context.Context, id uint64) (*model.Progress, error) {
	var e progressEntity

	res := c.db.First(&e, id)
	if res.Error != nil {
		return nil, res.Error
	}

	return e.toModel(), nil
}

func (c *Client) ListProgress(ctx context.Context, filters repository.Filters, offset, limit uint) ([]*model.Progress, error) {
	var entities []*progressEntity

	var res *gorm.DB
	if len(filters) > 0 {
		repoFilters := map[string]interface{}(filters)
		res = c.db.Where(repoFilters).Find(&entities).Offset(int(offset)).Limit(int(limit))
	} else {
		res = c.db.Find(&entities).Offset(int(offset)).Limit(int(limit))
	}

	if res.Error != nil {
		return nil, res.Error
	}

	progress := make([]*model.Progress, len(entities))

	for i, e := range entities {
		progress[i] = e.toModel()
	}

	return progress, nil
}
