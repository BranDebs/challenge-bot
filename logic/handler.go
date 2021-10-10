package logic

import (
	"github.com/BranDebs/challenge-bot/repository"
)

type Handler interface {
	ChallengeHandler
}

type handler struct {
	challengeHandler
}

func New(repo repository.Repository) Handler {
	return &handler{
		challengeHandler: challengeHandler{
			repo: repo,
		},
	}
}
