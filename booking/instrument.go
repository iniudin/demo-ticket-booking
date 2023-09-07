package booking

import "go.opentelemetry.io/otel"

var concertTracer = otel.Tracer("Concert:service")
var orderTracer = otel.Tracer("Order:service")
