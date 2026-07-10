# cut-it
A link shortener built in Go, from scratch, with a focus in infrastructure.

## Architecture
```
cut-it/
├── main.go              # server setup, dependency wiring
├── go.mod
├── dockerfile            # multi-stage build
├── handlers/
│   ├── shorten.go        # POST /shorten
│   ├── redirect.go       # GET /{slug}
│   └── routes.go         # Handler struct, route registration
├── store/
│   └── store.go          # in-memory, mutex-protected key-value store
└── models/
    └── models.go          # Link struct
```

## Running locally
```bash
go run .
```

## Running with Docker
```bash
docker build . -t cut-it
docker run -p 8080:8080 cut-it
```

## API

### `POST /shorten`
Request:
```json
{ "original_url": "https://example.com/some/long/path" }
```
 
Success (201):
```json
{ "shortened_url": "http://localhost:8080/aB3xY9z" }
```

Failure cases: 400 (malformed JSON), 422 (invalid URL — bad/missing scheme, empty or self-referential host), 500 (server-side failure generating a unique slug).

### `GET /{slug}`
 
- Found → 307 redirect to the original URL via the `Location` header.
- Not found → 404.