.PHONY: test-integration

test-integration:
	docker compose -f ./test/docker-compose-test.yml up
	go test ./test/...