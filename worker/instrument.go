package worker

import "go.opentelemetry.io/otel"

var orderTracer = otel.Tracer("Order:worker")
