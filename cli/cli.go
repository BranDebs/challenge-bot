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
	errExit       = errors.New("exit signal")
	errInvalidCmd = errors.New("invalid command")
)

var (
	cmdList = []string{exit.String(), challenge.String(), user.String(), goal.String(), progress.String()}
)

type CLI struct {
	cHandler logic.ChallengeHandler
	gHandler logic.GoalHandler
	pHandler logic.ProgressHandler
}

func New(handler logic.Handler) *CLI {
	return &CLI{
		cHandler: handler,
		gHandler: handler,
		pHandler: handler,
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
	case goal:
		return cli.goalCommand(context.Background(), subCmd)
	case progress:
		return cli.progressCommand(context.Background(), subCmd)
	case exit:
		return errExit
	}

	return errInvalidCmd
}

func (cli CLI) challengeCommand(ctx context.Context, subCmd string) error {
	switch subCmd {
	case "create":
		return cli.createChallenge(ctx)
	case "list":
		return cli.listChallenges(ctx)
	case "find":
		return cli.findChallenge(ctx)
	case "join":
		return cli.joinChallenge(ctx)
	}

	return errInvalidCmd
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

func (cli CLI) findChallenge(ctx context.Context) error {
	challenge, err := cli.cHandler.FindChallenge(ctx, 5)
	if err != nil {
		return err
	}

	log.Printf("Challenge: %+v", *challenge)

	return nil
}

func (cli CLI) joinChallenge(ctx context.Context) error {
	if err := cli.cHandler.JoinChallenge(ctx, 5, 987); err != nil {
		return err
	}

	log.Println("Joined challenge.")

	return nil
}

func (cli CLI) goalCommand(ctx context.Context, subCmd string) error {
	switch subCmd {
	case "create":
		return cli.createGoal(ctx)
	case "find":
		return cli.findGoal(ctx)
	}
	return nil
}

func (cli CLI) createGoal(ctx context.Context) error {
	goal := &model.Goal{
		UserID:      123,
		ChallengeID: 5,
		Value:       []byte(`{"weight": 10}`),
	}

	if err := cli.gHandler.CreateGoal(ctx, goal); err != nil {
		return err
	}

	return nil
}

func (cli CLI) findGoal(ctx context.Context) error {
	goal, err := cli.gHandler.FindGoal(ctx, 5, 123)
	if err != nil {
		return err
	}

	log.Printf("Find goal: %+v\n", *goal)
	return nil
}

func (cli CLI) progressCommand(ctx context.Context, subCmd string) error {
	switch subCmd {
	case "create":
		return cli.createProgress(ctx)
	case "list":
		return cli.listProgress(ctx)
	}
	return nil
}

func (cli CLI) createProgress(ctx context.Context) error {
	progress := &model.Progress{
		UserID:      123,
		ChallengeID: 5,
		Value:       []byte(`{"weight": 11}`),
		Date:        uint64(time.Now().Unix()),
	}

	completed, err := cli.pHandler.CreateProgress(ctx, progress)
	if err != nil {
		return err
	}

	fmt.Printf("Create progress: %+v\n", *progress)

	if completed {
		fmt.Println("Completed challenge!")
	}

	return nil
}

func (cli CLI) listProgress(ctx context.Context) error {
	progress, err := cli.pHandler.ListProgress(ctx, 5, 123)
	if err != nil {
		return err
	}

	fmt.Println("Listing progress")
	for _, p := range progress {
		fmt.Printf("progress: %+v\n", *p)
	}

	return nil
}
