.PHONY: test security run stop compose-up compose-down

BUILD_DIR = $(PWD)/app

security:
	gosec -quiet ./...

test: security
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

docker_build_image:
	docker build -t api-portofolio:latest .

docker_app: docker_build_image
	docker run -d \
        		--name api_portofolio_container \
        		-p 3008:3008 \
        		-e SERVER_NAME="Portofolio Service API" \
        		-e SERVER_URL="0.0.0.0:3008" \
        		-e SERVER_READ_TIMEOUT="60" \
        		-e JWT_SECRET_KEY="XL67xKPmECeEkMdN" \
        		-e JWT_EXPIRE_MINUTES="240" \
        		-e DB_SERVER_URL="host=172.24.5.241 port=5432 user=pusdatin password=Kemendagri.1 dbname=p3dn sslmode=disable" \
        		-e LOGIN_URL="https://sipd.kemendagri.go.id/api/auth-user-sipd" \
        		-e LOGIN_TOKEN="58a176016a480b21a5b3b720f90f9123" \
        		-e SIRUP_URL="https://sipd.kemendagri.go.id/api/vSIRUP" \
        		-e SIRUP_TOKEN="58a176016a480b21a5b3b720f90f9123" \
        		-e WILAYAH_URL="https://sipd.kemendagri.go.id/api/vSIPD-Wilayah" \
        		-e WILAYAH_TOKEN="70a739816e9df4d009aadc15bb00094d" \
        		api-portofolio:latest

run: docker_app
stop:
	docker container stop api_portofolio_container
	docker container rm api_portofolio_container
	docker rmi api-portofolio:latest

compose-up:
	docker-compose up -d
compose-down:
	docker-compose down
