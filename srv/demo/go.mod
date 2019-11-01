module example

go 1.13

replace github.com/2637309949/dolphin => ../../

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43

replace github.com/go-xorm/core v0.6.3 => xorm.io/core v0.6.3

require (
	github.com/2637309949/dolphin v0.0.0-00010101000000-000000000000
	github.com/bu/gin-method-override v0.0.0-20180528072813-3f56fc145a4b
	github.com/clbanning/mxj v1.8.4 // indirect
	github.com/cockroachdb/apd v1.1.0 // indirect
	github.com/denisenkom/go-mssqldb v0.0.0-20191001013358-cfbb681360f0 // indirect
	github.com/ekyoung/gin-nice-recovery v0.0.0-20160510022553-1654dca486db
	github.com/gin-gonic/gin v1.4.0
	github.com/go-errors/errors v1.0.1 // indirect
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/core v0.6.3 // indirect
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/jackc/fake v0.0.0-20150926172116-812a484cc733 // indirect
	github.com/jackc/pgx v3.6.0+incompatible // indirect
	github.com/lib/pq v1.2.0 // indirect
	github.com/mattn/go-sqlite3 v1.11.0 // indirect
	github.com/shopspring/decimal v0.0.0-20191009025716-f1972eb1d1f5 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/xormplus/xorm v0.0.0-20190926190243-42377f593eb1
	github.com/ziutek/mymysql v1.5.4 // indirect
	go.uber.org/goleak v0.10.0 // indirect
	gopkg.in/guregu/null.v3 v3.4.0
)
