package ntracing

import (
	"context"

	opentracing "github.com/opentracing/opentracing-go"
)

type spanKey string

// SpanKey is the key to use to store the span in request context
var SpanKey = spanKey("span")

// NewChildSpanFromContext Generates a child span with given name
// if a span is found on the context
func NewChildSpanFromContext(
	ctx context.Context,
	name string,
) (opentracing.Span, bool) {
	if span, ok := ctx.Value(SpanKey).(opentracing.Span); ok && span != nil {
		var childSpan = opentracing.StartSpan(
			name,
			opentracing.ChildOf(span.Context()),
		)
		return childSpan, true
	}
	return nil, false
}

// NewChildSpanAndContext generates a child span with given name if a span is
// found on the context and creates a new context from that child span
func NewChildSpanAndContext(
	ctx context.Context,
	name string,
) (context.Context, opentracing.Span) {
	if span, ok := NewChildSpanFromContext(ctx, name); ok {
		return context.WithValue(ctx, SpanKey, span), span
	}
	return ctx, nil
}
