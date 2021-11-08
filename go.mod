module github.com/2637309949/dolphin

go 1.16

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/Chronokeeper/anyxml v0.0.0-20160530174208-54457d8e98c6
	github.com/CloudyKit/jet v2.1.2+incompatible
	github.com/KyleBanks/depth v1.2.1
	github.com/agrison/go-tablib v0.0.0-20160310143025-4930582c22ee
	github.com/boombuler/barcode v1.0.1
	github.com/bradfitz/gomemcache v0.0.0-20180710155616-bc664df96737
	github.com/eriklott/mustache v0.10.1
	github.com/fsnotify/fsnotify v1.4.9
	github.com/ghodss/yaml v1.0.0
	github.com/gin-contrib/cache v1.1.0
	github.com/gin-gonic/gin v1.7.1
	github.com/go-errors/errors v1.0.1
	github.com/go-openapi/spec v0.20.3
	github.com/go-redis/redis/v8 v8.8.0
	github.com/go-session/session v3.1.2+incompatible
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/google/uuid v1.2.0
	github.com/gorilla/securecookie v1.1.1
	github.com/json-iterator/go v1.1.11
	github.com/labstack/echo/v4 v4.6.1
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lib/pq v1.10.2
	github.com/mattn/go-sqlite3 v1.14.7
	github.com/memcachier/mc/v3 v3.0.3
	github.com/pkg/errors v0.9.1
	github.com/robfig/cron/v3 v3.0.1
	github.com/robfig/go-cache v0.0.0-20130306151617-9fc39e0dbf62
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749
	github.com/shurcooL/vfsgen v0.0.0-20200824052919-0d455de96546
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.8.1
	github.com/syndtr/goleveldb v1.0.0
	github.com/thoas/go-funk v0.8.0
	github.com/tidwall/buntdb v1.1.2
	go.uber.org/multierr v1.6.0
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5
	golang.org/x/oauth2 v0.0.0-20210402161424-2e8d93401602
	golang.org/x/sys v0.0.0-20210910150752-751e447fb3d0
	golang.org/x/term v0.0.0-20210503060354-a79de5458b56
	golang.org/x/tools v0.1.7
	google.golang.org/grpc v1.38.0
	gopkg.in/flosch/pongo2.v3 v3.0.0-20141028000813-5e81b817a0c4
	gopkg.in/go-playground/validator.v9 v9.31.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/CloudyKit/fastprinter v0.0.0-20200109182630-33d98a066a53 // indirect
	github.com/agrison/mxj v0.0.0-20160310142625-1269f8afb3b4 // indirect
	github.com/bndr/gotabulate v1.1.2 // indirect
	github.com/clbanning/mxj v1.8.4 // indirect
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/golang/snappy v0.0.3 // indirect
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/jonboulle/clockwork v0.1.0 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lestrrat-go/strftime v1.0.1 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/tealeg/xlsx v1.0.5 // indirect
	github.com/tebeka/strftime v0.1.3 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
)

// fix go version
// replace (
// 	github.com/golang-jwt/jwt => github.com/golang-jwt/jwt v3.2.1+incompatible
// 	github.com/spf13/afero => github.com/spf13/afero v1.5.1
// 	golang.org/x/tools => github.com/golang/tools v0.1.0
// )
