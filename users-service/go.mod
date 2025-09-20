module github.com/valdevay/users-service

go 1.25.0

require (
	github.com/joho/godotenv v1.5.1
	github.com/valdevay/project-protos/proto/users v0.0.0-20250918103921-0481ab8d9c16
	google.golang.org/grpc v1.75.1
	gorm.io/driver/postgres v1.5.7
	gorm.io/gorm v1.25.7
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.4.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/crypto v0.39.0 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250707201910-8d1bb00bc6a7 // indirect
	google.golang.org/protobuf v1.36.9 // indirect
)

replace github.com/valdevay/project-protos => ../proto
