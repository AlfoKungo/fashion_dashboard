# Tasks: Personal Men's Fashion Dashboard MVP

**Input**: Design documents from `specs/001-fashion-dashboard-mvp/`
**Prerequisites**: plan.md, spec.md, research.md, data-model.md, contracts/, quickstart.md
**Tests**: Included because the spec defines independent test criteria and quickstart.md recommends endpoint, route, selection, cleanup, and image coverage.
**Organization**: Tasks are grouped by user story to enable independent implementation and testing.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel because it touches different files and does not depend on incomplete tasks.
- **[Story]**: User story label for story phases only.
- Every task includes an exact file path.

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Initialize the Go project, directory layout, and local configuration artifacts.

- [X] T001 Create Go module `fashion_dashboard` in go.mod
- [X] T002 Create project directories from plan in cmd/server, internal/config, internal/db, internal/models, internal/repository, internal/processing, internal/scheduler, internal/web/templates, internal/web/static, internal/fetchers/articles, internal/fetchers/looks, internal/fetchers/items, tests/contract, tests/integration, and tests/unit
- [X] T003 Add MongoDB Go driver dependency in go.mod
- [X] T004 [P] Create .gitignore with Go, environment, editor, and OS patterns in .gitignore
- [X] T005 [P] Create environment sample with PORT, MONGODB_URI, MONGODB_DATABASE, APP_ENV, and DAILY_WORKFLOW_HOUR in .env.example
- [X] T006 [P] Create package documentation for runtime commands and feature scope in README.md

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Core types, configuration, database wiring, repository interfaces, routing shell, and shared helpers required by all user stories.

**CRITICAL**: No user story work can begin until this phase is complete.

- [X] T007 Implement environment configuration loading and defaults in internal/config/config.go
- [X] T008 Implement MongoDB client connection and graceful disconnect helpers in internal/db/mongo.go
- [X] T009 [P] Define Article, Look, Item, TrendSummary, and DailyCategory models in internal/models/models.go
- [X] T010 [P] Define repository interfaces for articles, looks, items, images, and trend summaries in internal/repository/interfaces.go
- [X] T011 Implement MongoDB repository structs and collection/index setup in internal/repository/mongo.go
- [X] T012 Implement common amount parsing and JSON error response helpers in internal/web/http_helpers.go
- [X] T013 Implement application dependency container and route registration shell in internal/web/server.go
- [X] T014 Implement process startup, config loading, database connection, route registration, and shutdown in cmd/server/main.go
- [X] T015 [P] Add reusable test fixture builders for Article, Look, and Item in tests/unit/fixtures_test.go

**Checkpoint**: Foundation ready; user story implementation can now begin.

---

## Phase 3: User Story 1 - View Today's Fashion Briefing (Priority: P1) MVP

**Goal**: Opening `/` shows today's articles, looks, item focus, navigation, current date, and footer as a complete daily briefing.

**Independent Test**: Open the default dashboard view and confirm it presents 4 article cards, 4 look cards, 6 item cards, current date, active Today navigation, and populated card metadata.

### Tests for User Story 1

- [X] T016 [P] [US1] Add unit tests for amount parsing defaults and 1-50 validation in tests/unit/http_helpers_test.go
- [X] T017 [P] [US1] Add handler test for `/api/articles?amount=4` response shape in tests/contract/articles_api_test.go
- [X] T018 [P] [US1] Add handler test for `/api/looks?amount=4` response shape in tests/contract/looks_api_test.go
- [X] T019 [P] [US1] Add handler test for `/api/items?amount=6` response shape and category field in tests/contract/items_api_test.go
- [X] T020 [P] [US1] Add integration test for Today view section counts and active navigation in tests/integration/today_view_test.go

### Implementation for User Story 1

- [X] T021 [P] [US1] Implement article list repository methods and in-memory fallback seed data in internal/repository/articles.go
- [X] T022 [P] [US1] Implement look list repository methods and in-memory fallback seed data in internal/repository/looks.go
- [X] T023 [P] [US1] Implement item list repository methods, daily category lookup, and in-memory fallback seed data in internal/repository/items.go
- [X] T024 [US1] Implement dashboard data service that loads 4 articles, 4 looks, 6 items, active category, current date, quote, and weather in internal/processing/dashboard.go
- [X] T025 [US1] Implement JSON handlers for `/api/articles`, `/api/looks`, and `/api/items` in internal/web/api_handlers.go
- [X] T026 [US1] Implement Today page handler for `/` in internal/web/page_handlers.go
- [X] T027 [US1] Create shared layout and navigation template in internal/web/templates/layout.html
- [X] T028 [US1] Create Today page template with top articles, daily inspiration, daily item focus, and footer in internal/web/templates/today.html
- [X] T029 [US1] Create dark dashboard CSS for cards, horizontal sections, tags, article overlays, navigation, and responsive behavior in internal/web/static/styles.css
- [X] T030 [US1] Register static asset serving and Today/API routes in internal/web/server.go

**Checkpoint**: User Story 1 is independently functional and demoable.

---

## Phase 4: User Story 2 - Browse Section-Specific Views (Priority: P2)

**Goal**: Users can navigate to dedicated Looks, Items, and Articles views from navigation and section call-to-action controls.

**Independent Test**: Select each top navigation item and each section call-to-action, then confirm the matching view opens and the active navigation state updates.

### Tests for User Story 2

- [X] T031 [P] [US2] Add integration test for `/looks` route rendering and active Looks navigation in tests/integration/looks_view_test.go
- [X] T032 [P] [US2] Add integration test for `/items` route rendering and active Items navigation in tests/integration/items_view_test.go
- [X] T033 [P] [US2] Add integration test for `/articles` route rendering and active Articles navigation in tests/integration/articles_view_test.go

### Implementation for User Story 2

- [X] T034 [US2] Add View All Looks and View All category links to Today template in internal/web/templates/today.html
- [X] T035 [P] [US2] Create Looks view template in internal/web/templates/looks.html
- [X] T036 [P] [US2] Create Items view template in internal/web/templates/items.html
- [X] T037 [P] [US2] Create Articles view template in internal/web/templates/articles.html
- [X] T038 [US2] Implement Looks, Items, and Articles page handlers in internal/web/page_handlers.go
- [X] T039 [US2] Register `/looks`, `/items`, and `/articles` routes in internal/web/server.go
- [X] T040 [US2] Extend CSS for dedicated view grids and active navigation states in internal/web/static/styles.css

**Checkpoint**: User Stories 1 and 2 both work independently.

---

## Phase 5: User Story 3 - Receive Fresh Daily Content (Priority: P3)

**Goal**: The system refreshes content daily, deduplicates source records, selects daily looks/items, generates trend summaries, and removes expired data.

**Independent Test**: Trigger a daily refresh cycle and confirm that daily selections exist, duplicate source records are avoided, and records older than 7 days are cleaned up.

### Tests for User Story 3

- [X] T041 [P] [US3] Add unit tests for deterministic daily category rotation in tests/unit/category_test.go
- [X] T042 [P] [US3] Add unit tests for deduplication by canonical links in tests/unit/dedup_test.go
- [X] T043 [P] [US3] Add unit tests for selecting exactly 4 looks and 6 items when enough valid records exist in tests/unit/selection_test.go
- [X] T044 [P] [US3] Add unit tests for 7-day cleanup eligibility in tests/unit/cleanup_test.go
- [X] T045 [US3] Add integration test for daily workflow orchestration in tests/integration/daily_workflow_test.go

### Implementation for User Story 3

- [X] T046 [P] [US3] Implement article fetcher interface and MVP sample/public-source fetch behavior in internal/fetchers/articles/fetcher.go
- [X] T047 [P] [US3] Implement look fetcher interface and MVP sample/public-source fetch behavior in internal/fetchers/looks/fetcher.go
- [X] T048 [P] [US3] Implement item fetcher interface and MVP sample/public-source fetch behavior in internal/fetchers/items/fetcher.go
- [X] T049 [US3] Implement content normalization, canonical link hashing, read-time defaults, and tag defaults in internal/processing/normalize.go
- [X] T050 [US3] Implement deterministic daily category rotation and daily look/item selection in internal/processing/selection.go
- [X] T051 [US3] Implement trend summary generation placeholder and persistence integration in internal/processing/trends.go
- [X] T052 [US3] Implement retention cleanup for articles, looks, items, and trend summaries in internal/processing/cleanup.go
- [X] T053 [US3] Implement daily workflow orchestration across fetch, normalize, store, select, summarize, and cleanup in internal/scheduler/workflow.go
- [X] T054 [US3] Implement configured daily scheduler loop using DAILY_WORKFLOW_HOUR in internal/scheduler/scheduler.go
- [X] T055 [US3] Start scheduler from cmd/server/main.go when APP_ENV is not test

**Checkpoint**: User Stories 1, 2, and 3 work independently.

---

## Phase 6: User Story 4 - Handle Content and Image Gaps Gracefully (Priority: P4)

**Goal**: The dashboard remains readable when image storage fails or when fewer complete records are available than the target count.

**Independent Test**: Use records with missing or unavailable images and short content lists, then confirm pages and cards render without broken layouts or empty placeholders.

### Tests for User Story 4

- [X] T056 [P] [US4] Add contract tests for `/images/articles/{id}`, `/images/looks/{id}`, and `/images/items/{id}` byte, redirect, and not-found behavior in tests/contract/images_test.go
- [X] T057 [P] [US4] Add integration test for partial content rendering without empty cards in tests/integration/partial_content_test.go

### Implementation for User Story 4

- [X] T058 [US4] Implement image download and content-type capture helper with URL-only fallback in internal/processing/images.go
- [X] T059 [US4] Implement image lookup repository methods for articles, looks, and items in internal/repository/images.go
- [X] T060 [US4] Implement image handlers for `/images/articles/{id}`, `/images/looks/{id}`, and `/images/items/{id}` in internal/web/image_handlers.go
- [X] T061 [US4] Register image routes in internal/web/server.go
- [X] T062 [US4] Update templates to render only complete cards and avoid broken image elements in internal/web/templates/today.html
- [X] T063 [US4] Update dedicated view templates to render partial content safely in internal/web/templates/looks.html, internal/web/templates/items.html, and internal/web/templates/articles.html
- [X] T064 [US4] Add CSS fallbacks for missing imagery and compact partial sections in internal/web/static/styles.css

**Checkpoint**: All user stories are independently functional.

---

## Phase 7: Polish & Cross-Cutting Concerns

**Purpose**: Final validation, documentation, and implementation cleanup across all stories.

- [X] T065 [P] Add repository unit tests for MongoDB query filters and index setup behavior in tests/unit/repository_test.go
- [X] T066 [P] Add README instructions for setup, environment variables, running, testing, and scheduled workflow behavior in README.md
- [X] T067 Run gofmt across Go source and tests in cmd, internal, and tests
- [X] T068 Run `go test ./...` and fix any failures in cmd, internal, and tests
- [X] T069 Validate quickstart commands and document any environment-dependent limitations in specs/001-fashion-dashboard-mvp/quickstart.md
- [X] T070 Review implementation against FR-001 through FR-028 and update this tasks file with completed checkboxes in specs/001-fashion-dashboard-mvp/tasks.md

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies; start immediately.
- **Foundational (Phase 2)**: Depends on Setup completion; blocks all user stories.
- **User Stories (Phase 3+)**: Depend on Foundational completion.
- **Polish (Phase 7)**: Depends on all desired user stories being complete.

### User Story Dependencies

- **US1 (P1)**: Starts after Foundation; no dependency on other stories.
- **US2 (P2)**: Starts after Foundation; uses shared layout and data loading from US1 for best continuity, but routes can be tested independently.
- **US3 (P3)**: Starts after Foundation; provides live freshness behind the same repository/service contracts used by US1.
- **US4 (P4)**: Starts after Foundation; integrates with UI and repository contracts from US1/US2 but is independently testable with missing images and partial data.

### Within Each User Story

- Tests before implementation.
- Repository/data methods before services.
- Services before handlers.
- Handlers before route registration.
- Templates and CSS before final route validation.

### Parallel Opportunities

- Setup tasks T004-T006 can run in parallel after T001-T003.
- Foundational tasks T009, T010, and T015 can run in parallel after T007-T008.
- US1 tests T016-T020 can run in parallel.
- US1 repository tasks T021-T023 can run in parallel.
- US2 tests T031-T033 and templates T035-T037 can run in parallel.
- US3 tests T041-T044 and fetchers T046-T048 can run in parallel.
- US4 tests T056-T057 can run in parallel.
- Polish tasks T065-T066 can run in parallel.

---

## Parallel Example: User Story 1

```text
Task: "T016 [P] [US1] Add unit tests for amount parsing defaults and 1-50 validation in tests/unit/http_helpers_test.go"
Task: "T017 [P] [US1] Add handler test for /api/articles?amount=4 response shape in tests/contract/articles_api_test.go"
Task: "T018 [P] [US1] Add handler test for /api/looks?amount=4 response shape in tests/contract/looks_api_test.go"
Task: "T019 [P] [US1] Add handler test for /api/items?amount=6 response shape and category field in tests/contract/items_api_test.go"
Task: "T020 [P] [US1] Add integration test for Today view section counts and active navigation in tests/integration/today_view_test.go"
```

## Parallel Example: User Story 3

```text
Task: "T046 [P] [US3] Implement article fetcher interface and MVP sample/public-source fetch behavior in internal/fetchers/articles/fetcher.go"
Task: "T047 [P] [US3] Implement look fetcher interface and MVP sample/public-source fetch behavior in internal/fetchers/looks/fetcher.go"
Task: "T048 [P] [US3] Implement item fetcher interface and MVP sample/public-source fetch behavior in internal/fetchers/items/fetcher.go"
```

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1 setup.
2. Complete Phase 2 foundation.
3. Complete Phase 3 User Story 1.
4. Stop and validate `/`, `/api/articles`, `/api/looks`, and `/api/items`.
5. Demo the Today dashboard before adding dedicated browsing, daily refresh, or image resilience.

### Incremental Delivery

1. Add US2 to make navigation and dedicated views complete.
2. Add US3 to replace static/fallback content with daily refresh behavior.
3. Add US4 to harden the UI and image handling against partial external failures.
4. Complete polish tasks and run the quickstart validation.
