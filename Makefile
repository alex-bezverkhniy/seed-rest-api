.PHONY: start
start:
	docker-compose up -d --build

.PHONY: stop
stop:
	docker-compose rm -v --force --stop
	docker image rm seed-rest-api:latest

.PHONY: test
test:
	sh ./scripts/e2e-testing.sh