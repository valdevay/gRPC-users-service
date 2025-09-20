module github.com/valdevay/users-service

go 1.21

require (
	github.com/lib/pq v1.10.9
	github.com/valdevay/project-protos/proto/user v0.0.0-20250919073248-a72ec6340465
	google.golang.org/grpc v1.64.0
)

require (
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)

replace github.com/valdevay/project-protos => ../proto
