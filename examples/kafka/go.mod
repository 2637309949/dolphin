module kafka

go 1.13

replace github.com/2637309949/dolphin => ../../

require (
	github.com/2637309949/dolphin v1.0.53
	github.com/go-errors/errors v1.0.1
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.3.3
	github.com/klauspost/compress v1.11.7 // indirect
	github.com/pierrec/lz4 v2.6.0+incompatible // indirect
	github.com/pkg/errors v0.8.1
	github.com/segmentio/kafka-go v0.4.9
	golang.org/x/net v0.0.0-20201224014010-6772e930b67b
	google.golang.org/grpc v1.26.0
)
