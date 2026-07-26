// Package httpheaders centralizes HTTP header and JSON field names for
// cross-cutting concerns (idempotency, correlation) shared across go-ride's
// HTTP-facing services. Struct tags must be literal strings in Go, so the
// *Field constants are for non-tag usages (header lookups, error messages),
// not a replacement for JSON tags.
package httpheaders

const (
	Idempotency = "Idempotency-Key"
	Correlation = "X-Correlation-ID"

	IdempotencyKeyField = "idempotency_key"
	CorrelationIDField  = "correlation_id"
)
