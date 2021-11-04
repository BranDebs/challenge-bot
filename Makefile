build:
	docker build -t challenge-bot .

dev: env
	docker-compose -f docker-compose.dev.yml up --build app

env:
	docker-compose -f docker-compose.dev.yml up --detach db dba

clean:
	docker-compose -f docker-compose.dev.yml down

.PHONY: build dev env clean
