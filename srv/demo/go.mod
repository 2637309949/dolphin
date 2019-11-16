module example

go 1.13

replace github.com/2637309949/dolphin => ../../

// replace github.com/go-xorm/core v0.6.3 => xorm.io/core v0.6.3

require (
	github.com/2637309949/dolphin v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.4.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/spf13/viper v1.4.0
	github.com/thoas/go-funk v0.4.0
	gopkg.in/guregu/null.v3 v3.4.0
)
