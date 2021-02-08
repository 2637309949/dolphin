package rediscensus

import (
	"context"

	"github.com/2637309949/dolphin/packages/redis"
	"github.com/go-redis/redis/extra/rediscmd"
	"go.opencensus.io/trace"
)

type TracingHook struct{}

var _ redis.Hook = TracingHook{}

func (TracingHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	ctx, span := trace.StartSpan(ctx, cmd.FullName())
	span.AddAttributes(trace.StringAttribute("db.system", "redis"),
		trace.StringAttribute("redis.cmd", rediscmd.CmdString(cmd)))

	return ctx, nil
}

func (TracingHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	span := trace.FromContext(ctx)
	if err := cmd.Err(); err != nil {
		recordErrorOnOCSpan(ctx, span, err)
	}
	span.End()
	return nil
}

func (TracingHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	return ctx, nil
}

func (TracingHook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	return nil
}

func recordErrorOnOCSpan(ctx context.Context, span *trace.Span, err error) {
	if err != redis.Nil {
		span.AddAttributes(trace.BoolAttribute("error", true))
		span.Annotate([]trace.Attribute{trace.StringAttribute("Error", "redis error")}, err.Error())
	}
}
