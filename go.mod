module zmin

go 1.14

replace github.com/ouczbs/zmin => ../zmin

require (
	github.com/go-redis/redis/v8 v8.0.0-beta.6
	github.com/golang/protobuf v1.4.2
	github.com/ouczbs/zmin v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.8.1
	go.mongodb.org/mongo-driver v1.3.5
	go.uber.org/zap v1.15.0
	google.golang.org/protobuf v1.25.0
)
