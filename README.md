# challenge-bot
Telegram bot to handle challenges

## Quick Start
To run the bot in local env:
```bash
# Runs locally challengebot along with the resources.
make local
```

To run the env only:
```bash
# Creates the resources needed like DB, DBA.
make env
```

To clean:
```bash
make clean
```

## Using DBA
```bash
# Go to localhost:1337

System: PostgreSQL
Server: challenge_db
Username: postgres
Password: postgres123
Database: challenge

```

## Schema
```golang
type Challenge struct {
    ID uint64
    Name string
    UserIDs []uint64
    StartDate uint64
    EndDate uint64
    Schema []byte // example: "{value: int}", "{age: int, money: int64}"
}

type User struct {
    ID uint64 // retrieve from Telegram
    Username string // retrieve from Telegram

}

type Goal struct {
    ID uint64
    UserID uint64
    ChallengeID uint64
    Value []byte // {value: 2}, "{age: 30, money: 50}"
}


type Progress struct {    //cannot be updated
    ID uint64
    UserID uint64
    ChallengeID uint64
    Value []byte // {value: 2}, "{age: 30, money: 50}"
    Date uint64
}
```

