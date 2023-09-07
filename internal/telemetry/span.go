package telemetry

import (
	"go.opentelemetry.io/otel/trace"
)

type SpanContext struct {
	RequestID string `json:"request_id,omitempty"`
	TraceID   string `json:"trace_id,omitempty"`
	SpanID    string `json:"span_id,omitempty"`
}

func NewSpanContext(ctx SpanContext) (trace.SpanContext, error) {
	traceID, err := trace.TraceIDFromHex(ctx.TraceID)
	if err != nil {
		return trace.SpanContext{}, err
	}
	spanID, err := trace.SpanIDFromHex(ctx.SpanID)
	if err != nil {
		return trace.SpanContext{}, err
	}
	var spanContextConfig trace.SpanContextConfig
	spanContextConfig.TraceID = traceID
	spanContextConfig.SpanID = spanID
	spanContextConfig.TraceFlags = 01
	spanContextConfig.Remote = false
	var spanContext trace.SpanContext
	spanContext = trace.NewSpanContext(spanContextConfig)
	return spanContext, nil
}
