package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

type Tracer struct {
	trace.TracerProvider
}

func (*Tracer) TraceQueryStart(ctx context.Context, _ *pgx.Conn, _ pgx.TraceQueryStartData) context.Context {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		semconv.DBSystemPostgreSQL,
	)

	return ctx
}

func (*Tracer) TraceQueryEnd(ctx context.Context, _ *pgx.Conn, data pgx.TraceQueryEndData) {
	span := trace.SpanFromContext(ctx)

	if data.Err != nil {
		span.RecordError(data.Err)
	}

	span.End()
}
