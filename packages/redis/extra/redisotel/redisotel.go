package redisotel

import (
	"context"

	"github.com/2637309949/dolphin/packages/redis"
	"github.com/go-redis/redis/extra/rediscmd"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("github.com/go-redis/redis")

type TracingHook struct{}

var _ redis.Hook = TracingHook{}

func (TracingHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	if !trace.SpanFromContext(ctx).IsRecording() {
		return ctx, nil
	}

	ctx, span := tracer.Start(ctx, cmd.FullName())
	span.SetAttributes(
		label.String("db.system", "redis"),
		label.String("db.statement", rediscmd.CmdString(cmd)),
	)

	return ctx, nil
}

func (TracingHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	span := trace.SpanFromContext(ctx)
	if err := cmd.Err(); err != nil {
		recordError(ctx, span, err)
	}
	span.End()
	return nil
}

func (TracingHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	if !trace.SpanFromContext(ctx).IsRecording() {
		return ctx, nil
	}

	summary, cmdsString := rediscmd.CmdsString(cmds)

	ctx, span := tracer.Start(ctx, "pipeline "+summary)
	span.SetAttributes(
		label.String("db.system", "redis"),
		label.Int("db.redis.num_cmd", len(cmds)),
		label.String("db.statement", cmdsString),
	)

	return ctx, nil
}

func (TracingHook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	span := trace.SpanFromContext(ctx)
	if err := cmds[0].Err(); err != nil {
		recordError(ctx, span, err)
	}
	span.End()
	return nil
}

func recordError(ctx context.Context, span trace.Span, err error) {
	if err != redis.Nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
}
