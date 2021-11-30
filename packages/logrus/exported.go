package logrus

import (
	"context"

	"github.com/2637309949/dolphin/packages/trace/tracespec"
	"github.com/sirupsen/logrus"
)

const traceField = "traceid"

// Trace logs a message at level Trace on the standard logger.
func Trace(ctx context.Context, args ...interface{}) {
	span, ok := ctx.(tracespec.SpanContext)
	if ok {
		logrus.WithField(traceField, span.TraceId()).Trace(args...)
	} else {
		logrus.Trace(args...)
	}
}

// Debug logs a message at level Debug on the standard logger.
func Debug(ctx context.Context, args ...interface{}) {
	span, ok := ctx.(tracespec.SpanContext)
	if ok {
		logrus.WithField(traceField, span.TraceId()).Debug(args...)
	} else {
		logrus.Debug(args...)
	}
}

// Print logs a message at level Info on the standard logger.
func Print(ctx context.Context, args ...interface{}) {
	span, ok := ctx.(tracespec.SpanContext)
	if ok {
		logrus.WithField(traceField, span.TraceId()).Print(args...)
	} else {
		logrus.Print(args...)
	}
}

// Info logs a message at level Info on the standard logger.
func Info(ctx context.Context, args ...interface{}) {
	span, ok := ctx.(tracespec.SpanContext)
	if ok {
		logrus.WithField(traceField, span.TraceId()).Info(args...)
	} else {
		logrus.Info(args...)
	}
}

// Warn logs a message at level Warn on the standard logger.
func Warn(ctx context.Context, args ...interface{}) {
	span, ok := ctx.(tracespec.SpanContext)
	if ok {
		logrus.WithField(traceField, span.TraceId()).Warn(args...)
	} else {
		logrus.Warn(args...)
	}
}

// Warning logs a message at level Warn on the standard logger.
func Warning(ctx context.Context, args ...interface{}) {
	span, ok := ctx.(tracespec.SpanContext)
	if ok {
		logrus.WithField(traceField, span.TraceId()).Warning(args...)
	} else {
		logrus.Warning(args...)
	}
}

// Error logs a message at level Error on the standard logger.
func Error(ctx context.Context, args ...interface{}) {
	span, ok := ctx.(tracespec.SpanContext)
	if ok {
		logrus.WithField(traceField, span.TraceId()).Error(args...)
	} else {
		logrus.Error(args...)
	}
}

// Panic logs a message at level Panic on the standard logger.
func Panic(ctx context.Context, args ...interface{}) {
	span, ok := ctx.(tracespec.SpanContext)
	if ok {
		logrus.WithField(traceField, span.TraceId()).Panic(args...)
	} else {
		logrus.Panic(args...)
	}
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(ctx context.Context, args ...interface{}) {
	span, ok := ctx.(tracespec.SpanContext)
	if ok {
		logrus.WithField(traceField, span.TraceId()).Fatal(args...)
	} else {
		logrus.Fatal(args...)
	}
}

// Tracef logs a message at level Trace on the standard logger.
func Tracef(ctx context.Context, format string, args ...interface{}) {
	span, ok := ctx.(tracespec.SpanContext)
	if ok {
		logrus.WithField(traceField, span.TraceId()).Tracef(format, args...)
	} else {
		logrus.Tracef(format, args...)
	}
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(ctx context.Context, format string, args ...interface{}) {
	span, ok := ctx.(tracespec.SpanContext)
	if ok {
		logrus.WithField(traceField, span.TraceId()).Debugf(format, args...)
	} else {
		logrus.Debugf(format, args...)
	}
}

// Printf logs a message at level Info on the standard logger.
func Printf(ctx context.Context, format string, args ...interface{}) {
	span, ok := ctx.(tracespec.SpanContext)
	if ok {
		logrus.WithField(traceField, span.TraceId()).Printf(format, args...)
	} else {
		logrus.Printf(format, args...)
	}
}

// Infof logs a message at level Info on the standard logger.
func Infof(ctx context.Context, format string, args ...interface{}) {
	span, ok := ctx.(tracespec.SpanContext)
	if ok {
		logrus.WithField(traceField, span.TraceId()).Infof(format, args...)
	} else {
		logrus.Infof(format, args...)
	}
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(ctx context.Context, format string, args ...interface{}) {
	span, ok := ctx.(tracespec.SpanContext)
	if ok {
		logrus.WithField(traceField, span.TraceId()).Warnf(format, args...)
	} else {
		logrus.Warnf(format, args...)
	}
}

// Warningf logs a message at level Warn on the standard logger.
func Warningf(ctx context.Context, format string, args ...interface{}) {
	span, ok := ctx.(tracespec.SpanContext)
	if ok {
		logrus.WithField(traceField, span.TraceId()).Warningf(format, args...)
	} else {
		logrus.Warningf(format, args...)
	}
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(ctx context.Context, format string, args ...interface{}) {
	span, ok := ctx.(tracespec.SpanContext)
	if ok {
		logrus.WithField(traceField, span.TraceId()).Errorf(format, args...)
	} else {
		logrus.Errorf(format, args...)
	}
}

// Panicf logs a message at level Panic on the standard logger.
func Panicf(ctx context.Context, format string, args ...interface{}) {
	span, ok := ctx.(tracespec.SpanContext)
	if ok {
		logrus.WithField(traceField, span.TraceId()).Panicf(format, args...)
	} else {
		logrus.Panicf(format, args...)
	}
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalf(ctx context.Context, format string, args ...interface{}) {
	span, ok := ctx.(tracespec.SpanContext)
	if ok {
		logrus.WithField(traceField, span.TraceId()).Fatalf(format, args...)
	} else {
		logrus.Fatalf(format, args...)
	}
}

// Traceln logs a message at level Trace on the standard logger.
func Traceln(ctx context.Context, args ...interface{}) {
	span, ok := ctx.(tracespec.SpanContext)
	if ok {
		logrus.WithField(traceField, span.TraceId()).Traceln(args...)
	} else {
		logrus.Traceln(args...)
	}
}

// Debugln logs a message at level Debug on the standard logger.
func Debugln(ctx context.Context, args ...interface{}) {
	span, ok := ctx.(tracespec.SpanContext)
	if ok {
		logrus.WithField(traceField, span.TraceId()).Debugln(args...)
	} else {
		logrus.Debugln(args...)
	}
}

// Println logs a message at level Info on the standard logger.
func Println(ctx context.Context, args ...interface{}) {
	span, ok := ctx.(tracespec.SpanContext)
	if ok {
		logrus.WithField(traceField, span.TraceId()).Println(args...)

	} else {
		logrus.Println(args...)
	}
}

// Infoln logs a message at level Info on the standard logger.
func Infoln(ctx context.Context, args ...interface{}) {
	span, ok := ctx.(tracespec.SpanContext)
	if ok {
		logrus.WithField(traceField, span.TraceId()).Infoln(args...)
	} else {
		logrus.Infoln(args...)
	}
}

// Warnln logs a message at level Warn on the standard logger.
func Warnln(ctx context.Context, args ...interface{}) {
	span, ok := ctx.(tracespec.SpanContext)
	if ok {
		logrus.WithField(traceField, span.TraceId()).Warnln(args...)
	} else {
		logrus.Warnln(args...)
	}
}

// Warningln logs a message at level Warn on the standard logger.
func Warningln(ctx context.Context, args ...interface{}) {
	span, ok := ctx.(tracespec.SpanContext)
	if ok {
		logrus.WithField(traceField, span.TraceId()).Warningln(args...)
	} else {
		logrus.Warningln(args...)
	}
}

// Errorln logs a message at level Error on the standard logger.
func Errorln(ctx context.Context, args ...interface{}) {
	span, ok := ctx.(tracespec.SpanContext)
	if ok {
		logrus.WithField(traceField, span.TraceId()).Errorln(args...)
	} else {
		logrus.Errorln(args...)
	}
}

// Panicln logs a message at level Panic on the standard logger.
func Panicln(ctx context.Context, args ...interface{}) {
	span, ok := ctx.(tracespec.SpanContext)
	if ok {
		logrus.WithField(traceField, span.TraceId()).Panicln(args...)
	} else {
		logrus.Panicln(args...)
	}
}

// Fatalln logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalln(ctx context.Context, args ...interface{}) {
	span, ok := ctx.(tracespec.SpanContext)
	if ok {
		logrus.WithField(traceField, span.TraceId()).Fatalln(args...)
	} else {
		logrus.Fatalln(args...)
	}
}
