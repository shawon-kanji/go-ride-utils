# go-ride-utils

Shared common values package for Go Ride services.

## Includes

- Kafka topic name constants under `kafkatopics/`
- Kafka event contract structs under `events/`
- Shared HTTP header/field name constants under `httpheaders/`

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
