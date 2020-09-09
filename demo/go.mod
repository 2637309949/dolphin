module demo

go 1.13

replace github.com/2637309949/dolphin => ../

require (
	github.com/2637309949/dolphin v0.0.0-20200909073611-9a94feac5fa6
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.3.3
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553
	google.golang.org/grpc v1.26.0
)
