PROTOS  := proto/tasks/*.proto proto/users/*.proto
OUT_DIR := .

generate:
	./bin/protoc.exe \
	  --go_out=$(OUT_DIR) --go_opt=paths=source_relative \
	  --go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
	  --plugin=protoc-gen-go=C:\Users\vladimir.kireev\go\bin\protoc-gen-go.exe \
	  --plugin=protoc-gen-go-grpc=C:\Users\vladimir.kireev\go\bin\protoc-gen-go-grpc.exe \
	  --proto_path=proto \
	  $(PROTOS)

clean:
	find . -name "*.pb.go" -delete

# Docker команды
docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f postgres

docker-clean:
	docker-compose down -v

# Разработка
dev-setup: docker-up
	@echo "Waiting for database to be ready..."
	@timeout /t 10 /nobreak > nul
	@echo "Database is ready!"

dev-run:
	cd users-service && go run cmd/server/main.go

dev-build:
	cd users-service && go build -o ../bin/users-service cmd/server/main.go

# Полная перезагрузка окружения
reset: docker-clean docker-up dev-setup