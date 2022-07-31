.PHONY: test security run stop compose-up compose-down

BUILD_DIR = $(PWD)/app

security:
	gosec -quiet ./...

test: security
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

docker_build_image:
	docker build -t api_portofolio:latest .

docker_app: docker_build_image
	docker run -d \
        		--name api_portofolio_container \
        		-p 3008:3008 \
        		-e SERVER_NAME="Portofolio Service API" \
        		-e SERVER_URL="0.0.0.0:3008" \
        		-e SERVER_READ_TIMEOUT="60" \
        		-e JWT_SECRET_KEY="XL67xKPmECeEkMdN" \
        		-e JWT_EXPIRE_MINUTES="240" \
        		-e DB_SERVER_URL="host=localhost port=5432 user=postgres password=5g6B7UcTBJVDkNvL dbname=portofolio sslmode=disable" \
				api_portofolio:latest

run: docker_app
stop:
	docker container stop api_portofolio_container
	docker container rm api_portofolio_container
	docker rmi api_portofolio:latest

compose-up:
	docker-compose up -d
compose-down:
	docker-compose down
