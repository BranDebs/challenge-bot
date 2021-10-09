env:
	docker-compose -f docker-compose.dev.yml up -d

clean:
	docker-compose -f docker-compose.dev.yml down

.PHONY: env clean
