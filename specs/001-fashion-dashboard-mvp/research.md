# Research: Personal Men's Fashion Dashboard MVP

## Decision: Use a single Go server process

**Rationale**: The feature scope combines HTTP routing, server-rendered pages, JSON data endpoints, scheduled content workflows, content normalization, image handling, and storage. A single service keeps the personal MVP simple and matches the provided design constraints.

**Alternatives considered**: Separate frontend/backend services were rejected because the MVP explicitly avoids a frontend framework and client-side external fetching. A CLI-only workflow was rejected because the primary user value is an interactive dashboard.

## Decision: Use Go standard library routing first, with Chi as optional fallback

**Rationale**: The route set is small and predictable: 4 page routes, 3 data endpoints, and 3 image-serving route groups. Standard library routing is sufficient for the MVP and minimizes dependencies.

**Alternatives considered**: Chi was considered for cleaner path parameter handling. It should be introduced only if image route parsing or middleware organization becomes materially clearer.

## Decision: Use server-rendered HTML templates and custom CSS

**Rationale**: The UI is content-heavy, the design is static enough for server rendering, and the MVP explicitly excludes a frontend framework. Server rendering also enforces that the UI reads from backend-owned data rather than external browser-side sources.

**Alternatives considered**: React or another client framework was rejected by the MVP constraints. Static HTML only was rejected because daily content, dynamic routes, and stored images require server data.

## Decision: Use MongoDB Atlas as the persistence layer

**Rationale**: `design.md` specifies MongoDB Atlas and document-shaped records for articles, looks, items, and trend summaries. The content records include optional metadata and image bytes, which fit document storage for an MVP.

**Alternatives considered**: File storage was rejected because deduplication, date queries, cleanup, and selected-for-day lookups would become ad hoc. A relational database was not selected because the supplied design already defines document collections.

## Decision: Store image bytes when possible and retain source image URLs

**Rationale**: Stored images make dashboard rendering more reliable if source images change or rate-limit later. Retaining the original URL gives a graceful fallback when download or storage fails.

**Alternatives considered**: URL-only images were rejected because they make the dashboard fragile. Mandatory byte storage was rejected because external image downloads can fail and should not block content display.

## Decision: Implement the daily workflow inside the service

**Rationale**: A personal MVP can run a lightweight internal scheduled loop configured by `DAILY_WORKFLOW_HOUR`. This avoids external scheduler setup while satisfying daily refresh, trend summary, selection, and cleanup requirements.

**Alternatives considered**: A hosted cron or queue system was rejected as unnecessary MVP infrastructure. Manual refresh only was rejected because the feature requires daily freshness without manual maintenance.

## Decision: Rotate daily item category deterministically by date

**Rationale**: A deterministic date-based category rotation keeps the UI fresh while making behavior reproducible for tests and operations. The initial category list comes from `design.md`.

**Alternatives considered**: Random selection was rejected because it is harder to test and explain. User personalization was rejected by MVP constraints.

## Decision: Bound API amount requests to 1 through 50

**Rationale**: This directly satisfies the spec and keeps endpoint responses manageable for a personal dashboard. Invalid values should return a clear client error instead of silently choosing surprising limits.

**Alternatives considered**: Unbounded requests were rejected due to performance and UI risks. Silently coercing invalid values was rejected because contract tests should be unambiguous.

## Decision: Prefer complete cards over placeholders when content is short

**Rationale**: The spec requires the dashboard to avoid broken cards when fewer records are available. Showing fewer complete cards communicates partial availability more cleanly than empty placeholders.

**Alternatives considered**: Placeholder cards were rejected because they reduce perceived quality and can obscure real content shortages. Blocking the page until exact counts exist was rejected because it harms resilience.
