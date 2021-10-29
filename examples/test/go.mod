module test

go 1.13

// fix io/fs
replace (
	github.com/spf13/afero => github.com/spf13/afero v1.5.1
	golang.org/x/tools => github.com/golang/tools v0.1.0
)

require (
	github.com/2637309949/dolphin v1.2.15
	github.com/gin-contrib/cache v1.1.0
	github.com/mattn/go-sqlite3 v1.14.9
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.9.0
	github.com/thoas/go-funk v0.9.1
)
