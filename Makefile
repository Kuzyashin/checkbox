docker_run:
	docker-compose -f build/docker-compose.yaml up
docker_run_detach:
	docker-compose -f build/docker-compose.yaml up -d
run_local:
	go run cmd/main/main.go