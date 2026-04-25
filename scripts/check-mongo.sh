#!/usr/bin/env bash
set -euo pipefail

if [ -f .env ]; then
  set -a
  # shellcheck disable=SC1091
  source .env
  set +a
fi

if [ -z "${MONGODB_URI:-}" ]; then
  echo "ERROR: MONGODB_URI is not set. Put your Atlas connection string in .env or export it." >&2
  exit 1
fi

host="$(
  node -e 'const u = new URL(process.env.MONGODB_URI); console.log(u.hostname)'
)"

if [ "$host" = "CLUSTER.mongodb.net" ] || [ "$host" = "cluster.mongodb.net" ] || [[ "$host" == *"<"* ]] || [[ "$host" == *">"* ]]; then
  echo "ERROR: MONGODB_URI still contains a placeholder host: $host" >&2
  echo "Use the real Atlas host, usually like cluster0.xxxxx.mongodb.net." >&2
  exit 1
fi

if [[ "$MONGODB_URI" == mongodb+srv://* ]]; then
  echo "Checking SRV DNS: _mongodb._tcp.$host"
  if ! dig +short SRV "_mongodb._tcp.$host" >/tmp/fashion-dashboard-mongo-srv.txt; then
    echo "ERROR: DNS lookup failed for _mongodb._tcp.$host" >&2
    exit 1
  fi
  if [ ! -s /tmp/fashion-dashboard-mongo-srv.txt ]; then
    echo "ERROR: No MongoDB SRV records found for _mongodb._tcp.$host" >&2
    echo "Check that the Atlas cluster hostname is exact." >&2
    exit 1
  fi
  cat /tmp/fashion-dashboard-mongo-srv.txt
fi

db="${MONGODB_DATABASE:-fashion_dashboard}"
echo "Pinging MongoDB database: $db"
mongosh "$MONGODB_URI" --quiet --eval "db.getSiblingDB('$db').runCommand({ ping: 1 })"
