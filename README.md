# Personal Men's Fashion Dashboard

Server-rendered Go MVP for a personal men's fashion dashboard. It shows a Today briefing with fashion articles, outfit inspiration, and a daily item focus, plus dedicated Looks, Items, and Articles views.

## Run

```sh
cp .env.example .env
go run ./cmd/server
```

Open `http://localhost:8080/`.

## Environment

- `PORT`: HTTP port, defaults to `8080`
- `MONGODB_URI`: optional MongoDB Atlas connection string
- `MONGODB_DATABASE`: defaults to `fashion_dashboard`
- `APP_ENV`: set to `test` to avoid starting the scheduler
- `DAILY_WORKFLOW_HOUR`: local hour for daily refresh, defaults to `7`

When MongoDB is not configured, the app serves curated in-memory fallback content so the UI and contracts remain testable.

The server loads `.env` from the repository root automatically. Shell environment variables take precedence over values in `.env`.

## Local MongoDB

Start a local MongoDB container with known development credentials:

```sh
docker compose up -d mongo
```

The included `.env` points the app at:

```text
mongodb://fashion:dashboard@localhost:27017/fashion_dashboard?authSource=admin
```

## Mongo CLI

Connect with `mongosh` using the same URI:

```sh
mongosh "$MONGODB_URI" --eval 'use fashion_dashboard; show collections'
```

Or paste the Atlas URI directly:

```sh
mongosh 'mongodb+srv://USER:PASSWORD@CLUSTER.mongodb.net/fashion_dashboard?retryWrites=true&w=majority'
```

Validate the local Mongo CLI, DNS, and credentials:

```sh
bash scripts/check-mongo.sh
```

## Test

```sh
go test ./...
```
