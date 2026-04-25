# Implementation Plan: Personal Men's Fashion Dashboard MVP

**Branch**: `001-fashion-dashboard-mvp` | **Date**: 2026-04-25 | **Spec**: [spec.md](./spec.md)  
**Input**: Feature specification from `specs/001-fashion-dashboard-mvp/spec.md`

## Summary

Build a lightweight personal men's fashion dashboard that renders a dark, card-based Today view with articles, looks, and daily item focus, plus dedicated Looks, Items, and Articles views. The implementation will use a single Go web service with server-rendered HTML, MongoDB Atlas storage, internal daily content workflows, image storage/fallback handling, and simple JSON endpoints for section data.

## Technical Context

**Language/Version**: Go 1.26.2 locally; target Go 1.26+ for implementation  
**Primary Dependencies**: Go standard library HTTP server and templates; official MongoDB Go driver; optional Chi router only if route handling becomes clearer than standard library routing  
**Storage**: MongoDB Atlas database `fashion_dashboard` with collections for articles, looks, items, and trend summaries  
**Testing**: `go test ./...`; HTTP handler tests with `httptest`; repository tests with mocked interfaces or integration tests gated by MongoDB test configuration  
**Target Platform**: Local/deployable web service running as a single server process  
**Project Type**: Server-rendered web application with JSON data endpoints  
**Performance Goals**: Today view identifies all primary sections within 10 seconds for a user; content amount endpoints return up to 50 records quickly enough for interactive browsing; daily refresh completes before the configured morning display time on 95% of scheduled runs  
**Constraints**: No authentication, no personalization, no frontend framework, no client-side fetching from external sources, no real-time updates, keep source definitions during cleanup, retain content for 7 days  
**Scale/Scope**: Personal MVP with 4 article cards, 4 daily look cards, 6 daily item cards, maximum 50 records per data request, 4 web views, 3 JSON data endpoints, 3 image-serving route groups, and one daily workflow

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

The constitution at `.specify/memory/constitution.md` still contains placeholder principles and no ratified project-specific gates. No enforceable constitutional constraints were found. Planning proceeds under the explicit constraints in `spec.md` and `design.md`.

**Initial Gate Result**: PASS, no active gates to evaluate.

## Project Structure

### Documentation (this feature)

```text
specs/001-fashion-dashboard-mvp/
├── plan.md
├── research.md
├── data-model.md
├── quickstart.md
├── contracts/
│   ├── openapi.yaml
│   └── ui-routes.md
├── checklists/
│   └── requirements.md
└── tasks.md
```

### Source Code (repository root)

```text
cmd/
└── server/
    └── main.go

internal/
├── config/
├── db/
├── fetchers/
│   ├── articles/
│   ├── items/
│   └── looks/
├── models/
├── processing/
├── repository/
├── scheduler/
└── web/
    ├── static/
    └── templates/

tests/
├── contract/
├── integration/
└── unit/
```

**Structure Decision**: Use the single-service Go project structure from `design.md`. `cmd/server` owns process startup, `internal` owns application modules, `internal/web` owns server-rendered templates and static assets, and `tests` holds cross-package contract and integration coverage when package-local tests are not sufficient.

## Phase 0: Research Summary

Research decisions are recorded in [research.md](./research.md). All technical context unknowns are resolved.

## Phase 1: Design Summary

Data model is recorded in [data-model.md](./data-model.md). Interface contracts are recorded in [contracts/openapi.yaml](./contracts/openapi.yaml) and [contracts/ui-routes.md](./contracts/ui-routes.md). Developer startup and verification flow is recorded in [quickstart.md](./quickstart.md).

## Constitution Check - Post-Design

The post-design review found no active constitutional gates because the constitution remains uncustomized. The generated artifacts remain aligned with the spec constraints: single service, no authentication, no personalization, no frontend framework, bounded data endpoints, 7-day cleanup, and server-rendered UI backed by stored data.

**Post-Design Gate Result**: PASS, no active gates to evaluate.

## Complexity Tracking

No constitution violations or justified complexity exceptions.
