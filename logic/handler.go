package logic

import (
	"github.com/BranDebs/challenge-bot/repository"
)

type Handler interface {
	ChallengeHandler
	GoalHandler
}

type handler struct {
	challengeHandler
	goalHandler
}

func New(repo repository.Repository) Handler {
	return &handler{
		challengeHandler: challengeHandler{
			repo: repo,
		},
		goalHandler: goalHandler{
			repo: repo,
		},
	}
}
