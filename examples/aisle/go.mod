module aisle

go 1.13

require (
	github.com/2637309949/dolphin v1.0.67
	github.com/gin-contrib/cache v1.1.0
	github.com/mattn/go-sqlite3 v1.14.7
	github.com/shopspring/decimal v1.2.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.8.1
	github.com/thoas/go-funk v0.8.0
)

replace github.com/2637309949/dolphin => ../../

// fix io/fs
replace (
	github.com/spf13/afero => github.com/spf13/afero v1.5.1
	golang.org/x/tools => github.com/golang/tools v0.1.0
)
