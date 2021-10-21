local: env
	docker-compose -f docker-compose.dev.yml up --detach --build app

env:
	docker-compose -f docker-compose.dev.yml up --detach db dba

clean:
	docker-compose -f docker-compose.dev.yml down

.PHONY: local env clean
