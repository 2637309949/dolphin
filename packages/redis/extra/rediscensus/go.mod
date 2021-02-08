module github.com/go-redis/redis/extra/rediscensus

go 1.15

replace github.com/go-redis/redis/extra/rediscmd => ../rediscmd

require (
	github.com/go-redis/redis/extra/rediscmd v0.2.0
	github.com/2637309949/dolphin/packages/redis v8.4.4
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	go.opencensus.io v0.22.6
)
