# go-ride-utils

Shared common values package for Go Ride services.

## Includes

- Kafka topic name constants under [`kafkatopics/`](kafkatopics/kafkatopics.go)
- Kafka event contract structs under [`events/`](events/)
- Shared HTTP header/field name constants under [`httpheaders/`](httpheaders/httpheaders.go)
- AWS Secrets Manager fetch helper under [`awssecrets/`](awssecrets/awssecrets.go) — see "Deployment usage" below

## Deployment usage

`awssecrets.FetchJSON(ctx, secretName)` is what
[`go-ride-kafka-consumers`](https://github.com/shawon-kanji/go-ride-kafka-consumers)
and [`go-ride-backend`](https://github.com/shawon-kanji/go-ride-backend)
call from their `internal/config/config.go` `Load()` functions to fetch DB
credentials and the JWT secret from AWS Secrets Manager in staging/
production, authenticated via IRSA (no explicit AWS credentials or region
wiring needed — it uses the SDK's default credential chain, which IRSA
populates on the pod automatically). Locally, consumers fall back to plain
env vars and never call this package. The Secrets Manager entries
themselves are created by
[`go-ride-infra`](https://github.com/shawon-kanji/go-ride-infra)'s
[`terraform/modules/secrets-manager`](https://github.com/shawon-kanji/go-ride-infra/blob/main/terraform/modules/secrets-manager)
— see that repo's
[`docs/architecture.md`](https://github.com/shawon-kanji/go-ride-infra/blob/main/docs/architecture.md)
for the full secrets-flow picture.

## Local usage in sibling repos

Add to `go.mod`:

```go
require github.com/shawon-kanji/go-ride-utils v0.0.0
replace github.com/shawon-kanji/go-ride-utils => ../go-ride-utils
```

Or pull a tagged release:

```bash
go get github.com/shawon-kanji/go-ride-utils@vX.Y.Z
```
