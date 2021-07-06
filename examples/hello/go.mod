module hello

go 1.13

replace github.com/2637309949/dolphin => ../../

require (
	github.com/2637309949/dolphin v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.7.2
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/mattn/go-sqlite3 v1.14.7
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.8.1
	github.com/thoas/go-funk v0.8.0
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e
	google.golang.org/grpc v1.39.0
)
