package cli

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/BranDebs/challenge-bot/logic"
	"github.com/BranDebs/challenge-bot/model"
)

var (
	errExit = errors.New("exit signal")
)

var (
	cmdList = []string{exit.String(), challenge.String(), user.String(), goal.String(), progress.String()}
)

type CLI struct {
	cHandler logic.ChallengeHandler
}

func New(handler logic.Handler) *CLI {
	return &CLI{
		cHandler: handler,
	}
}

func (cli CLI) Listen() error {
	log.Println("CLI is listening.")
	var cmd string
	var subCmd string
	fmt.Println("Enter command: ")
	fmt.Scanln(&cmd, &subCmd)
	if err := cli.Process(strToCommand(cmd), subCmd); err != nil {
		return err
	}

	return nil
}

func (cli CLI) Process(cmd command, subCmd string) error {
	switch cmd {
	case challenge:
		return cli.challengeCommand(context.Background(), subCmd)
	case exit:
		return errExit
	}

	return nil
}

func (cli CLI) challengeCommand(ctx context.Context, subCmd string) error {
	switch subCmd {
	case "create":
		return cli.createChallenge(ctx)
	case "list":
		return cli.listChallenges(ctx)
	case "join":
		return cli.joinChallenge(ctx)
	}

	return nil
}

func (cli CLI) createChallenge(ctx context.Context) error {
	var (
		name   = "test challenge"
		userID = uint64(123)
		schema = `{"weight": "int64"}`
	)

	currentDate := time.Now()
	endDate := currentDate.Add(24 * time.Hour)

	challenge := &model.Challenge{
		Name:      name,
		UserIDs:   []uint64{uint64(userID)},
		StartDate: uint64(currentDate.Unix()),
		EndDate:   uint64(endDate.Unix()),
		Schema:    []byte(schema),
	}

	log.Printf("Challenge created: %v", challenge)

	return cli.cHandler.CreateChallenge(ctx, challenge)
}

func (cli CLI) listChallenges(ctx context.Context) error {
	challenges, err := cli.cHandler.ListChallenges(ctx)
	if err != nil {
		return err
	}

	log.Println("Listing challenges:")
	for _, c := range challenges {
		log.Printf("%+v\n", *c)
	}

	return nil
}

func (cli CLI) joinChallenge(ctx context.Context) error {
	if err := cli.cHandler.JoinChallenge(ctx, 5, 987); err != nil {
		return err
	}

	log.Println("Joined challenge.")

	return nil

}
