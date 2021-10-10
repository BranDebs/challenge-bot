package logic

import (
	"github.com/BranDebs/challenge-bot/repository"
)

type Handler interface {
	ChallengeHandler
	GoalHandler
	ProgressHandler
}

type handler struct {
	challengeHandler
	goalHandler
	progressHandler
}

func New(repo repository.Repository) Handler {
	return &handler{
		challengeHandler: challengeHandler{
			repo: repo,
		},
		goalHandler: goalHandler{
			repo: repo,
		},
		progressHandler: progressHandler{
			cRepo: repo,
			pRepo: repo,
			gRepo: repo,
		},
	}
}
