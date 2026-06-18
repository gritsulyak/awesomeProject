.PHONY: test test-integration

test:
	go test ./... -coverprofile=coverage.out -covermode=atomic

test-integration:
	docker compose -f test/docker-compose-test.yml up -d
	go test --tags=test_integration ./test/... -v -timeout 60s; \
	docker compose -f test/docker-compose-test.yml down