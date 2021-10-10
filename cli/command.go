package cli

import "strings"

type command uint

const (
	unknown = command(iota)
	exit
	challenge
	user
	goal
	progress
)

func (c command) String() string {
	switch c {
	case unknown:
		return "unknown"
	case exit:
		return "exit"
	case challenge:
		return "challenge"
	case user:
		return "user"
	case goal:
		return "goal"
	case progress:
		return "progress"
	}
	return ""
}

func strToCommand(cmdStr string) command {
	switch lower := strings.ToLower(cmdStr); lower {
	case "exit":
		return exit
	case "challenge":
		return challenge
	case "user":
		return user
	case "goal":
		return goal
	case "progress":
		return progress
	}
	return unknown
}
