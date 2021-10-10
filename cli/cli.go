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
	for {
		var cmd string
		fmt.Println("Enter command: ")
		fmt.Scanln(&cmd)
		if err := cli.Process(strToCommand(cmd)); err != nil {
			log.Fatalf("Processing err: %v", err)
			break
		}
	}

	return nil
}

func (cli CLI) Process(cmd command) error {
	switch cmd {
	case challenge:
		return cli.CreateChallenge(context.Background())
	case exit:
		return errExit
	}

	return nil
}

func (cli CLI) CreateChallenge(ctx context.Context) error {
	var (
		name   = "test challenge"
		userID = uint64(123)
		schema = `{"weight": "int64"}`
	)

	// fmt.Scanln()
	// fmt.Println("Enter challenge name:")
	// fmt.Scanln(&name)

	// fmt.Println("Enter user ID:")
	// fmt.Scanln(&userID)

	// fmt.Println("Enter schema:")
	// fmt.Scanln(&schema)

	// fmt.Println("Enter end date:")
	// fmt.Scanln(&endDateStr)
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
