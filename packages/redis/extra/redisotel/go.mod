module github.com/go-redis/redis/extra/redisotel

go 1.15

replace github.com/2637309949/dolphin/packages/redis => ../..

replace github.com/go-redis/redis/extra/rediscmd => ../rediscmd

require (
	github.com/go-redis/redis/extra/rediscmd v0.2.0
	github.com/2637309949/dolphin/packages/redis v8.4.4
	go.opentelemetry.io/otel v0.16.0
)
