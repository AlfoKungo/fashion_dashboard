# Quickstart: Personal Men's Fashion Dashboard MVP

## Prerequisites

- Go 1.26+
- MongoDB Atlas connection string
- Network access from the server to configured content/image sources

## Environment

Create a local `.env` or export these variables before running:

```sh
PORT=8080
MONGODB_URI=
MONGODB_DATABASE=fashion_dashboard
APP_ENV=development
DAILY_WORKFLOW_HOUR=7
```

## Run Locally

```sh
go run ./cmd/server
```

Open:

- `http://localhost:8080/`
- `http://localhost:8080/looks`
- `http://localhost:8080/items`
- `http://localhost:8080/articles`

## Verify Data Endpoints

```sh
curl 'http://localhost:8080/api/articles?amount=4'
curl 'http://localhost:8080/api/looks?amount=4'
curl 'http://localhost:8080/api/items?amount=6'
curl 'http://localhost:8080/api/items?amount=51'
```

Expected behavior:

- Valid amount requests return JSON arrays with up to the requested count.
- Invalid amount requests return a client error.
- Item responses include the active daily category.

## Verify Images

Use image paths from endpoint responses or rendered HTML:

```sh
curl -I 'http://localhost:8080/images/articles/{id}'
curl -I 'http://localhost:8080/images/looks/{id}'
curl -I 'http://localhost:8080/images/items/{id}'
```

Expected behavior:

- Stored images return bytes with an image content type.
- Records without stored bytes redirect or fall back to the original image URL.

## Test

```sh
go test ./...
```

Recommended coverage:

- Amount validation for articles, looks, and items.
- Route rendering for Today, Looks, Items, and Articles.
- Daily selection counts for looks and items.
- Deduplication by canonical links.
- Cleanup of records older than 7 days.
- Image byte serving and URL fallback.
