module example

go 1.13

replace github.com/2637309949/dolphin => ../../

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43

require (
	github.com/2637309949/dolphin v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.4.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/viper v1.4.0
	github.com/xormplus/xorm v0.0.0-20190926190243-42377f593eb1
	gopkg.in/guregu/null.v3 v3.4.0
)
